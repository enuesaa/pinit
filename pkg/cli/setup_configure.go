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
				fmt.Printf("error: Missing required flag `--database-dsn`.\n")
				return
			}
			configSrv := service.NewConfigSevice(repos)
			if err := configSrv.WriteDatabaseDsn(databaseDsn); err != nil {
				fmt.Printf("error: %s\n", err)
				return
			}
		},
	}
	cmd.Flags().String("database-dsn", "", "[Required] Database Dsn")

	return cmd
}
