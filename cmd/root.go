package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sidly",
	Short: "A CLI tool for twilio sids",
	Long:  `A CLI twilio to get twilio sids detailed information from`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to sidly! Use 'sidly help' to see available commands.")
	},
}

// Execute executes the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		// Exit with error code 1 if something goes wrong
		// os.Exit(1)  // optional
	}
}
