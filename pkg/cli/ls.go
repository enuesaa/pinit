package cli

import (
	"github.com/spf13/cobra"
)

func createLsCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "ls",
		Run: func(cmd *cobra.Command, args []string) {
			// list notes
		},
	}
	cmd.Flags().StringArrayP("filter", "", []string{}, "search vavlue")

	return cmd
}
