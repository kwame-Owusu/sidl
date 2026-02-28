package cmd

import (
	"fmt"
	"strings"

	"github.com/kwame-Owusu/sidl/internal"
	"github.com/kwame-Owusu/sidl/internal/tui"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info [sid or prefix]",
	Short: "Get information about a SID or SID prefix",
	Long:  "Get detailed information about a Twilio SID or its 2-character prefix.",
	Args:  cobra.ExactArgs(1),
	Run:   runInfo,
}

func runInfo(cmd *cobra.Command, args []string) {
	input := strings.ToUpper(args[0])

	// Case 1: Prefix only (2 characters)
	if len(input) == 2 {
		if description, ok := sids[input]; ok {
			runTUI(description.Name, description.Description, input)
			return
		}
		fmt.Printf("SID prefix %s not found\n", input)
		return
	}

	// Case 2: Full SID
	if len(input) < 2 {
		fmt.Println("Invalid SID input")
		return
	}

	prefix := input[0:2]

	isValidSid, err := internal.IsValidSid(input, sids)
	if err != nil {
		fmt.Printf("Error validating SID: %s\n", err)
		return
	}

	if !isValidSid {
		fmt.Printf("Unknown SID prefix: %s\n", prefix)
		return
	}

	if description, ok := sids[prefix]; ok {
		runTUI(description.Name, description.Description, prefix)
	} else {
		fmt.Printf("SID prefix %s not found\n", prefix)
	}
}

func runTUI(name, description, prefix string) {
	output := tui.RenderInfo(prefix, description, description)
	fmt.Println(output)
}
