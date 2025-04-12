package sdk

import (
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/truthvalidator/TruthValidatorAgent/VerifyProposalAIAgent/contracts/TruthValidatorSentientNet"
)

type TruthValidatorClient struct {
	contract *TruthValidatorSentientNet.TruthValidatorSentientNet
	client   *ethclient.Client
}

// NewClient creates a new TruthValidator client
func NewClient(cfg *Config) (*TruthValidatorClient, error) {
	// Validate configuration
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	// Connect to Ethereum node
	client, err := ethclient.Dial(cfg.RPCURL)
	if err != nil {
		return nil, err
	}

	// Instantiate contract
	contract, err := TruthValidatorSentientNet.NewTruthValidatorSentientNet(cfg.ContractAddress, client)
	if err != nil {
		return nil, err
	}

	return &TruthValidatorClient{
		contract: contract,
		client:   client,
	}, nil
}

// Close closes the client connection
func (c *TruthValidatorClient) Close() {
	if c.client != nil {
		c.client.Close()
	}
}
