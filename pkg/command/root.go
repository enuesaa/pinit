package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/enuesaa/pinit/pkg/handler"
)

func createRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "pinit",
		Args: cobra.MinimumNArgs(0),
		Run:  func(cmd *cobra.Command, args []string) {
			input := ParseArgs(cmd, args)

			switch {
			case input.IsOperationAmbiguous():
				fmt.Printf("Error: Operation Ambiguous\n\n")
				fmt.Printf("Cannot use these flags at the same time: --register, --remove\n")
			case input.HasRegisterFlag():
				handler.HandleRegister()
			case input.HasRemoveFlag():
				handler.HandleRemove()
			default:
				handler.HandleList()
			}
		},
	}

	return cmd
}

func CreateCli() *cobra.Command {
	cli := createRootCmd()

	// global options
	cli.PersistentFlags().String("register", "", "Register file.")
	cli.PersistentFlags().String("remove", "", "Remove file.")
	cli.PersistentFlags().String("tag", "", "Tag")

	// disable default
	cli.SetHelpCommand(&cobra.Command{Hidden: true})
	cli.PersistentFlags().BoolP("help", "", false, "Show help information")
	cli.CompletionOptions.DisableDefaultCmd = true

	return cli
}
