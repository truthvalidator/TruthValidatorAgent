package sdk

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Wallet encapsulates Ethereum wallet operations
type Wallet struct {
	privateKey *ecdsa.PrivateKey
	address    common.Address
}

// NewWalletFromPrivateKey creates a wallet from a private key
func NewWalletFromPrivateKey(privateKeyHex string) (*Wallet, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &Wallet{
		privateKey: privateKey,
		address:    address,
	}, nil
}

// GetAddress gets the wallet address
func (w *Wallet) GetAddress() common.Address {
	return w.address
}

// NewTransactor creates a transaction signer
func (w *Wallet) NewTransactor(chainID *big.Int) (*bind.TransactOpts, error) {
	return bind.NewKeyedTransactorWithChainID(w.privateKey, chainID)
}
