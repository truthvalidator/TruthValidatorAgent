# TruthValidator Go SDK

Go SDK for interacting with TruthValidator smart contracts

## Features

- Connect to Ethereum nodes
- Wallet management
- Proposal submission
- Result querying
- Voting operations

## Installation

```bash
go get github.com/truthvalidator/TruthValidatorAgent/VerifyProposalAIAgent/sdk
```

## Usage Example

### Initialize Client

```go
package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/truthvalidator/TruthValidatorAgent/VerifyProposalAIAgent/sdk"
)

func main() {
	// Configuration
	cfg := &sdk.Config{
		RPCURL:         "https://mainnet.infura.io/v3/YOUR_PROJECT_ID",
		ContractAddress: common.HexToAddress("0x..."), // Contract address
		ChainID:        big.NewInt(1), // Mainnet
	}

	// Create client
	client, err := sdk.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// Create wallet
	wallet, err := sdk.NewWalletFromPrivateKey("YOUR_PRIVATE_KEY")
	if err != nil {
		panic(err)
	}

	// Submit proposal
	contentHash := "Qm...IPFS hash..."
	txHash, err := client.SubmitProposal(
		context.Background(),
		wallet,
		cfg.ChainID,
		contentHash,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Proposal submitted, transaction hash: %s\n", txHash.Hex())

	// Query proposal
	proposalID := big.NewInt(1)
	proposal, err := client.GetProposal(context.Background(), proposalID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Proposal content hash: %s\n", proposal.ContentHash)
	fmt.Printf("Total votes: %s\n", proposal.VoteCount.String())

	// Query vote results
	results, err := client.GetVoteResults(context.Background(), proposalID)
	if err != nil {
		panic(err)
	}
	for _, result := range results {
		fmt.Printf("Voter: %s, Result: %v, Reason: %s\n", 
			result.Voter.Hex(), result.IsApproved, result.Reason)
	}
}
```

## API Documentation

### TruthValidatorClient

- `NewClient(cfg *Config)`: Create new client
- `Close()`: Close connection
- `SubmitProposal()`: Submit new proposal
- `GetProposal()`: Get proposal information
- `GetVoteResult()`: Get single vote result
- `GetVoteResults()`: Get all vote results

### Wallet

- `NewWalletFromPrivateKey()`: Create wallet from private key
- `GetAddress()`: Get wallet address
- `NewTransactor()`: Create transaction signer
