// Package connection provides Ethereum network connectivity and transaction utilities
//
// This package handles:
// - Establishing secure connections to Ethereum networks
// - Managing cryptographic keys and wallet addresses
// - Preparing transactions with proper gas estimation
// - Network-specific configurations
//
// Example usage:
//   conn, err := connection.NewConnection(context.Background(), "https://mainnet.infura.io/v3/YOUR-PROJECT-ID")
//   if err != nil {
//       log.Fatal(err)
//   }
//   
//   txOpts, err := conn.PrepareNextTx(context.Background())
//   if err != nil {
//       log.Fatal(err)
//   }
//
// Version: 1.2.0
// Author: TruthValidator Team
package connection

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Connection represents a secure connection to an Ethereum network
// with wallet and transaction capabilities.
type Connection struct {
	// Client is the Go Ethereum client instance used for RPC calls.
	// Handles all network communication including:
	// - Querying blockchain state
	// - Submitting transactions
	// - Event subscription
	Client *ethclient.Client

	// ChainID identifies the Ethereum network (1=Mainnet, 5=Goerli, etc.)
	// Used for EIP-155 transaction replay protection
	ChainID *big.Int

	// PrivateKey is the ECDSA private key loaded from PRIVATE_KEY env var.
	// WARNING: Ensure proper key management in production environments.
	PrivateKey *ecdsa.PrivateKey

	// PublicKey is derived from PrivateKey using secp256k1 curve.
	// Used for address generation and signature verification.
	PublicKey *ecdsa.PublicKey

	// Address is the Ethereum wallet address derived from PublicKey.
	// Format: 0x followed by 20-byte hex string (40 characters)
	Address common.Address
}

// NewConnection establishes a new Ethereum or Filecoin FEVM network connection
//
// Parameters:
//   ctx - Context for cancellation and timeouts
//   where - Network RPC endpoint URL (e.g. "https://mainnet.infura.io" or FEVM endpoint)
//
// Returns:
//   *Connection - Configured connection instance
//   error - Any initialization error
//
// Example:
//   conn, err := NewConnection(context.Background(), "https://ropsten.infura.io/v3/YOUR-PROJECT-ID")
//   if err != nil {
//       return nil, fmt.Errorf("connection failed: %v", err)
//   }
//
// Notes:
// - Requires PRIVATE_KEY environment variable set
// - Automatically handles 0x prefix in private key
// - Validates network connectivity during initialization
func NewConnection(ctx context.Context, where string) (*Connection, error) {
	var (
		c   Connection
		err error
		ok  bool
	)

	c.Client, err = ethclient.Dial(where)
	if err != nil {
		return nil, err
	}

	// Special handling for Filecoin FEVM
	if strings.Contains(where, "filecoin") || strings.Contains(where, "fevm") {
		// FEVM uses Ethereum chain ID 314 for mainnet
		c.ChainID = big.NewInt(314)
	} else {
		c.ChainID, err = c.Client.NetworkID(ctx)
		if err != nil {
			return nil, fmt.Errorf("unable to get chain ID: %v", err)
		}
	}

	pk := os.Getenv("PRIVATE_KEY")
	if strings.HasPrefix(pk, "0x") {
		pk = strings.Replace(pk, "0x", "", 1)
	}

	c.PrivateKey, err = crypto.HexToECDSA(pk)
	if err != nil {
		return nil, fmt.Errorf("unable to get private key from environment: %v", err)
	}
	c.PublicKey, ok = c.PrivateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("public key is not in ECDSA format")
	}
	c.Address = crypto.PubkeyToAddress(*c.PublicKey)

	// Removed Sapphire wrapping
	// wc, err := sapphire.WrapClient(c.Client, func(digest [32]byte) ([]byte, error) {
	// 	return crypto.Sign(digest[:], c.PrivateKey)
	// })
	// if err != nil {
	// 	return nil, fmt.Errorf("unable to wrap backend: %v", err)
	// }
	// c.Sapphire = *wc
	// For standard eth, we don't need Sapphire wrapping, we can directly use Client
	// We can set Sapphire to nil, or remove it from struct if you don't need it at all in other parts of your code.
	// For now, let's just keep it nil to minimize changes if other parts of your code expect `c.Sapphire` to exist.
	// c.Sapphire = nil // Or remove Sapphire field from struct if not needed elsewhere

	return &c, nil
}

// PrepareNextTx prepares transaction options for the next outgoing transaction
//
// Handles:
// - Nonce management (auto-incrementing)
// - Gas price estimation (network suggested)
// - Gas limit setting (fixed at 30M gas)
// - Signer configuration
//
// Parameters:
//   ctx - Context for cancellation and timeouts
//
// Returns:
//   *bind.TransactOpts - Configured transaction options
//   error - Any preparation error
//
// Example:
//   txOpts, err := conn.PrepareNextTx(context.Background())
//   if err != nil {
//       return nil, fmt.Errorf("tx preparation failed: %v", err)
//   }
//   tx, err := contract.Method(txOpts, args...)
//
// Notes:
// - Gas limit may need adjustment for complex contracts
// - For EIP-1559 transactions, use SuggestGasPrice with type 2 transactions
func (c *Connection) PrepareNextTx(ctx context.Context) (*bind.TransactOpts, error) {
	nonce, err := c.Client.PendingNonceAt(ctx, c.Address) // Use c.Client directly
	if err != nil {
		return nil, fmt.Errorf("unable to get pending nonce: %v", err)
	}

	gasPrice, err := c.Client.SuggestGasPrice(ctx) // Use c.Client directly
	if err != nil {
		return nil, fmt.Errorf("unable to get suggested gas price: %v", err)
	}

	auth := bind.NewKeyedTransactor(c.PrivateKey) // Use standard bind.NewKeyedTransactor
	auth.Nonce = new(big.Int).SetUint64(nonce)
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(30000000) // You might want to adjust gas limit based on your needs

	return auth, nil
}
