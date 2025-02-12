package cmd

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/spf13/cobra"

	"github.com/truthvalidator/TruthValidatorAgent/VerifyProposalAIAgent/connection"
	"github.com/truthvalidator/TruthValidatorAgent/VerifyProposalAIAgent/contracts/TruthValidatorSentientNet"
)

var voteCmd = &cobra.Command{
	Use:   "vote [contract_address] [proposalId] [isApproved] [reason]",
	Short: "Cast a vote on a specific proposal",
	Args:  cobra.ExactArgs(4),
	Run:   Vote,
}

func init() {
	RootCmd.AddCommand(voteCmd)
}

func Vote(cmd *cobra.Command, args []string) {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(time.Second*60))
	defer cancelCtx()

	contractAddr, err := ParseAddress(args[0])
	if err != nil {
		ExitWithError("Unable to parse contract address", err)
	}

	proposalId, err := ParseUint256(args[1])
	if err != nil {
		ExitWithError("Invalid proposal ID", err)
	}

	// Parse the isApproved argument (true or false)
	isApproved, err := strconv.ParseBool(args[2])
	if err != nil {
		ExitWithError("Invalid isApproved value", err)
	}

	// The reason is a string, no need to parse
	reason := args[3]

	conn, err := connection.NewConnection(ctx, GetNetworkAddress())
	if err != nil {
		ExitWithError("Unable to connect", err)
	}

	sas, err := TruthValidatorSentientNet.NewTruthValidatorSentientNet(contractAddr, conn.Client)
	if err != nil {
		ExitWithError("Unable to get instance of contract", err)
	}

	// Prepare the transaction options using the PrepareNextTx method
	auth, err := conn.PrepareNextTx(ctx)
	if err != nil {
		ExitWithError("Failed to prepare transaction", err)
	}

	// Set the from address in the transact options
	auth.From = conn.Address
	log.Println("auth.From", auth.From)
	tx, err := sas.Vote(auth, proposalId, isApproved, reason)
	if err != nil {
		ExitWithError("Failed to cast vote", err)
	}

	fmt.Printf("Vote cast successfully. Transaction hash: %s\n", tx.Hash().Hex())
}
