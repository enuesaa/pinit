package main

import (
	"github.com/enuesaa/pinit/pkg/cli"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   "pinit",
		Short: "A CLI client for a personal note-taking application.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	// sub commands
	cmd.AddCommand(cli.CreateConfigureCmd())
	cmd.AddCommand(cli.CreateNewCmd())
	cmd.AddCommand(cli.CreateLsCmd())
	cmd.AddCommand(cli.CreateDescribeCmd())
	cmd.AddCommand(cli.CreateRmCmd())

	// disable default
	cmd.SetHelpCommand(&cobra.Command{Hidden: true})
	cmd.PersistentFlags().BoolP("help", "", false, "Show help information")
	cmd.CompletionOptions.DisableDefaultCmd = true

	cmd.Execute()
}
