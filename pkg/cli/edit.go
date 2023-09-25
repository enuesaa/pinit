package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/erikgeiser/promptkit/textinput"
	"github.com/spf13/cobra"
)

func CreateEditCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:  "edit <name>",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			noteSrv := service.NewNoteService(repos)
			note, err := noteSrv.Get(name)
			if err != nil {
				fmt.Println(err)
				return
			}

			newName, err := textinput.New("Name").RunPrompt()
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
			note.Name = newName
			note.Content = content
			note.Comment = comment
			noteSrv.Update(*note)

		},
	}

	return cmd
}
