package cli

import (
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateRmCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "rm <name>",
		Short: "Remove a binder",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			binderName := args[0]
			usecase.DeleteBinder(repos, binderName)
		},
	}

	return cmd
}
