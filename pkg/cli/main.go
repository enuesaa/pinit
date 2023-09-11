package cli

import (
	"github.com/spf13/cobra"
)

func createRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pinit",
		Short: "A CLI client for a personal note-taking application.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return cmd
}

func CreateCli() *cobra.Command {
	cli := createRootCmd()
	cli.AddCommand(createConfigureCmd())
	cli.AddCommand(createNewCmd())
	cli.AddCommand(createLsCmd())
	cli.AddCommand(createDescribeCmd())
	cli.AddCommand(createRmCmd())

	// disable default
	cli.SetHelpCommand(&cobra.Command{Hidden: true})
	cli.PersistentFlags().BoolP("help", "", false, "Show help information")
	cli.CompletionOptions.DisableDefaultCmd = true

	return cli
}
