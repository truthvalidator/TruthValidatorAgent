// TruthValidatorAgent - The core agent for validating and processing truth proposals
// 
// This agent provides the main entry point for the Truth Validator system, handling:
// - Proposal validation and verification
// - Voting and consensus mechanisms 
// - Integration with blockchain networks
// - AI-assisted decision making
//
// Version: 1.0.0
// Author: TruthValidator Team
package main

import (
	"github.com/truthvalidator/TruthValidatorAgent/VerifyProposalAIAgent/cmd"
)

// main - Entry point for the TruthValidatorAgent
// Initializes and executes the root command handler that manages:
// - CLI command parsing
// - Configuration loading
// - Subcommand routing
func main() {
	cmd.Execute()
}
