package cli

import (
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/spf13/cobra"
)

func CreateSetupCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use: "setup",
		Short: "setup to take notes.",
	}
	cmd.AddCommand(CreateSetupStatusCmd(repos))
	cmd.AddCommand(CreateSetupConfigureCmd(repos))
	cmd.AddCommand(CreateSetupMigrationCmd(repos))

	return cmd
}
