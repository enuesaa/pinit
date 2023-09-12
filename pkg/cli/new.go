package cli

import (
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/spf13/cobra"
)

func createNewCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "new",
		Run: func(cmd *cobra.Command, args []string) {
			repos := repository.NewRepos()
			configSrv := service.NewConfigSevice(repos)
			databaseDsn, err := configSrv.ReadDatabaseDsn()
			if err != nil {
				return
			}
			repos.Database.WithDsn(databaseDsn)

			noteSrv := service.NoteService{}
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
