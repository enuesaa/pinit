package cli

import (
	"github.com/enuesaa/pinit/pkg/handler"
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/spf13/cobra"
)

func createRemoveCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:  "remove",
		Args: cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			input := ParseArgs(cmd, args)

			itemsRepo := repository.NewItemsRepository()
			handler.HandleList(itemsRepo, input.Tag)
		},
	}

	return cmd
}
