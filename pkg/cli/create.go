package cli

import (
	"log"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateCreateCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "Create new binder and write a note",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return usecase.OnStartUp(repos)
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return usecase.CloseDbConnection(repos)
		},
		Run: func(cmd *cobra.Command, args []string) {
			if err := usecase.RunCreatePrompt(repos); err != nil {
				log.Fatalf("Error: %s", err.Error())
			}
		},
	}

	return cmd
}
