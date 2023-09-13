package cli

import (
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/spf13/cobra"
)

func CreateNewCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use: "new",
		Run: func(cmd *cobra.Command, args []string) {
			noteSrv := service.NewNoteService(repos)
			noteSrv.Create(service.Note{
				Name:    "a",
				Content: "b",
				Comment: "c",
			})
			// create tags
		},
	}

	return cmd
}
