package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"text/tabwriter"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/spf13/cobra"

	"github.com/truthvalidator/TruthValidatorAgent/VerifyProposalAIAgent/connection"
	"github.com/truthvalidator/TruthValidatorAgent/VerifyProposalAIAgent/contracts/TruthValidatorSentientNet"
)

// aigentListenProposalFinalizedCmd represents the aigent_listen_proposal_finalized command
var aigentListenProposalFinalizedCmd = &cobra.Command{
	Use:   "aigent_listen_proposal_finalized [contract_address]",
	Short: "Listen for ProposalFinalized events and process the finalized proposal",
	Args:  cobra.ExactArgs(1),
	Run:   aigentListenProposalFinalized,
}

func init() {
	RootCmd.AddCommand(aigentListenProposalFinalizedCmd)
}

func aigentListenProposalFinalized(cmd *cobra.Command, args []string) {
	contractAddr, err := ParseAddress(args[0])
	if err != nil {
		ExitWithError("Unable to parse contract address", err)
		return
	}

	for {
		ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(time.Minute*1)) // 1 minute timeout

		// Create a new connection
		conn, err := connection.NewConnection(ctx, GetNetworkAddress())
		if err != nil {
			log.Printf("Unable to connect: %v, retrying in 10 seconds", err)
			cancelCtx()
			time.Sleep(10 * time.Second)
			continue
		}
		log.Println("Connected to the network.")

		// Create an instance of the contract using the Sapphire backend
		sas, err := TruthValidatorSentientNet.NewTruthValidatorSentientNet(contractAddr, conn.Client)
		if err != nil {
			log.Printf("Unable to get instance of contract: %v, retrying in 10 seconds", err)
			conn.Client.Close()
			cancelCtx()
			time.Sleep(10 * time.Second)
			continue
		}
		log.Println("Contract instance created.")

		// Start listening for ProposalFinalized events
		proposalFinalizedChan := make(chan *TruthValidatorSentientNet.TruthValidatorSentientNetProposalFinalized)
		sub, err := sas.WatchProposalFinalized(&bind.WatchOpts{Context: ctx}, proposalFinalizedChan, nil)
		if err != nil {
			log.Printf("Failed to watch ProposalFinalized events: %v, retrying in 10 seconds", err)
			conn.Client.Close()
			cancelCtx()
			time.Sleep(10 * time.Second)
			continue
		}
		log.Println("Listening for ProposalFinalized events...")

	eventLoop:
		for {
			select {
			case event := <-proposalFinalizedChan:

				go func(event *TruthValidatorSentientNet.TruthValidatorSentientNetProposalFinalized) {
					fmt.Printf("Proposal finalized! Proposal ID: %s, Final Result: %v\n", event.ProposalId.String(), event.FinalResult)

					fmt.Println("Voter Results:")
					// Initialize tabwriter for formatted output
					w := tabwriter.NewWriter(os.Stdout, 12, 0, 2, ' ', 0)
					fmt.Fprintln(w, "Voter\tVote\tReason\t")   // Header row
					fmt.Fprintln(w, "-------\t----\t------\t") // Separator row

					for _, vote := range event.VoterResults {
						voteInfo := "Approved"
						if !vote.IsApproved {
							voteInfo = "Rejected"
						}
						reason := vote.Reason
						if len(reason) > 100 { // Limit reason length for better table display
							reason = reason[:100] + "..."
						}
						fmt.Fprintf(w, "%s\t%s\t%s\t\n", vote.Voter.String(), voteInfo, reason)
					}
					w.Flush() // Flush to output table

					// Process the finalized proposal (e.g., log the result, notify users, or perform other actions)
					processFinalizedProposal(ctx, event.ProposalId, event.FinalResult)
				}(event)

			case err := <-sub.Err():
				log.Printf("Subscription error: %v, attempting to reconnect: %v", err, errors.Unwrap(err))
				break eventLoop

			case <-ctx.Done():
				if errors.Is(ctx.Err(), context.DeadlineExceeded) {
					log.Println("Context timeout. Reconnecting after timeout...")
				} else if errors.Is(ctx.Err(), context.Canceled) {
					log.Println("Context cancelled. Reconnecting...")
				} else {
					log.Printf("Context done with error: %v. Reconnecting...", ctx.Err())
				}
				break eventLoop
			}
		}

		// Cleanup resources before reconnecting
		sub.Unsubscribe()
		conn.Client.Close()
		cancelCtx()

		fmt.Println("Reconnecting...")
		time.Sleep(1 * time.Second)
	}
}

// processFinalizedProposal processes the finalized proposal
func processFinalizedProposal(ctx context.Context, proposalId *big.Int, finalResult bool) {
	log.Printf("Processing finalized proposal ID: %s, Final Result: %v\n", proposalId.String(), finalResult)

	// Here you can add any additional logic to process the finalized proposal,
	// such as notifying users, storing the result in a database, or triggering other actions.
}
