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
			repo := repository.ConfigRepository{}

			isReadTask, _ := cmd.Flags().GetBool("read")
			if isReadTask {
				config, err := repo.ReadConfig() 
				if err != nil {
					return
				}
				fmt.Printf("Database DSN: %s \n", config.DatabaseDsn)
			} else {
				config := repository.Config {
					DatabaseDsn: "",
				}
	
				databaseDsn, err := textinput.New("Database DSN").RunPrompt()
				if err != nil {
					return
				}
				config.DatabaseDsn = databaseDsn
	
				if err := repo.WriteConfig(config); err != nil {
					return
				}
			}
		},
	}

	cmd.Flags().Bool("read", false, "Read Config")

	return cmd
}
