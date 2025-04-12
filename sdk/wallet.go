package sdk

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Wallet 封装以太坊钱包操作
type Wallet struct {
	privateKey *ecdsa.PrivateKey
	address    common.Address
}

// NewWalletFromPrivateKey 从私钥创建钱包
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

// GetAddress 获取钱包地址
func (w *Wallet) GetAddress() common.Address {
	return w.address
}

// NewTransactor 创建交易签名者
func (w *Wallet) NewTransactor(chainID *big.Int) (*bind.TransactOpts, error) {
	return bind.NewKeyedTransactorWithChainID(w.privateKey, chainID)
}
