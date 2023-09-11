package cli

import (
	"github.com/spf13/cobra"
)

func createRmCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:  "rm <name>",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// remove
		},
	}

	return cmd
}
