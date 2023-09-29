package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/erikgeiser/promptkit/textinput"
	"github.com/spf13/cobra"
)

func CreateSetupConfigureCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use: "configure",
		Run: func(cmd *cobra.Command, args []string) {
			configSrv := service.NewConfigSevice(repos)

			databaseDsn, err := textinput.New("Database DSN").RunPrompt()
			if err != nil {
				return
			}
			if err := configSrv.WriteDatabaseDsn(databaseDsn); err != nil {
				fmt.Println(err)
				return
			}
		},
	}
	cmd.Flags().String("database-dsn", "", "Using when migration flag provided.")

	return cmd
}
