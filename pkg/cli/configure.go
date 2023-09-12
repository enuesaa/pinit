package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/erikgeiser/promptkit/textinput"
	"github.com/spf13/cobra"
)

func CreateConfigureCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "configure",
		Run: func(cmd *cobra.Command, args []string) {
			input := parseConfigureInput(cmd)

			repos := repository.NewRepos()
			configSrv := service.NewConfigSevice(repos)

			if input.read {
				config, err := configSrv.Read()
				if err != nil {
					return
				}
				fmt.Printf("Database DSN: %s \n", config.DatabaseDsn)
			} else {
				databaseDsn, err := textinput.New("Database DSN").RunPrompt()
				if err != nil {
					return
				}
				if err := configSrv.WriteDatabaseDsn(databaseDsn); err != nil {
					return
				}
			}
		},
	}
	cmd.Flags().Bool("read", false, "Read Config")

	return cmd
}


type ConfigureInput struct {
	read bool
}
func parseConfigureInput(cmd *cobra.Command) ConfigureInput {
	isReadTask, _ := cmd.Flags().GetBool("read")

	return ConfigureInput{
		read: isReadTask,
	}
}
