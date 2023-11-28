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
		Run: func(cmd *cobra.Command, args []string) {
			if err := usecase.CreateWithPrompt(repos); err != nil {
				log.Printf("Error: %s", err.Error())
				return
			}
		},
	}

	return cmd
}
