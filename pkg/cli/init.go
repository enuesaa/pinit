package cli

import (
	"log"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateInitCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize pinit app. Run database migration and setup application.",
		Run: func(cmd *cobra.Command, args []string) {
			if err := usecase.CreateRegistryIfNotExist(repos); err != nil {
				log.Fatalf("Error: %s", err.Error())
			}
			if err := usecase.OpenDb(repos); err != nil {
				log.Fatalf("Error: %s", err.Error())
			}
			if err := usecase.ConfigureAppConfig(repos); err != nil {
				log.Fatalf("Error: %s", err.Error())
			}
		},
	}

	return cmd
}
