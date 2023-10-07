package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/spf13/cobra"
)

func CreateSetupStatusCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use: "status",
		Short: "show setup status.",
		Run: func(cmd *cobra.Command, args []string) {
			configSrv := service.NewConfigSevice(repos)
			config, err := configSrv.Read()
			if err != nil {
				return
			}
			fmt.Printf("Database DSN: %s \n", config.DatabaseDsn)
		},
	}

	return cmd
}
