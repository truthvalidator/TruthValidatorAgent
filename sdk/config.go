package sdk

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const (
	// DefaultRPCURL default RPC node URL
	DefaultRPCURL = "http://localhost:8545"
)

// Config SDK configuration
type Config struct {
	// RPCURL Ethereum node RPC URL
	RPCURL string

	// ContractAddress contract address
	ContractAddress common.Address

	// ChainID chain ID
	ChainID *big.Int
}

// NewDefaultConfig creates default configuration
func NewDefaultConfig(contractAddress common.Address, chainID *big.Int) *Config {
	return &Config{
		RPCURL:         DefaultRPCURL,
		ContractAddress: contractAddress,
		ChainID:        chainID,
	}
}

// Validate validates configuration
func (c *Config) Validate() error {
	if c.RPCURL == "" {
		return ErrRPCConnection
	}
	if c.ContractAddress == (common.Address{}) {
		return ErrContractCall
	}
	if c.ChainID == nil {
		return ErrInvalidChainID
	}
	return nil
}
