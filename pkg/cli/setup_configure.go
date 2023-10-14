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
		Short: "setup pinit.",
		Run: func(cmd *cobra.Command, args []string) {
			configSrv := service.NewConfigSevice(repos)
			config, err := configSrv.Read()
			if err != nil {
				// create config file.
				config, err = configSrv.RunCreatePrompt()
				if err != nil {
					fmt.Printf("Error: %s\n", err.Error())
					return
				}
			} else {
				config, err = configSrv.RunEditPrompt(*config)
				if err != nil {
					fmt.Printf("Error: %s\n", err.Error())
					return
				}
			}

			if err := configSrv.Write(*config); err != nil {
				fmt.Printf("error: %s\n", err)
			}
		},
	}

	return cmd
}
