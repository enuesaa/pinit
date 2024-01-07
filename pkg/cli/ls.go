package cli

import (
	"fmt"
	"log"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateLsCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "ls",
		Short: "List binders",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return usecase.OnStartUp(repos)
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return usecase.CloseDbConnection(repos)
		},
		Run: func(cmd *cobra.Command, args []string) {
			binders, err := usecase.ListBinders(repos)
			if err != nil {
				log.Fatalf("Error: %s", err.Error())
			}

			fmt.Printf("%d binder(s) found.\n", len(binders))
			if len(binders) == 0 {
				return
			}

			for _, binder := range binders {
				note, err := usecase.DescribeLatestNote(repos, binder.Name)
				latestContent := "<nocontent>"
				if err == nil {
					latestContent = note.Content
				}
				fmt.Printf("%s: %s\n", binder.Name, latestContent)
			}
		},
	}

	return cmd
}
