package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/spf13/cobra"
)

func CreateActionCreateCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "create an action",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("create an action\n")
		},
	}

	return cmd
}
