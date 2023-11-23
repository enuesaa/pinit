package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/enuesaa/pinit/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateConfigureCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "configure",
		Short: "Setup pinit",
		Run: func(cmd *cobra.Command, args []string) {
			migrate, _ := cmd.Flags().GetBool("migrate")

			corecase := usecase.NewCorecase()
			if err := corecase.Configure(repos); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				return
			}

			configSrv := service.NewConfigSevice(repos)

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
					if err := corecase.Migrate(repos); err != nil {
						fmt.Printf("Error: %s\n", err.Error())
					}
					fmt.Printf("Migration succeeded.\n")
				}
			}
		},
	}
	cmd.Flags().Bool("migrate", false, "do migration")


	return cmd
}
