package sdk

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// ProposalInfo proposal information
type ProposalInfo struct {
	Proposer    common.Address
	ContentHash string
	Status      uint8
	VoteCount   *big.Int
}

// SubmitProposal submits a new proposal
func (c *TruthValidatorClient) SubmitProposal(
	ctx context.Context,
	wallet *Wallet,
	chainID *big.Int,
	contentHash string,
) (common.Hash, error) {
	// Create transaction signer
	opts, err := wallet.NewTransactor(chainID)
	if err != nil {
		return common.Hash{}, err
	}
	opts.Context = ctx

	// Call contract to submit proposal
	tx, err := c.contract.SubmitProposal(opts, contentHash)
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), nil
}

// GetProposal gets proposal information
func (c *TruthValidatorClient) GetProposal(
	ctx context.Context,
	proposalID *big.Int,
) (*ProposalInfo, error) {
	// Query basic proposal info
	proposal, err := c.contract.Proposals(&bind.CallOpts{
		Context: ctx,
	}, proposalID)
	if err != nil {
		return nil, err
	}

	// Query vote counts
	voteCounts, err := c.contract.GetVoteCounts(&bind.CallOpts{
		Context: ctx,
	}, proposalID)
	if err != nil {
		return nil, err
	}

	totalVotes := new(big.Int).Add(voteCounts.YesVotes, voteCounts.NoVotes)

	return &ProposalInfo{
		Proposer:    common.HexToAddress("0x0"), // Default value since contract doesn't store proposer
		ContentHash: proposal.Content,
		Status:      uint8(0), // 0=active, 1=approved, 2=rejected
		VoteCount:   totalVotes,
	}, nil
}
