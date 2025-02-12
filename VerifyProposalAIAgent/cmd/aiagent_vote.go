package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/accounts/abi/bind" // Import bind package
	"github.com/truthvalidator/TruthValidatorAgent/VerifyProposalAIAgent/connection"
	"github.com/truthvalidator/TruthValidatorAgent/VerifyProposalAIAgent/contracts/TruthValidatorSentientNet"
)

// aigentListenAndVoteCmd represents the aigent_listenAndVote command
var aigentListenAndVoteCmd = &cobra.Command{
	Use:   "aigent_listenAndVote [contract_address]",
	Short: "Listen for ProposalSubmitted events and automatically vote on the proposal",
	Args:  cobra.ExactArgs(1),
	Run:   aigentListenAndVote,
}

func init() {
	RootCmd.AddCommand(aigentListenAndVoteCmd)
}

func aigentListenAndVote(cmd *cobra.Command, args []string) {
	contractAddr, err := ParseAddress(args[0])
	if err != nil {
		ExitWithError("Unable to parse contract address", err)
		return // Critical error, cannot proceed without a valid address
	}

	for {
		ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(time.Minute*1)) // 1 hour timeout

		// Create a new connection
		conn, err := connection.NewConnection(ctx, GetNetworkAddress())
		if err != nil {
			log.Printf("Unable to connect: %v, retrying in 10 seconds", err)
			cancelCtx() // Ensure context is cancelled even if connection fails
			time.Sleep(10 * time.Second)
			continue // Retry connection
		}

		// Create an instance of the contract using the Sapphire backend
		sas, err := TruthValidatorSentientNet.NewTruthValidatorSentientNet(contractAddr, conn.Client)
		if err != nil {
			log.Printf("Unable to get instance of contract: %v, retrying in 10 seconds", err)
			conn.Client.Close() // Close connection if contract instantiation fails
			cancelCtx()         // Cancel context if contract instantiation fails
			time.Sleep(10 * time.Second)
			continue // Retry connection
		}

		// Start listening for ProposalSubmitted events
		proposalSubmittedChan := make(chan *TruthValidatorSentientNet.TruthValidatorSentientNetProposalSubmitted)
		sub, err := sas.WatchProposalSubmitted(&bind.WatchOpts{Context: ctx}, proposalSubmittedChan, nil)
		if err != nil {
			log.Printf("Failed to watch ProposalSubmitted events: %v, retrying in 10 seconds", err)
			conn.Client.Close() // Close connection if subscription fails
			cancelCtx()         // Cancel context if subscription fails
			time.Sleep(10 * time.Second)
			continue // Retry connection
		}

		fmt.Println("Listening for ProposalSubmitted events...")

	eventLoop: // Label for the inner loop to allow breaking from it
		for {
			select {
			case event := <-proposalSubmittedChan:
				fmt.Printf("New proposal submitted! Proposal ID: %s, Content: %s\n", event.ProposalId.String(), event.Content)

				go func(event *TruthValidatorSentientNet.TruthValidatorSentientNetProposalSubmitted) {
					// Extract the proposal ID and content
					proposalId := event.ProposalId
					content := event.Content

					// Make a decision based on the content using AI
					aiResponse, err := decideVote(ctx, content)
					if err != nil {
						log.Printf("Error during AI decision: %v", err)
						return // Skip voting if AI decision fails, maybe try again next time or implement retry logic
					}

					isApproved := aiResponse.Judge
					reason := aiResponse.Reason

					// Prepare the transaction options
					auth, err := conn.PrepareNextTx(ctx)
					if err != nil {
						log.Printf("Failed to prepare transaction: %v", err)
						return
					}
					auth.From = conn.Address

					// Vote on the proposal with the reason
					tx, err := sas.Vote(auth, proposalId, isApproved, reason)
					if err != nil {
						log.Printf("Failed to vote on proposal %s: %v", proposalId.String(), err)
						return
					}

					fmt.Printf("Voted on proposal %s. Transaction hash: %s, Vote: %v, Reason: %s\n", proposalId.String(), tx.Hash().Hex(), isApproved, reason)
				}(event)

			case err := <-sub.Err():
				log.Printf("Subscription error: %v, attempting to reconnect", err)
				break eventLoop // Break the inner loop and reconnect

			case <-ctx.Done():
				fmt.Println("Context timeout or canceled. Reconnecting after timeout...")
				break eventLoop // Break the inner loop and reconnect. Do NOT return.
			}
		}

		sub.Unsubscribe()
		conn.Client.Close()
		cancelCtx()

		fmt.Println("Reconnecting...")
		// time.Sleep(1 * time.Second) // Add a delay before attempting to reconnect
	}
}

// AIReponse struct to hold the AI API response
type AIReponse struct {
	Judge  bool   `json:"judge"`
	Reason string `json:"reason"`
}

// decideVote calls the AI API to make a decision based on the proposal content
func decideVote(ctx context.Context, content string) (*AIReponse, error) {
	log.Println("Calling AI API for decision on content:", content)

	// Call the AI API
	response, err := callAISearchAPIForDecision(ctx, content)
	if err != nil {
		log.Printf("AI API call error: %v", err)
		return nil, fmt.Errorf("call AI API failed: %w", err)
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read AI API response body: %w", err)
	}

	// Trim the response body to remove extra spaces, newlines, and Markdown ```json ``` tags
	trimmedBody := trimMarkdownJSON(string(body))

	// Parse the trimmed JSON response
	var aiResponse AIReponse
	err = json.Unmarshal([]byte(trimmedBody), &aiResponse)
	if err != nil {
		log.Printf("Failed to unmarshal AI API response: %v, body: %s", err, trimmedBody)
		return nil, fmt.Errorf("failed to unmarshal AI API response: %w, body: %s", err, trimmedBody)
	}

	log.Printf("AI API Response: Judge=%v, Reason=%s", aiResponse.Judge, aiResponse.Reason)
	return &aiResponse, nil
}

// trimMarkdownJSON removes Markdown ```json ``` tags and extra spaces/newlines
func trimMarkdownJSON(input string) string {
	// Remove the Markdown ```json ``` tags
	input = strings.TrimSpace(input)
	trimmed := strings.TrimPrefix(input, "```json")
	trimmed = strings.TrimSuffix(trimmed, "```")

	// Remove extra spaces and newlines
	trimmed = strings.TrimSpace(trimmed)

	return trimmed
}

// callAISearchAPIForDecision calls the AI search API with the provided question for decision
func callAISearchAPIForDecision(ctx context.Context, question string) (*http.Response, error) {
	log.Println("callAISearchAPIForDecision question:", question)
	url := os.Getenv("AI_DECISION_URL") // Use a different env var for decision API URL, or reuse AI_SEARCH_URL if it's the same API endpoint
	if url == "" {
		return nil, fmt.Errorf("AI_DECISION_URL environment variable is not set")
	}
	payload := map[string]interface{}{
		"question": question,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("accept", "application/json") // Expect JSON response
	req.Header.Set("content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to AI API: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, fmt.Errorf("AI API returned status %d: %s", resp.StatusCode, string(body))
	}

	return resp, nil
}

// generateVoteReason generates a reason for the vote based on the AI response
func generateVoteReason(content string, isApproved bool) string {
	// This function is kept for backward compatibility, but ideally, the reason should come from the AI response directly.
	if isApproved {
		return fmt.Sprintf("AI Agent approved the proposal related to: %s", content) // More generic reason as AI provides detailed reason
	} else {
		return "AI Agent rejected the proposal." // More generic reason as AI provides detailed reason
	}
}
