package cli

import (
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/spf13/cobra"
)

func CreateLsCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "ls",
		Run: func(cmd *cobra.Command, args []string) {
			repos := repository.NewRepos()
			configSrv := service.NewConfigSevice(repos)
			databaseDsn, err := configSrv.ReadDatabaseDsn()
			if err != nil {
				return
			}
			repos.Database.WithDsn(databaseDsn)
		},
	}
	cmd.Flags().StringArrayP("filter", "", []string{}, "search vavlue")

	return cmd
}
