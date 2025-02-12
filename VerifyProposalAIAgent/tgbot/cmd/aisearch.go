package cmd

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/joho/godotenv"
	"github.com/mnt-ltd/daemore"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/truthvalidator/TruthValidatorAgent/VerifyProposalAIAgent/connection"
	"github.com/truthvalidator/TruthValidatorAgent/VerifyProposalAIAgent/contracts/TruthValidatorSentientNet"
)

var whitelist []string

func getWhitelist() []string {
	return whitelist
}

func isUserAllowed(username string) bool {
	log.Println(whitelist)
	for _, allowedUsername := range whitelist {
		if allowedUsername == username {
			return true
		}
	}
	return false
}

func init() {
	if err := loadWhitelist(); err != nil {
		log.Printf("Failed to load whitelist: %v", err)
	}
	if err := loadAdminUsernames(); err != nil {
		log.Printf("Failed to load admin usernames: %v", err)
	}
}

func addToWhitelist(username string) {
	for _, allowedUsername := range whitelist {
		if allowedUsername == username {
			return
		}
	}
	whitelist = append(whitelist, username)
	saveWhitelist()
}

func removeFromWhitelist(username string) {
	for i, allowedUsername := range whitelist {
		if allowedUsername == username {
			whitelist = append(whitelist[:i], whitelist[i+1:]...)
			saveWhitelist()
			return
		}
	}
}

// loadAdminUsernames
func loadAdminUsernames() error {
	adminUsernamesEnv := os.Getenv("TG_BOT_ADMINS")
	if adminUsernamesEnv == "" {
		return nil
	}
	adminUsernames = strings.Split(adminUsernamesEnv, ",")
	return nil
}

// aisearchCmd represents the aisearch command
var aisearchCmd = &cobra.Command{
	Use:   "aisearch",
	Short: "Start the Telegram bot service with AI search integration",
	Long:  `Start the TruthValidatorSentientNet-tgbot Telegram bot service with AI search integration.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := godotenv.Load(); err != nil {
			logger.Warn("No .env file found or error loading .env file:", zap.Error(err))
		}

		botToken := os.Getenv("TG_BOT_TOKEN")
		if botToken == "" {
			logger.Error("TG_BOT_TOKEN environment variable is required")
			return
		}

		// Start the bot with the provided token
		d, err := NewAISearchDaemon(botToken)
		if err != nil {
			logger.Error("Failed to start service:", zap.Error(err))
			return
		}
		d.Service.Run()
	},
}

func init() {
	rootCmd.AddCommand(aisearchCmd)
}

// NewAISearchDaemon initializes the AI search daemon
func NewAISearchDaemon(botToken string) (d *daemore.Daemon, err error) {
	d, err = daemore.NewDaemon(daemore.DaemonOption{
		Name:        "TruthValidatorSentientNet-tgbot",
		DisplayName: "TruthValidatorSentientNet-tgbot AISearch Service",
		Description: "A Telegram bot service with AI search integration.",
		Callback:    func() { runAISearchServer(botToken) },
	})
	return
}

// runAISearchServer starts the Telegram bot with AI search integration
func runAISearchServer(botToken string) {
	// Create a context that listens for interrupt signals
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Define bot options, setting the AI search handler as the default handler
	opts := []bot.Option{
		bot.WithDefaultHandler(aiSearchHandler(botToken)),
	}

	// Initialize the bot with the provided token and options
	b, err := bot.New(botToken, opts...)
	if err != nil {
		log.Panicf("Failed to initialize bot: %v", err)
		return
	}

	// **配置命令菜单**
	commands := []models.BotCommand{
		{
			Command:     "start",
			Description: "Submit information for truth verification",
		},
		// 你可以添加更多命令到菜单中，例如 /help, /about 等
	}

	_, err = b.SetMyCommands(ctx, &bot.SetMyCommandsParams{
		Commands: commands,
	})
	if err != nil {
		log.Printf("Failed to set bot commands: %v", err)
	} else {
		log.Println("Bot commands set successfully.")
	}

	// Start the bot (blocking call)
	log.Println("Bot started successfully with AI search integration")
	b.Start(ctx)

	// Bot will stop gracefully when the context is canceled
	log.Println("Bot stopped gracefully")
}

type Evidence struct {
	URL         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type APIResponse struct {
	Evidences []Evidence  `json:"evidences"`
	ID        string      `json:"id"`
	Related   interface{} `json:"related"`
	BundledIn struct {
		DataTxID      string `json:"dataTxId"`
		EntityID      string `json:"entityId"`
		EntityName    string `json:"entityName"`
		MetadataTxID  string `json:"metadataTxId"`
		Web3UploadURL string `json:"web3uploadURL"`
	} `json:"bundledIn"`
}

func reloadWhitelist() {
	whitelist = getWhitelist()
}

// getAdminUsernames
func getAdminUsernames() []string {
	adminUsernamesEnv := os.Getenv("TG_BOT_ADMINS")
	if adminUsernamesEnv == "" {
		return []string{}
	}
	return strings.Split(adminUsernamesEnv, ",")
}

var adminUsernames = getAdminUsernames()

func isAdmin(username string) bool {
	for _, adminUsername := range adminUsernames {
		if adminUsername == username {
			return true
		}
	}
	return false
}

// addAdmin
func addAdmin(username string) {
	for _, adminUsername := range adminUsernames {
		if adminUsername == username {
			return
		}
	}
	adminUsernames = append(adminUsernames, username)
}

// removeAdmin
func removeAdmin(username string) {
	for i, adminUsername := range adminUsernames {
		if adminUsername == username {
			adminUsernames = append(adminUsernames[:i], adminUsernames[i+1:]...)
			return
		}
	}
}

// saveWhitelist
func saveWhitelist() error {
	file, err := os.Create("whitelist.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(whitelist)
}

// loadWhitelist
func loadWhitelist() error {
	whitelistEnv := os.Getenv("TG_BOT_WHITELIST")
	if whitelistEnv != "" {
		envWhitelist := strings.Split(whitelistEnv, ",")
		whitelist = append(whitelist, envWhitelist...)
	}

	file, err := os.Open("whitelist.json")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	var fileWhitelist []string
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&fileWhitelist); err != nil {
		return err
	}

	whitelist = append(whitelist, fileWhitelist...)

	whitelist = removeDuplicates(whitelist)

	return nil
}

func removeDuplicates(list []string) []string {
	seen := make(map[string]bool)
	result := []string{}
	for _, item := range list {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}

func adminHandler(botToken string) func(context.Context, *bot.Bot, *models.Update) {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message == nil || update.Message.Text == "" {
			return
		}

		username := update.Message.From.Username
		if !isAdmin(username) {

			_, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "Sorry, you are not authorized to use this command.",
				ReplyParameters: &models.ReplyParameters{
					MessageID: update.Message.ID,
				},
			})
			if err != nil {
				log.Printf("Failed to send admin check message: %v", err)
			}
			return
		}

		parts := strings.Fields(update.Message.Text)
		if len(parts) < 2 {
			return
		}

		command := parts[0]
		targetUsername := strings.TrimPrefix(parts[1], "@") // remove @

		switch command {
		case "/add_admin":
			addAdmin(targetUsername)
			_, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   fmt.Sprintf("User @%s has been added to the admin list.", targetUsername),
			})
			if err != nil {
				log.Printf("Failed to send add admin message: %v", err)
			}

		case "/remove_admin":
			removeAdmin(targetUsername)
			_, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   fmt.Sprintf("User @%s has been removed from the admin list.", targetUsername),
			})
			if err != nil {
				log.Printf("Failed to send remove admin message: %v", err)
			}

		case "/add_to_whitelist":
			addToWhitelist(targetUsername)
			_, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   fmt.Sprintf("User @%s has been added to the whitelist.", targetUsername),
			})
			if err != nil {
				log.Printf("Failed to send add to whitelist message: %v", err)
			}

		case "/remove_from_whitelist":
			removeFromWhitelist(targetUsername)
			_, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   fmt.Sprintf("User @%s has been removed from the whitelist.", targetUsername),
			})
			if err != nil {
				log.Printf("Failed to send remove from whitelist message: %v", err)
			}

		case "/reload_whitelist":
			if err := loadWhitelist(); err != nil {
				log.Printf("Failed to reload whitelist: %v", err)
			}
			_, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: update.Message.Chat.ID,
				Text:   "Whitelist reloaded successfully.",
			})
			if err != nil {
				log.Printf("Failed to send reload whitelist message: %v", err)
			}
		}
	}
}

func formatForTelegram(text string) string {
	var resultBuilder strings.Builder

	// Store placeholders for different types of content
	var (
		codeBlocks = make(map[string]string)
		mathBlocks = make(map[string]string)
		inlineMath = make(map[string]string)
		inlineCode = make(map[string]string)
	)

	// First, extract and save math blocks ($$...$$)
	mathBlockCount := 0
	text = regexp.MustCompile(`\$\$([\s\S]*?)\$\$`).ReplaceAllStringFunc(text, func(block string) string {
		content := strings.Trim(block, "$")
		placeholder := fmt.Sprintf("___MATH_BLOCK_%d___", mathBlockCount)
		mathBlocks[placeholder] = content
		mathBlockCount++
		return placeholder
	})

	// Extract and save inline math ($...$)
	inlineMathCount := 0
	text = regexp.MustCompile(`\$([^\$\n]+?)\$`).ReplaceAllStringFunc(text, func(block string) string {
		content := strings.Trim(block, "$")
		placeholder := fmt.Sprintf("___INLINE_MATH_%d___", inlineMathCount)
		inlineMath[placeholder] = content
		inlineMathCount++
		return placeholder
	})

	// Extract code blocks
	codeBlockCount := 0
	text = regexp.MustCompile("```(?:.*?)\n([\\s\\S]*?)```").ReplaceAllStringFunc(text, func(block string) string {
		content := regexp.MustCompile("^```.*?\n").ReplaceAllString(block, "")
		content = strings.TrimSuffix(content, "```")
		content = strings.TrimSpace(content)
		placeholder := fmt.Sprintf("___CODE_BLOCK_%d___", codeBlockCount)
		codeBlocks[placeholder] = content
		codeBlockCount++
		return placeholder
	})

	// Extract inline code
	inlineCodeCount := 0
	text = regexp.MustCompile("`([^`]+)`").ReplaceAllStringFunc(text, func(block string) string {
		content := strings.Trim(block, "`")
		placeholder := fmt.Sprintf("___INLINE_CODE_%d___", inlineCodeCount)
		inlineCode[placeholder] = content
		inlineCodeCount++
		return placeholder
	})

	// Define escape characters
	escapeChars := []struct {
		char    string
		escaped string
	}{
		{"\\", "\\\\"},
		{"_", "\\_"},
		{"[", "\\["},
		{"]", "\\]"},
		{"(", "\\("},
		{")", "\\)"},
		{"~", "\\~"},
		{">", "\\>"},
		{"#", "\\#"},
		{"+", "\\+"},
		{"-", "\\-"},
		{"=", "\\="},
		{"|", "\\|"},
		{"{", "\\{"},
		{"}", "\\}"},
		{".", "\\."},
		{"!", "\\!"},
	}

	// Process text line by line
	lines := strings.Split(text, "\n")
	lastLineEmpty := false

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			if !lastLineEmpty {
				resultBuilder.WriteString("\n")
			}
			lastLineEmpty = true
			continue
		}
		lastLineEmpty = false

		// Handle math blocks
		if mathBlock, exists := mathBlocks[line]; exists {
			// Format math block as monospace
			resultBuilder.WriteString("```\n")
			resultBuilder.WriteString(mathBlock)
			resultBuilder.WriteString("\n```\n")
			continue
		}

		// Handle code blocks
		if codeBlock, exists := codeBlocks[line]; exists {
			resultBuilder.WriteString("```\n")
			resultBuilder.WriteString(codeBlock)
			resultBuilder.WriteString("\n```\n")
			continue
		}

		// Process the line
		processedLine := line

		// Handle markdown headers
		if strings.HasPrefix(processedLine, "###") {
			processedLine = "*" + strings.TrimSpace(strings.TrimPrefix(processedLine, "###")) + "*"
		}

		// Handle bullet points
		if strings.HasPrefix(processedLine, "- ") {
			processedLine = "• " + strings.TrimPrefix(processedLine, "- ")
		}

		// Restore inline math as monospace
		for placeholder, content := range inlineMath {
			processedLine = strings.ReplaceAll(processedLine, placeholder, fmt.Sprintf("`%s`", content))
		}

		// Restore inline code
		for placeholder, content := range inlineCode {
			processedLine = strings.ReplaceAll(processedLine, placeholder, fmt.Sprintf("`%s`", content))
		}

		// Escape special characters outside of code blocks
		segments := strings.Split(processedLine, "`")
		for i := range segments {
			if i%2 == 0 { // Only escape outside of code/math blocks
				for _, ec := range escapeChars {
					segments[i] = strings.ReplaceAll(segments[i], ec.char, ec.escaped)
				}
			}
		}
		processedLine = strings.Join(segments, "`")

		resultBuilder.WriteString(processedLine)
		resultBuilder.WriteString("\n")
	}

	// Clean up the result
	result := resultBuilder.String()
	result = strings.TrimSpace(result)
	result = regexp.MustCompile(`\n{3,}`).ReplaceAllString(result, "\n\n")

	return result
}

func escapeForTelegram(text string) string {
	escapeChars := []struct {
		char    string
		escaped string
	}{
		{"_", "\\_"},
		{"*", "\\*"},
		{"`", "\\`"},
		{"[", "\\["},
		{"]", "\\]"},
		{"(", "\\("},
		{")", "\\)"},
		{"~", "\\~"},
		{">", "\\>"},
		{"#", "\\#"},
		{"+", "\\+"},
		{"-", "\\-"},
		{"=", "\\="},
		{"|", "\\|"},
		{"{", "\\{"},
		{"}", "\\}"},
		{".", "\\."},
		{"!", "\\!"},
	}

	for _, ec := range escapeChars {
		text = strings.ReplaceAll(text, ec.char, ec.escaped)
	}

	return text
}

func listenForProposalResult(ctx context.Context, b *bot.Bot, chatID int64, proposalID uint64, replyToID int) {
	// Create a context with a timeout of 1 hour
	ctx, cancel := context.WithTimeout(ctx, time.Hour)
	defer cancel()

	contractAddr := os.Getenv("CONTRACT_ADDRESS")
	if contractAddr == "" {
		log.Printf("CONTRACT_ADDRESS environment variable is not set")
		return
	}

	contractAddrParsed, err := ParseAddress(contractAddr)
	if err != nil {
		log.Printf("Failed to parse contract address: %v", err)
		return
	}

	// Create a new connection
	conn, err := connection.NewConnection(ctx, GetNetworkAddress())
	if err != nil {
		log.Printf("Failed to connect to the network: %v", err)
		return
	}
	defer conn.Client.Close()

	// Create an instance of the contract
	sas, err := TruthValidatorSentientNet.NewTruthValidatorSentientNet(contractAddrParsed, conn.Client)
	if err != nil {
		log.Printf("Failed to get contract instance: %v", err)
		return
	}

	// Create a channel to receive ProposalFinalized events
	proposalFinalizedChan := make(chan *TruthValidatorSentientNet.TruthValidatorSentientNetProposalFinalized)
	sub, err := sas.WatchProposalFinalized(&bind.WatchOpts{Context: ctx}, proposalFinalizedChan, []*big.Int{big.NewInt(int64(proposalID))})
	if err != nil {
		log.Printf("Failed to watch ProposalFinalized events: %v", err)
		return
	}
	defer sub.Unsubscribe()

	log.Printf("Listening for ProposalFinalized event for Proposal ID: %d", proposalID)

	// Wait for the event or timeout
	select {
	case event := <-proposalFinalizedChan:
		log.Printf("ProposalFinalized event received for Proposal ID: %d", proposalID)

		// Format the final result
		finalResult := "❌"
		if event.FinalResult {
			finalResult = "✅"
		}

		// Format the voter results using blocks and line breaks
		var voterResults strings.Builder
		voterResults.WriteString("Voter Results:\n")

		for _, vote := range event.VoterResults {
			voteInfo := "TRUE"
			if !vote.IsApproved {
				voteInfo = "FALSE"
			}
			reason := vote.Reason

			voterResults.WriteString("----------------------\n") // Separator between voters
			voterResults.WriteString(fmt.Sprintf("AI-Voter: %s\n", vote.Voter.String()))
			voterResults.WriteString(fmt.Sprintf("Vote: %s \n", voteInfo))
			voterResults.WriteString(fmt.Sprintf("Reason: %s\n\n", reason))
			// voterResults.WriteString("\n") // Extra line break after each voter block
		}
		// voterResults.WriteString("----------------------\n") // Final separator

		// Send the final result and voter results to the user with more line breaks
		messageText := fmt.Sprintf("Verification Request Result (ID: %d):\n\n"+ // Changed message title to English
			"Final Result: %s\n\n"+
			"%s\n", proposalID, finalResult, voterResults.String())

		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   messageText,
			ReplyParameters: &models.ReplyParameters{
				MessageID: replyToID,
			},
		})
		if err != nil {
			log.Printf("Failed to send proposal result message: %v", err)
		}

	case <-ctx.Done():
		log.Printf("Timeout reached while waiting for ProposalFinalized event for Proposal ID: %d", proposalID)
		_, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   fmt.Sprintf("No result received for Proposal ID %d within 1 hour. Please check later.", proposalID),
			ReplyParameters: &models.ReplyParameters{
				MessageID: replyToID,
			},
		})
		if err != nil {
			log.Printf("Failed to send timeout message: %v", err)
		}
	}
}

func aiSearchHandler(botToken string) func(context.Context, *bot.Bot, *models.Update) {
	// State to track if a user is waiting to submit a proposal
	var proposalSubmissionState = make(map[int64]bool) // ChatID -> IsSubmittingProposal

	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered from panic: %v", r)
				_, err := b.SendMessage(ctx, &bot.SendMessageParams{
					ChatID: update.Message.Chat.ID,
					Text:   "Sorry, an error occurred. Please try again later.",
					ReplyParameters: &models.ReplyParameters{
						MessageID: update.Message.ID,
					},
				})
				if err != nil {
					log.Printf("Failed to send panic recovery message: %v", err)
				}
			}
		}()

		username := update.Message.From.Username
		chatID := update.Message.Chat.ID
		log.Println("UserName:", username)
		if !isUserAllowed(username) {
			_, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: chatID,
				Text:   "Sorry, you are not allowed to use this bot.",
				ReplyParameters: &models.ReplyParameters{
					MessageID: update.Message.ID,
				},
			})
			if err != nil {
				log.Printf("Failed to send whitelist message: %v", err)
			}
			return
		}

		if update.Message == nil || update.Message.Text == "" {
			_, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: chatID,
				Text:   "Hi! I only understand text messages. Please send me a text question! 😊",
				ReplyParameters: &models.ReplyParameters{
					MessageID: update.Message.ID,
				},
			})
			if err != nil {
				log.Printf("Failed to send non-text message response: %v", err)
			}
			return
		}

		if proposalSubmissionState[chatID] {
			proposalContent := update.Message.Text
			log.Println(proposalContent)
			proposalSubmissionState[chatID] = false // Reset state

			contractAddr := os.Getenv("CONTRACT_ADDRESS") // Get contract address from environment variable
			if contractAddr == "" {
				_, err := b.SendMessage(ctx, &bot.SendMessageParams{
					ChatID:    chatID,
					Text:      "Error: CONTRACT_ADDRESS environment variable is not set.",
					ParseMode: "MarkdownV2",
					ReplyParameters: &models.ReplyParameters{
						MessageID: update.Message.ID,
					},
				})
				if err != nil {
					log.Printf("Failed to send contract address error message: %v", err)
				}
				return
			}

			txHash, proposalId, err := SubmitProposalToContract(ctx, contractAddr, proposalContent)
			if err != nil {
				_, err := b.SendMessage(ctx, &bot.SendMessageParams{
					ChatID: chatID,
					Text:   fmt.Sprintf("Failed to submit proposal: %v", formatError(err)),
					// ParseMode: "MarkdownV2",
					ReplyParameters: &models.ReplyParameters{
						MessageID: update.Message.ID,
					},
				})
				if err != nil {
					log.Printf("Failed to send Verification submission error message: %v", err)
				}
			} else {
				_, err := b.SendMessage(ctx, &bot.SendMessageParams{
					ChatID: chatID,
					Text: fmt.Sprintf("Verification Request Submission Successful🎉\n\n"+
						"Transaction Hash: %s\n\n"+
						"Verification ID: %d\n\n"+
						"Thank you for your submission. The Verification will be reviewed and voted on by the AI Agents.", txHash, proposalId),
					// ParseMode: "MarkdownV2",
					ReplyParameters: &models.ReplyParameters{
						MessageID: update.Message.ID,
					},
				})
				if err != nil {
					log.Printf("Failed to send Verification submission success message: %v", err)
				}

				// Start a goroutine to listen for the proposal result
				go listenForProposalResult(ctx, b, chatID, proposalId, update.Message.ID)
			}
			return
		}

		if isAdmin(username) {

			parts := strings.Fields(update.Message.Text)
			if len(parts) >= 2 {
				command := parts[0]
				targetUsername := strings.TrimPrefix(parts[1], "@") // remove @

				switch command {
				case "/add_admin":
					addAdmin(targetUsername)
					_, err := b.SendMessage(ctx, &bot.SendMessageParams{
						ChatID: chatID,
						Text:   fmt.Sprintf("User @%s has been added to the admin list.", targetUsername),
					})
					if err != nil {
						log.Printf("Failed to send add admin message: %v", err)
					}
					return

				case "/remove_admin":
					removeAdmin(targetUsername)
					_, err := b.SendMessage(ctx, &bot.SendMessageParams{
						ChatID: chatID,
						Text:   fmt.Sprintf("User @%s has been removed from the admin list.", targetUsername),
					})
					if err != nil {
						log.Printf("Failed to send remove admin message: %v", err)
					}
					return

				case "/add_to_whitelist":
					addToWhitelist(targetUsername)
					_, err := b.SendMessage(ctx, &bot.SendMessageParams{
						ChatID: chatID,
						Text:   fmt.Sprintf("User @%s has been added to the whitelist.", targetUsername),
					})
					if err != nil {
						log.Printf("Failed to send add to whitelist message: %v", err)
					}
					return

				case "/remove_from_whitelist":
					removeFromWhitelist(targetUsername)
					_, err := b.SendMessage(ctx, &bot.SendMessageParams{
						ChatID: chatID,
						Text:   fmt.Sprintf("User @%s has been removed from the whitelist.", targetUsername),
					})
					if err != nil {
						log.Printf("Failed to send remove from whitelist message: %v", err)
					}
					return

				case "/reload_whitelist":
					if err := loadWhitelist(); err != nil {
						log.Printf("Failed to reload whitelist: %v", err)
					}
					_, err := b.SendMessage(ctx, &bot.SendMessageParams{
						ChatID: chatID,
						Text:   "Whitelist reloaded successfully.",
					})
					if err != nil {
						log.Printf("Failed to send reload whitelist message: %v", err)
					}
					return
				}
			}
		}

		if update.Message.Text == "/start" {
			proposalSubmissionState[chatID] = true
			_, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: chatID,
				Text:   "🔎  Enter the information you want to verify. Just type it in here!",
				ReplyParameters: &models.ReplyParameters{
					MessageID: update.Message.ID,
				},
			})
			if err != nil {
				log.Printf("Failed to send Verification content request message: %v", err)
			}
			return
		}

		var inputText string
		if update.Message.ReplyToMessage != nil && update.Message.ReplyToMessage.Text != "" {
			// 如果有回复信息，将其追加到输入的首部，并使用 Markdown 引用符号修饰
			replyText := update.Message.ReplyToMessage.Text
			inputText = fmt.Sprintf("%s ```%s```", update.Message.Text, replyText)
		} else {
			// 如果没有回复信息，直接使用用户输入的内容
			inputText = update.Message.Text
		}

		sentMessage, err := b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:    chatID,
			Text:      "Your request is being processed\\.\\.\\.",
			ParseMode: "MarkdownV2",
			ReplyParameters: &models.ReplyParameters{
				MessageID: update.Message.ID,
			},
		})
		if err != nil {
			log.Printf("Failed to send initial message: %v", err)
			return
		}

		// call API
		response, err := callAISearchAPI(ctx, inputText)
		if err != nil {
			log.Printf("API call error: %v", err)
			_, err = b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:    chatID,
				Text:      "Call API FAIL.",
				ParseMode: "MarkdownV2",
				ReplyParameters: &models.ReplyParameters{
					MessageID: update.Message.ID,
				},
			})
			if err != nil {
				log.Printf("Failed to send error message: %v", err)
			}
			return
		}
		defer response.Body.Close()

		var messageBuilder strings.Builder
		scanner := bufio.NewScanner(response.Body)
		var isFirstLine = true
		var lastContent string

		// buffer size 30
		const bufferSize = 30
		var buffer strings.Builder

		for scanner.Scan() {
			line := scanner.Text()
			if isFirstLine {
				isFirstLine = false
				continue
			}
			if strings.Contains(line, `"bundledIn"`) {
				break
			}

			messageBuilder.WriteString(line)
			messageBuilder.WriteString("\n")

			buffer.WriteString(line)
			buffer.WriteString("\n")

			// if content > buffer size then send
			if buffer.Len() >= bufferSize {
				newContent := messageBuilder.String()
				if newContent != lastContent {
					err = sendOrUpdateMessage(ctx, b, chatID,
						sentMessage.ID, newContent, true, update.Message.ID)
					if err != nil {
						log.Printf("Update error: %v", err)
						continue
					}
					lastContent = newContent
					buffer.Reset()
				}
			}
		}

		// send else buf
		if buffer.Len() > 0 {
			newContent := messageBuilder.String()
			if newContent != lastContent {
				err = sendOrUpdateMessage(ctx, b, chatID,
					sentMessage.ID, newContent, true, update.Message.ID)
				if err != nil {
					log.Printf("Update error: %v", err)
				}
			}
		}

		if scanner.Err() != nil {
			log.Printf("Scanner error: %v", scanner.Err())
			return
		}
	}
}

func sendOrUpdateMessage(ctx context.Context, b *bot.Bot, chatID int64, messageID int, text string, isUpdate bool, replyToID int) error {
	formattedText := formatForTelegram(text)

	if strings.TrimSpace(formattedText) == "" {
		return nil
	}

	const maxLength = 4000
	var parts []string

	if len(formattedText) > maxLength {
		for len(formattedText) > 0 {
			cutPoint := maxLength
			if len(formattedText) > maxLength {
				lastNewline := strings.LastIndex(formattedText[:maxLength], "\n")
				if lastNewline > 0 {
					cutPoint = lastNewline
				}
			}

			// Ensure cutPoint is within bounds
			if cutPoint > len(formattedText) {
				cutPoint = len(formattedText)
			}

			parts = append(parts, formattedText[:cutPoint])
			formattedText = strings.TrimSpace(formattedText[cutPoint:])
		}
	} else {
		parts = []string{formattedText}
	}

	// send or update
	for partIndex, part := range parts {
		if isUpdate && partIndex == 0 {
			_, err := b.EditMessageText(ctx, &bot.EditMessageTextParams{
				ChatID:    chatID,
				MessageID: messageID,
				Text:      part,
				ParseMode: "MarkdownV2",
			})
			if err != nil {
				return fmt.Errorf("failed to update message: %w", err)

			}
		} else {
			_, err := b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:    chatID,
				Text:      part,
				ParseMode: "MarkdownV2",
				ReplyParameters: &models.ReplyParameters{
					MessageID: replyToID,
				},
			})
			if err != nil {

				return fmt.Errorf("failed to send message: %w", err)

			}
		}
	}

	return nil
}

// callAISearchAPI calls the AI search API with the provided question
func callAISearchAPI(ctx context.Context, question string) (*http.Response, error) {

	log.Println("callAISearchAPI question:", question)
	url := os.Getenv("AI_SEARCH_URL")
	payload := map[string]interface{}{
		"question":                 question,
		"language":                 "en",
		"prog_lang":                "",
		"field":                    "comm",
		"blockchainStorageEnabled": false,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	req.Header.Set("content-type", "text/plain;charset=UTF-8")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	return resp, nil
}

func SubmitProposalToContract(ctx context.Context, contractAddrStr string, content string) (string, uint64, error) {
	ctxTimeout, cancelCtx := context.WithTimeout(ctx, time.Duration(time.Second*60))
	defer cancelCtx()

	contractAddr, err := ParseAddress(contractAddrStr)
	if err != nil {
		return "", 0, fmt.Errorf("unable to parse contract address: %w", err)
	}

	conn, err := connection.NewConnection(ctxTimeout, GetNetworkAddress())
	if err != nil {
		return "", 0, fmt.Errorf("unable to connect: %w", err)
	}
	defer conn.Client.Close() // Close the connection when done

	sas, err := TruthValidatorSentientNet.NewTruthValidatorSentientNet(contractAddr, conn.Client)
	if err != nil {
		return "", 0, fmt.Errorf("unable to get instance of contract: %w", err)
	}

	auth, err := conn.PrepareNextTx(ctxTimeout)
	if err != nil {
		return "", 0, fmt.Errorf("failed to prepare transaction: %w", err)
	}
	auth.From = conn.Address

	tx, err := sas.SubmitProposal(auth, content)
	if err != nil {
		return "", 0, fmt.Errorf("failed to submit Verification: %w", err)
	}

	// Wait for the transaction to be mined and get the receipt
	receipt, err := bind.WaitMined(ctxTimeout, conn.Client, tx)
	if err != nil {
		return "", 0, fmt.Errorf("failed to get transaction receipt: %w", err)
	}

	// Parse the ProposalSubmitted event from the receipt logs
	var proposalId uint64
	for _, log := range receipt.Logs {
		event, err := sas.ParseProposalSubmitted(*log)
		if err == nil {
			proposalId = event.ProposalId.Uint64() // Extract the proposal ID
			break
		}
	}

	if proposalId == 0 {
		return "", 0, fmt.Errorf("failed to retrieve Verification ID from event logs")
	}

	return tx.Hash().Hex(), proposalId, nil
}

func formatError(err error) string {
	return strings.ReplaceAll(err.Error(), "_", "\\_") // Escape underscores in error messages
}
