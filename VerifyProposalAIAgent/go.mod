// TruthValidator Agent Go Module
//
// This module defines the dependencies for the TruthValidator Agent,
// which provides decentralized proposal validation and voting functionality.
//
// Key Dependencies:
// - go-ethereum: Ethereum client library
// - sapphire-paratime: Oasis Sapphire compatibility
// - cobra: CLI framework
// - zap: Logging
//
// Go version: 1.22.5 (minimum required)
module github.com/truthvalidator/TruthValidatorAgent/VerifyProposalAIAgent

go 1.23.0 // Minimum Go version

toolchain go1.24.1

require (
	github.com/ethereum/go-ethereum v1.15.10
	github.com/go-telegram/bot v1.15.0
	github.com/joho/godotenv v1.5.1
	github.com/mnt-ltd/daemore v0.0.0-20240227053424-d57afb9f5e67
	github.com/spf13/cobra v1.9.1
	go.uber.org/zap v1.27.0
)

replace github.com/cometbft/cometbft => github.com/oasisprotocol/cometbft v0.37.2-oasis1

require (
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/bits-and-blooms/bitset v1.17.0 // indirect
	github.com/consensys/bavard v0.1.22 // indirect
	github.com/consensys/gnark-crypto v0.14.0 // indirect
	github.com/crate-crypto/go-ipa v0.0.0-20240724233137-53bbb0ceb27a // indirect
	github.com/crate-crypto/go-kzg-4844 v1.1.0 // indirect
	github.com/deckarep/golang-set/v2 v2.6.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.2.0 // indirect
	github.com/ethereum/c-kzg-4844 v1.0.0 // indirect
	github.com/ethereum/go-verkle v0.2.2 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/google/uuid v1.4.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/holiman/uint256 v1.3.2 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/kardianos/service v1.2.2 // indirect
	github.com/klauspost/compress v1.17.2 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mmcloughlin/addchain v0.4.0 // indirect
	github.com/prometheus/client_golang v1.17.0 // indirect
	github.com/rs/cors v1.11.0 // indirect
	github.com/shirou/gopsutil v3.21.4-0.20210419000835-c7a38de76ee5+incompatible // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	github.com/supranational/blst v0.3.14 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.35.0 // indirect
	golang.org/x/exp v0.0.0-20240222234643-814bf88cf225 // indirect
	golang.org/x/sync v0.11.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	rsc.io/tmplfunc v0.0.3 // indirect
)
