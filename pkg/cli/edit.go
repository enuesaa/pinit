package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/spf13/cobra"
)

func CreateEditCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:  "edit <name>",
		Short: "edit a note.",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]

			configSrv := service.NewConfigSevice(repos)
			if err := configSrv.Init(); err != nil {
				fmt.Printf("Error: %s", err.Error())
				return
			}

			noteSrv := service.NewNoteService(repos)
			note, err := noteSrv.Get(name)
			if err != nil {
				fmt.Printf("Error: %s", err.Error())
				return
			}

			note, err = noteSrv.RunEditPrompt(*note)
			if err != nil {
				fmt.Printf("Error: %s", err.Error())
				return
			}
			if err := noteSrv.Update(*note); err != nil {
				fmt.Printf("Error: %s", err.Error())
			}
		},
	}

	return cmd
}
