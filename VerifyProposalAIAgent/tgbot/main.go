// Package main implements the Telegram bot interface for the TruthValidator AI Agent system.
// The bot handles:
// - Proposal verification requests
// - Voting coordination 
// - Real-time status updates
// - Admin commands and notifications
package main

import (
	"log"
	"os"

	"github.com/truthvalidator/TruthValidatorAgent/VerifyProposalAIAgent/tgbot/cmd"
)

func init() {
	// Configure detailed logging with microsecond precision and full file paths
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	log.SetOutput(os.Stdout) // Explicitly set output to stdout
}

// main is the entry point for the Telegram bot service.
// It initializes and executes the root Cobra command which handles:
// - Command line flags
// - Configuration loading  
// - Bot service startup
func main() {
	cmd.Execute()
}
