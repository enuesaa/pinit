package cli

import (
	"fmt"
	"log"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateConfigureCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "configure",
		Short: "Setup pinit",
		Run: func(cmd *cobra.Command, args []string) {
			migrate, _ := cmd.Flags().GetBool("migrate")

			if err := usecase.Configure(repos); err != nil {
				log.Printf("Error: %s\n", err.Error())
				return
			}

			if migrate {
				fmt.Printf("Migrating..\n")
				if err := usecase.Migrate(repos); err != nil {
					log.Printf("Error: %s\n", err.Error())
					return
				}
				fmt.Printf("Migration succeeded.\n")
			}

			fmt.Printf("Checking table status..\n")
			if err := usecase.CheckTableStatus(repos); err != nil {
				log.Printf("Error: %s\n", err.Error())
				return
			}
			fmt.Printf("Binders table exists.\n")
			fmt.Printf("Notes table exists.\n")
		},
	}
	cmd.Flags().Bool("migrate", false, "do migration")

	return cmd
}
