package cli

import (
	"fmt"
	"log"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateLookupCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "lookup <name>",
		Short: "describe a binder",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			binderName := args[0]
			binder, err := usecase.Describe(repos, binderName)
			if err != nil {
				log.Printf("Error: %s", err.Error())
				return
			}
			fmt.Printf("Name: %s\n", binder.Name)
			fmt.Printf("Category: %s\n", binder.Category)

			note, err := usecase.DescribeFirstNote(repos, binder.ID)
			if err != nil {
				log.Printf("Error: %s", err.Error())
				return
			}
			fmt.Printf("Note Content: %s\n", note.Content)
			fmt.Printf("Note Comment: %s\n", note.Comment)
			fmt.Printf("Note Publisher: %s\n", note.Publisher)
		},
	}

	return cmd
}
