package cli

import (
	"github.com/spf13/cobra"
)

func createDescribeCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:  "describe <name>",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	return cmd
}
