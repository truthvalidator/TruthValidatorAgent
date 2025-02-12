package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	logger, _ = zap.NewProduction()
)

var rootCmd = &cobra.Command{
	Use:   "TruthValidatorSentientNet-tgbot",
	Short: "TruthValidatorSentientNet-tgbot, a Golang Bot for Telegram",
	Long:  `TruthValidatorSentientNet-tgbot is a Golang Bot for Telegram`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

const (
	// NetworkFlag is the name of the flag used to specify the network.
	NetworkFlag = "NETWORK"
)

func init() {

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// ExitWithError terminates the program after writing the error to stderr.
func ExitWithError(msg string, err error) {
	fmt.Fprintf(os.Stderr, "ERROR: %s: %v\n", msg, err)
	os.Exit(1)
}

// GetNetworkAddress returns the dial address of the network specified via the
// network flag. If the network is unknown, it exits with an error.
func GetNetworkAddress() string {
	networks := map[string]string{
		"sapphire":          "https://sapphire.oasis.io",
		"sapphire-testnet":  "https://testnet.sapphire.oasis.io",
		"sapphire-localnet": "ws://localhost:8546",
		"eth-localnet":      "ws://127.0.0.1:8545",
		"sepolia":           "https://sepolia.infura.io/v3/YOUR_INFURA_PROJECT_ID", // Replace with your Infura Project ID or other Sepolia provider
		"sepolia-testnet":   "wss://sepolia.infura.io/v3/YOUR_INFURA_PROJECT_ID",
	}

	net := os.Getenv(NetworkFlag)
	if net == "" {
		ExitWithError("GetNetworkAddress", fmt.Errorf("The NETWORK environment variable is not set. Please set it to specify the network to connect to (e.g., 'sapphire', 'sapphire-testnet', etc.)."))
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
			fmt.Println("WARNING: You are using the Sepolia network with a placeholder Infura Project ID.")
			fmt.Println("Please replace 'YOUR_INFURA_PROJECT_ID' in the code or set the INFURA_PROJECT_ID environment variable.")
			fmt.Println("For example, you can use `--network sepolia INFURA_PROJECT_ID=<your_project_id>`")
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
		if networks[net] == "ws://127.0.0.1:8545" {
			log.Println("GetNetworkAddress:", "eth-localnet", "ws://127.0.0.1:8545")
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

// ParseAddress converts the hex representation of an Ethereum address into
// common.Address. It returns an error if the address is malformed.
func ParseAddress(addrHex string) (common.Address, error) {
	if strings.HasPrefix(addrHex, "0x") {
		addrHex = strings.TrimPrefix(addrHex, "0x")
	}

	if len(addrHex) != 40 {
		return common.Address{}, fmt.Errorf("address is malformed")
	}

	return common.HexToAddress(addrHex), nil
}
