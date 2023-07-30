package cli

import (
	"github.com/spf13/cobra"
)

// todo: separate per sub command
type CliInput struct {
	Tag      string
	Filename string
}

func (cli *CliInput) HasTagFlag() bool {
	return cli.Tag != ""
}

func(cli *CliInput) HasFilename() bool {
	return cli.Filename != ""
}

func ParseArgs(cmd *cobra.Command, args []string) CliInput {
	tag, _ := cmd.Flags().GetString("tag")
	filename := ""
	if len(args) > 0 {
		filename = args[0]
	}

	input := CliInput {
		Tag: tag,
		Filename: filename,
	}

	return input
}
