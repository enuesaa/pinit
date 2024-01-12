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
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return usecase.LoadConfig(repos)
		},
		Run: func(cmd *cobra.Command, args []string) {
			migrate, _ := cmd.Flags().GetBool("migrate")

			if err := usecase.Configure(repos); err != nil {
				log.Fatalf("Error: %s", err.Error())
			}

			if migrate {
				fmt.Printf("Migration started.\n")
				if err := usecase.Migrate(repos); err != nil {
					log.Fatalf("Error: %s", err.Error())
				}
				fmt.Printf("Migration succeeded.\n")
			}

			fmt.Printf("Checking table existence..\n")
			if err := usecase.OpenDbConnection(repos); err != nil {
				log.Fatalf("Error: %s", err.Error())
			}
			if err := usecase.CheckTableStatus(repos); err != nil {
				log.Fatalf("Error: %s", err.Error())
			}
			fmt.Printf("All table exists.\n")
		},
	}
	cmd.Flags().Bool("migrate", false, "run migration")

	return cmd
}
