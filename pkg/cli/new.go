package cli

import (
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/erikgeiser/promptkit/textinput"
	"github.com/spf13/cobra"
)

func CreateNewCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use: "new",
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
			noteSrv.Create(service.Note{
				Name:    name,
				Content: content,
				Comment: comment,
			})
		},
	}

	return cmd
}
