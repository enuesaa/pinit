package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/handler"
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/spf13/cobra"
)

func createRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "pinit",
		Short: "A CLI tool to register useful files and make project setup easier.",
		Args: cobra.MinimumNArgs(0),
		Run:  func(cmd *cobra.Command, args []string) {
			input := ParseArgs(cmd, args)

			itemsRepo := repository.NewItemsRepository()

			switch {
			case input.IsOperationAmbiguous():
				fmt.Printf("Error: Operation Ambiguous\n\n")
				fmt.Printf("Cannot use these flags at the same time: --register, --remove\n")
			case input.Register:
				handler.HandleRegister(itemsRepo, input.Tag, input.Filename)
			case input.Remove:
				handler.HandleRemove(itemsRepo)
			case input.HasFilename():
				fmt.Printf("apply command here.")
			default:
				handler.HandleList(itemsRepo, input.Tag)
			}
		},
	}

	return cmd
}

func CreateCli() *cobra.Command {
	cli := createRootCmd()

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
 