package cli

import (
	"github.com/spf13/cobra"
)

type CliInput struct {
	Register bool
	Remove   bool
	Tag      string
	Filename string
}

func (cli *CliInput) HasTagFlag() bool {
	return cli.Tag != ""
}

func(cli *CliInput) HasFilename() bool {
	return cli.Filename != ""
}

func (cli *CliInput) IsOperationAmbiguous() bool {
	return cli.Register && cli.Remove
}

func ParseArgs(cmd *cobra.Command, args []string) CliInput {
	register, _ := cmd.Flags().GetBool("register")
	remove, _ := cmd.Flags().GetBool("remove")
	tag, _ := cmd.Flags().GetString("tag")
	filename := ""
	if len(args) > 0 {
		filename = args[0]
	}

	input := CliInput {
		Register: register,
		Remove: remove,
		Tag: tag,
		Filename: filename,
	}

	return input
}
