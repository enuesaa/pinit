package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/spf13/cobra"
)

func CreateServeCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "serve",
		Short: "serve admin console",
		Run: func(cmd *cobra.Command, args []string) {
			adminSrv := service.NewAdminService(repos)

			fmt.Printf("Admin console listening on http://localhost:3000 \n")
			adminSrv.Serve()
		},
	}

	return cmd
}
