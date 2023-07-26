package cli

import (
	"github.com/spf13/cobra"
)

func createListCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:  "list",
		Args: cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			// todo
		},
	}

	return cmd
}
