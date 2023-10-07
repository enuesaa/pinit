package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/spf13/cobra"
)

func CreateAddCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use: "add",
		Short: "create a note",
		Run: func(cmd *cobra.Command, args []string) {
			configSrv := service.NewConfigSevice(repos)
			if err := configSrv.ConfigureDatabaseDsn(); err != nil {
				return
			}

			noteSrv := service.NewNoteService(repos)
			note, err := noteSrv.Prompt()
			if err != nil {
				fmt.Printf("Error: %s", err.Error())
			}

			if err := noteSrv.Create(*note); err != nil {
				fmt.Printf("Error: %s", err.Error())
			}
		},
	}

	return cmd
}
