package cli

import (
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateInitCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize pinit. Run database migration and setup pinit",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := repos.Db.Init(); err != nil {
				return err
			}
			if !usecase.DBIsExist(repos) {
				repos.Log.Info("database migration start")
				if err := usecase.DBSetup(repos); err != nil {
					return err
				}
				repos.Log.Info("database migration succeeded")
			}

			if err := usecase.DBOpen(repos); err != nil {
				return err
			}
			defer usecase.DBClose(repos)

			repos.Log.Info("app setup...")
			if err := usecase.ConfigAsk(repos); err != nil {
				return err
			}
			repos.Log.Info("succeeded")

			return nil
		},
	}

	return cmd
}
