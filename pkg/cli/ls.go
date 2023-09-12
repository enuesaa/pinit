package cli

import (
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/spf13/cobra"
)

func CreateLsCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use: "ls",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	cmd.Flags().StringArrayP("filter", "", []string{}, "search vavlue")

	return cmd
}
