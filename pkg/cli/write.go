package cli

import (
	"log"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateWriteCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "write <name>",
		Short: "write a note",
		Args:  cobra.MinimumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := usecase.ConfigInit(repos); err != nil {
				return err
			}
			return usecase.OpenDbConnection(repos)
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return usecase.CloseDbConnection(repos)
		},
		Run: func(cmd *cobra.Command, args []string) {
			binderName := args[0]
			if err := usecase.WriteNewNoteWithPrompt(repos, binderName); err != nil {
				log.Printf("Error: %s", err.Error())
				return
			}
		},
	}

	return cmd
}
