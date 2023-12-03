package cli

import (
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/spf13/cobra"
)

func CreateActionCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "action",
		Short: "manage actions",
	}
	cmd.AddCommand(CreateActionCreateCmd(repos))

	return cmd
}
