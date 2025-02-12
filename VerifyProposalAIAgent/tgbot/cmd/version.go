package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version = "unknown"
	GitHash = "unknown"
	BuildAt = "unknown"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "View version information",
	Long:  `View TruthValidatorSentientNet-tgbot version information, including version number, build time, and Git commit hash.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("")
		fmt.Println("AppName: ", "TruthValidatorSentientNet-tgbot")
		fmt.Println("Version: ", Version)
		fmt.Println("BuildAt: ", BuildAt)
		fmt.Println("GitHash: ", GitHash)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
