package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/erikgeiser/promptkit/textinput"
	"github.com/spf13/cobra"
)

func createConfigureCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "configure",
		Run: func(cmd *cobra.Command, args []string) {
			repos := repository.NewRepos()
			configSrv := service.NewConfigSevice(repos)

			isReadTask, _ := cmd.Flags().GetBool("read")
			if isReadTask {
				config, err := configSrv.Read()
				if err != nil {
					return
				}
				fmt.Printf("Database DSN: %s \n", config.DatabaseDsn)
			} else {
				config := service.Config{
					DatabaseDsn: "",
				}

				databaseDsn, err := textinput.New("Database DSN").RunPrompt()
				if err != nil {
					return
				}
				config.DatabaseDsn = databaseDsn

				if err := configSrv.Write(config); err != nil {
					return
				}
			}
		},
	}

	cmd.Flags().Bool("read", false, "Read Config")

	return cmd
}
