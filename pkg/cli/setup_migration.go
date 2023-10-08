package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/spf13/cobra"
)

func CreateSetupMigrationCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use: "migration",
		Short: "migration tables.",
		Run: func(cmd *cobra.Command, args []string) {
			configSrv := service.NewConfigSevice(repos)
			if err := configSrv.Init(); err != nil {
				fmt.Printf("Error: %s", err.Error())
				return
			}

			isTableExist, err := repos.Database.IsTableExist(&service.Note{})
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				return
			}
			fmt.Printf("isTableExist: %t\n", isTableExist)
			if isTableExist {
				fmt.Printf("Error: Migration aborted because note table exists.\n")
			}
			
			repos.Database.CreateTable(&service.Note{})
			fmt.Printf("note table created.\n")	
		},
	}

	return cmd
}
