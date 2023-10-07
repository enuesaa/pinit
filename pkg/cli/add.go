package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/erikgeiser/promptkit/textinput"
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
			name, err := textinput.New("Name").RunPrompt()
			if err != nil {
				return
			}
			content, err := textinput.New("Content").RunPrompt()
			if err != nil {
				return
			}
			comment, err := textinput.New("Comment").RunPrompt()
			if err != nil {
				return
			}
			note := service.Note {
				Name:    name,
				Content: content,
				Comment: comment,
			}
			if err := noteSrv.Create(note); err != nil {
				fmt.Printf("Error: %s", err.Error())
			}
		},
	}

	return cmd
}
