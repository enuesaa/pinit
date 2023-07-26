package cli

import (
	"github.com/spf13/cobra"
)

type Operation int

const (
    List Operation = iota
    Register
	Remove
	Apply
)


type CliInput struct {
	Operation Operation
	Tag      string
	Filename string
}

func (cli *CliInput) HasTagFlag() bool {
	return cli.Tag != ""
}

func(cli *CliInput) HasFilename() bool {
	return cli.Filename != ""
}

func ParseArgs(operation Operation, cmd *cobra.Command, args []string) CliInput {
	tag, _ := cmd.Flags().GetString("tag")
	filename := ""
	if len(args) > 0 {
		filename = args[0]
	}

	input := CliInput {
		Operation: operation,
		Tag: tag,
		Filename: filename,
	}

	return input
}
