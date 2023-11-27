package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateWriteCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "write <name>",
		Short: "write a note",
		Args: cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			binderName := args[0]
			if err := usecase.WriteNote(repos, binderName); err != nil {
				fmt.Printf("Error: %s", err.Error())
				return
			}
		},
	}

	return cmd
}
