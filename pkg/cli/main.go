package cli

import (
	"github.com/spf13/cobra"
)

func createRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "pinit",
		Short: "A CLI tool to register useful files and make project setup easier.",
		Args: cobra.MinimumNArgs(0),
		Run:  func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	return cmd
}

func CreateCli() *cobra.Command {
	cli := createRootCmd()
	cli.AddCommand(createApplyCmd())
	cli.AddCommand(createListCmd())
	cli.AddCommand(createRegisterCmd())
	cli.AddCommand(createRemoveCmd())

	// global options
	cli.PersistentFlags().Bool("register", false, "Register file.")
	cli.PersistentFlags().Bool("remove", false,  "Remove file.")
	cli.PersistentFlags().String("tag", "", "Tag")

	// disable default
	cli.SetHelpCommand(&cobra.Command{Hidden: true})
	cli.PersistentFlags().BoolP("help", "", false, "Show help information")
	cli.CompletionOptions.DisableDefaultCmd = true

	return cli
}
 