package cmd

import (
	"fmt"

	"github.com/kwame-Owusu/sidl/internal"
	"github.com/spf13/cobra"
)

var sids map[string]internal.Field
var dataFile string = "sids.json"

func loadSids() {
	var err error
	sids, err = internal.LoadFile(dataFile)
	if err != nil {
		fmt.Printf("Error loading sids: %v\n", err)
		return
	}
}

var rootCmd = &cobra.Command{
	Use:   "sidl",
	Short: "A CLI tool for twilio sids",
	Long:  `A CLI tool to get detailed information about twilio sids`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to sidl! Use 'sidl help' to see available commands.")
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get version of sidl",
	Long:  `Get version of sidl`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("0.8.0")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(explainCmd)
	rootCmd.AddCommand(explainPrefixCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(typesCmd)
	rootCmd.AddCommand(tuiCmd)
	loadSids()
}

// Execute executes the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		// Exit with error code 1 if something goes wrong
		// os.Exit(1)  // optional
	}
}
