package cli

import (
	"github.com/enuesaa/pinit/pkg/handler"
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/spf13/cobra"
)

func createApplyCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:  "apply",
		Args: cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			input := ParseArgs(cmd, args)

			itemsRepo := repository.NewItemsRepository()
			handler.HandleApply(itemsRepo, input.Tag, input.Filename)
		},
	}

	return cmd
}
