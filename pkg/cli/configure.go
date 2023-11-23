package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/spf13/cobra"
)

func CreateConfigureCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "configure",
		Short: "setup pinit.",
		Run: func(cmd *cobra.Command, args []string) {
			migrate, _ := cmd.Flags().GetBool("migrate")

			configSrv := service.NewConfigSevice(repos)

			if configSrv.IsConfigExist() {
				config, err := configSrv.Read()
				if err != nil {
					fmt.Printf("Error: %s\n", err.Error())
					return
				}
				config, err = configSrv.RunEditPrompt(*config)
				if err != nil {
					fmt.Printf("Error: %s\n", err.Error())
					return
				}
				if err := configSrv.Write(*config); err != nil {
					fmt.Printf("error: %s\n", err)
				}
			} else {
				// create config file.
				config, err := configSrv.RunCreatePrompt()
				if err != nil {
					fmt.Printf("Error: %s\n", err.Error())
					return
				}
				if err := configSrv.Write(*config); err != nil {
					fmt.Printf("error: %s\n", err)
				}
			}

			if err := configSrv.Init(); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return
			}

			fmt.Printf("\n")
			fmt.Printf("Checking table exists...\n")
			isTableExist, err := configSrv.IsTableExist()
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return
			}
			if isTableExist {
				fmt.Printf("Table exists.\n")
			} else {
				fmt.Printf("Table DOES NOT exist.\n")

				if migrate {
					fmt.Printf("Migrating..\n")
					configSrv.Migration()
					fmt.Printf("Migration succeeded.\n")
				}
			}
		},
	}
	cmd.Flags().Bool("migrate", false, "do migration")


	return cmd
}
