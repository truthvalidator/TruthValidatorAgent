package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

const (
	// NetworkFlag is the name of the flag used to specify the network.
	NetworkFlag = "network"
)

var (
	// RootCmd is the vibrant heart of TruthValidatorSentientNet - your gateway to decentralized AI governance.
	// It orchestrates all subcommands to verify, vote on, and manage AI proposals in a trustless environment.
	RootCmd = &cobra.Command{
		Use:   "TruthValidatorSentientNet",
		Short: "🚀 SovereignAISearch - Take control of AI truth verification in a decentralized network",
		Long: `A next-gen tool for decentralized AI governance that lets you:
- Verify AI proposal authenticity
- Vote on proposal validity
- Manage decentralized AI search results
- Ensure tamper-proof storage and privacy`,
	}
)

func init() {
	// Add a persistent flag to specify the network.
	RootCmd.PersistentFlags().String(NetworkFlag, "eth-localnet", "Name of the network to connect to")
}

// ExitWithError provides a friendly yet firm goodbye when things go wrong.
// It prints a clear error message to stderr before exiting with status 1.
// Example:
//   ExitWithError("Network connection failed", err)
func ExitWithError(msg string, err error) {
	fmt.Fprintf(os.Stderr, "ERROR: %s: %v\n", msg, err)
	os.Exit(1)
}

// GetNetworkAddress is your network navigator - it translates human-friendly network names
// into actual connection endpoints. Supports:
// - Sapphire mainnet/testnet
// - Ethereum localnet
// - Sepolia testnet
// Returns the dial address or exits gracefully if network is unknown.
func GetNetworkAddress() string {
	networks := map[string]string{
		"sapphire":          "https://sapphire.oasis.io",
		"sapphire-testnet":  "https://testnet.sapphire.oasis.io",
		"sapphire-localnet": "ws://localhost:8546",
		"eth-localnet":      "ws://localhost:8545",
		"sepolia":           "wss://sepolia.infura.io/v3/YOUR_INFURA_PROJECT_ID", // Replace with your Infura Project ID or other Sepolia provider
		"sepolia-testnet":   "wss://sepolia.infura.io/v3/YOUR_INFURA_PROJECT_ID",
	}

	net, err := RootCmd.PersistentFlags().GetString(NetworkFlag)
	if err != nil {
		ExitWithError("GetNetworkAddress", fmt.Errorf("Please specify the network to connect to using --%s.", NetworkFlag))
	}
	net = strings.ToLower(net)

	addr, found := networks[net]
	if !found {
		validNets := make([]string, 0, len(networks))
		for n := range networks {
			validNets = append(validNets, n)
		}

		ExitWithError("GetNetworkAddress", fmt.Errorf("Unknown network specified, please use one of the following: %s.", strings.Join(validNets, ", ")))
	}

	// Inform the user about using Infura and needing to set up an environment variable.
	if net == "sepolia" || net == "sepolia-testnet" {
		if networks[net] == "wss://sepolia.infura.io/v3/YOUR_INFURA_PROJECT_ID" {
			// Consider using environment variable or command line flag to get Infura ID more dynamically.
			// For simplicity, I'm keeping it in the map for now, but recommend better handling.
		}
		// If you want to dynamically get INFURA_PROJECT_ID from environment:
		// infuraID := os.Getenv("INFURA_PROJECT_ID")
		// if infuraID != "" {
		// 	addr = strings.Replace(addr, "YOUR_INFURA_PROJECT_ID", infuraID, 1)
		// } else {
		// 	fmt.Println("WARNING: INFURA_PROJECT_ID environment variable is not set. Please set it for Sepolia network.")
		// }

	}

	if net == "eth-localnet" {
		if networks[net] == "ws://localhost:8545" {
			// fmt.Println("WARNING: eth-localnet")
			// Consider using environment variable or command line flag to get Infura ID more dynamically.
			// For simplicity, I'm keeping it in the map for now, but recommend better handling.
		}
		// If you want to dynamically get INFURA_PROJECT_ID from environment:
		// infuraID := os.Getenv("INFURA_PROJECT_ID")
		// if infuraID != "" {
		// 	addr = strings.Replace(addr, "YOUR_INFURA_PROJECT_ID", infuraID, 1)
		// } else {
		// 	fmt.Println("WARNING: INFURA_PROJECT_ID environment variable is not set. Please set it for Sepolia network.")
		// }

	}

	return addr
}

// ParseAddress is your Ethereum address validator - it takes a hex string and
// converts it to a proper common.Address. Handles both 0x-prefixed and raw hex.
// Returns error if address format is invalid (not 40 chars after 0x removal).
// Example:
//   addr, err := ParseAddress("0x123...abc")
func ParseAddress(addrHex string) (common.Address, error) {
	if strings.HasPrefix(addrHex, "0x") {
		addrHex = strings.TrimPrefix(addrHex, "0x")
	}

	if len(addrHex) != 40 {
		return common.Address{}, fmt.Errorf("address is malformed")
	}

	return common.HexToAddress(addrHex), nil
}

// Execute brings TruthValidatorSentientNet to life! It:
// 1. Parses command line arguments
// 2. Routes to appropriate subcommands
// 3. Handles errors gracefully
// This is the main entry point that makes the magic happen.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		ExitWithError("TruthValidatorSentientNet", err)
	}
}
