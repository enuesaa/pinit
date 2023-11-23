package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/spf13/cobra"
)

func CreateNewCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "new",
		Short: "Create new binder and write note",
		Run: func(cmd *cobra.Command, args []string) {
			configSrv := service.NewConfigSevice(repos)
			if err := configSrv.Init(); err != nil {
				fmt.Printf("Error: %s", err.Error())
				return
			}

			noteSrv := service.NewNoteService(repos)
			note, err := noteSrv.RunCreatePrompt()
			if err != nil {
				fmt.Printf("Error: %s", err.Error())
				return
			}

			if err := noteSrv.Create(*note); err != nil {
				fmt.Printf("Error: %s", err.Error())
			}
		},
	}

	return cmd
}
