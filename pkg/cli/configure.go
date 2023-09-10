package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/erikgeiser/promptkit/textinput"
	"github.com/spf13/cobra"
)

func createConfigureCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:  "configure",
		Run: func(cmd *cobra.Command, args []string) {
			config := repository.Config {
				DatabaseDsn: "",
			}

			databaseDsn, err := textinput.New("Database DSN").RunPrompt()
			if err != nil {
				return
			}
			config.DatabaseDsn = databaseDsn

			repo := repository.ConfigRepository{}
			if err := repo.WriteConfig(config); err != nil {
				return
			}

			savedConfig, err := repo.ReadConfig() 
			if err != nil {
				return
			}
			fmt.Printf("save completed. %+v", savedConfig)
		},
	}

	return cmd
}
