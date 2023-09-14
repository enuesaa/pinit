package cli

import (
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/spf13/cobra"
	"github.com/erikgeiser/promptkit/textinput"
)

func CreateNewCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use: "new",
		Run: func(cmd *cobra.Command, args []string) {
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
