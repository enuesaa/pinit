package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/spf13/cobra"
)

func CreateUpCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "up",
		Short: "Serve web console",
		Run: func(cmd *cobra.Command, args []string) {
			adminSrv := service.NewAdminService(repos)

			fmt.Printf("Admin console listening on http://localhost:3000 \n")
			adminSrv.Serve()
		},
	}

	return cmd
}
