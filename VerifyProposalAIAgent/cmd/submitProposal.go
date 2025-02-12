package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/spf13/cobra"

	"github.com/truthvalidator/TruthValidatorAgent/VerifyProposalAIAgent/connection"
	"github.com/truthvalidator/TruthValidatorAgent/VerifyProposalAIAgent/contracts/TruthValidatorSentientNet"
)

var submitProposalCmd = &cobra.Command{
	Use:   "submitProposal [contract_address] [content]",
	Short: "Submit a new proposal to the TruthValidatorSentientNet contract",
	Args:  cobra.ExactArgs(2),
	Run:   SubmitProposal,
}

func init() {
	RootCmd.AddCommand(submitProposalCmd)
}

func SubmitProposal(cmd *cobra.Command, args []string) {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(time.Second*60))
	defer cancelCtx()

	contractAddr, err := ParseAddress(args[0])
	if err != nil {
		ExitWithError("Unable to parse contract address", err)
	}

	content := args[1]

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

	// log.Println("auth.From", auth.From.Hex())

	tx, err := sas.SubmitProposal(auth, content)
	if err != nil {
		ExitWithError("Failed to submit proposal", err)
	}

	fmt.Printf("Proposal submitted. Transaction hash: %s\n", tx.Hash().Hex())

	// Wait for the transaction to be mined and get the receipt
	receipt, err := bind.WaitMined(ctx, conn.Client, tx)
	if err != nil {
		ExitWithError("Failed to get transaction receipt", err)
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

	if proposalId > 0 {
		fmt.Printf("Proposal ID: %d\n", proposalId)
	} else {
		ExitWithError("Failed to retrieve proposal ID from event logs", nil)
	}
}
