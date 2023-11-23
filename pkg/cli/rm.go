package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/spf13/cobra"
)

func CreateRmCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "rm <name>",
		Short: "Remove a binder",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			configSrv := service.NewConfigSevice(repos)
			if err := configSrv.Init(); err != nil {
				fmt.Printf("Error: %s", err.Error())
				return
			}

			name := args[0]
			binderSrv := service.NewBinderService(repos)
			binderSrv.Delete(name)
		},
	}

	return cmd
}
