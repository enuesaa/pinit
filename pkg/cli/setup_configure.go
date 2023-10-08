package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/spf13/cobra"
)

func CreateSetupConfigureCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use: "configure",
		Short: "setup db connection.",
		Run: func(cmd *cobra.Command, args []string) {
			databaseDsn, err := cmd.Flags().GetString("database-dsn")
			if databaseDsn == "" || err != nil {
				fmt.Printf("Error: flag `--database-dsn` should not be empty.\n")
				return
			}

			config := service.Config {
				DatabaseDsn: databaseDsn,
			}
			configSrv := service.NewConfigSevice(repos)
			if err := configSrv.Write(config); err != nil {
				fmt.Printf("error: %s\n", err)
			}
		},
	}
	cmd.Flags().String("database-dsn", "", "[Required] Database Dsn")
	cmd.MarkFlagRequired("database-dsn")

	return cmd
}
