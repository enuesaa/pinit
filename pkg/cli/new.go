package cli

import (
	"fmt"

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
			_, err := configSrv.Read()
			if err != nil {
				fmt.Printf("%+v", err)
				return
			}

			// dbRepo := repository.DatabaseRepository{
			// 	Dsn: config.DatabaseDsn,
			// }

			// create note
			// dbRepo.Create(&service.Note{
			// 	Name:    "a",
			// 	Content: "b",
			// 	Comment: "c",
			// })
			// create tags
		},
	}

	return cmd
}
