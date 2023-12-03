package cli

import (
	"log"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateServeCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "serve",
		Short: "Launch the web console",
		Run: func(cmd *cobra.Command, args []string) {
			port, _ := cmd.Flags().GetInt("port")

			if err := usecase.Serve(repos, port); err != nil {
				log.Printf("Error: %s\n", err.Error())
				return
			}
		},
	}
	cmd.Flags().Int("port", 3000, "port. Default: 3000")

	return cmd
}
