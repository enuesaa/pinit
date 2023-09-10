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

	// disable default
	cli.SetHelpCommand(&cobra.Command{Hidden: true})
	cli.PersistentFlags().BoolP("help", "", false, "Show help information")
	cli.CompletionOptions.DisableDefaultCmd = true

	return cli
}

// todo: separate per sub command
type CliInput struct {
	Tag      string
	Filename string
}

func (cli *CliInput) HasTagFlag() bool {
	return cli.Tag != ""
}

func (cli *CliInput) HasFilename() bool {
	return cli.Filename != ""
}

func ParseArgs(cmd *cobra.Command, args []string) CliInput {
	tag, _ := cmd.Flags().GetString("tag")
	filename := ""
	if len(args) > 0 {
		filename = args[0]
	}

	input := CliInput{
		Tag:      tag,
		Filename: filename,
	}

	return input
}
