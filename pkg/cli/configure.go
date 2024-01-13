package cli

import (
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
			if err := usecase.CreateRegistryIfNotExist(repos); err != nil {
				log.Fatalf("Error: %s", err.Error())
			}

			if err := usecase.OpenDb(repos); err != nil {
				log.Fatalf("Error: %s", err.Error())
			}

			// if err := usecase.Configure(repos); err != nil {
			// 	log.Fatalf("Error: %s", err.Error())
			// }
		},
	}

	return cmd
}
