package cli

import (
	"github.com/spf13/cobra"
)

func createRegisterCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:  "register",
		Args: cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			// input := ParseArgs(cmd, args)

			// itemsRepo := repository.NewItemsRepository()
			// handler.HandleRegister(itemsRepo, input.Tag, input.Filename)
		},
	}

	return cmd
}
