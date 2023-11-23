package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/usecase"
	"github.com/spf13/cobra"
)

func CreateNewCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "new",
		Short: "Create new binder and write note",
		Run: func(cmd *cobra.Command, args []string) {
			if err := usecase.CreateBinderWithPrompt(repos); err != nil {
				fmt.Printf("Error: %s", err.Error())
				return
			}
		},
	}

	return cmd
}
