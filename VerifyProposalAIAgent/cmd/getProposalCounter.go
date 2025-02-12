package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/truthvalidator/TruthValidatorAgent/VerifyProposalAIAgent/connection"
	"github.com/truthvalidator/TruthValidatorAgent/VerifyProposalAIAgent/contracts/TruthValidatorSentientNet"
)

var getProposalCounterCmd = &cobra.Command{
	Use:   "getProposalCounter [contract_address]",
	Short: "Get the current proposal counter",
	Args:  cobra.ExactArgs(1),
	Run:   GetProposalCounter,
}

func init() {
	RootCmd.AddCommand(getProposalCounterCmd)
}

func GetProposalCounter(cmd *cobra.Command, args []string) {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(time.Second*60))
	defer cancelCtx()

	contractAddr, err := ParseAddress(args[0])
	if err != nil {
		ExitWithError("Unable to parse contract address", err)
	}

	conn, err := connection.NewConnection(ctx, GetNetworkAddress())
	if err != nil {
		ExitWithError("Unable to connect", err)
	}

	sas, err := TruthValidatorSentientNet.NewTruthValidatorSentientNet(contractAddr, conn.Client)
	if err != nil {
		ExitWithError("Unable to get instance of contract", err)
	}

	counter, err := sas.ProposalCounter(&bind.CallOpts{From: conn.Address})
	if err != nil {
		ExitWithError("Failed to retrieve proposal counter", err)
	}

	fmt.Printf("Proposal Counter: %s\n", counter.String())
}
