package cli

import (
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/spf13/cobra"
)

func CreateRmCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:  "rm <name>",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			configSrv := service.NewConfigSevice(repos)
			if err := configSrv.ConfigureDatabaseDsn(); err != nil {
				return
			}

			name := args[0]
			noteSrv := service.NewNoteService(repos)
			noteSrv.Delete(name)
		},
	}

	return cmd
}
