package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/spf13/cobra"
)

func CreateServeCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "serve",
		Short: "serve admin console",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("admin console")
		},
	}

	return cmd
}
