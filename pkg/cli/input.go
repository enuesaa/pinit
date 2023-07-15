package cli

import (
	"github.com/spf13/cobra"
)

type CliInput struct {
	Register bool
	Remove   bool
	Tag      string
}

func (cli *CliInput) HasTagFlag() bool {
	return cli.Tag != ""
}

func (cli *CliInput) IsOperationAmbiguous() bool {
	return cli.Register && cli.Remove
}

func (cli *CliInput) HasFilename() bool {
	return false
}

func ParseArgs(cmd *cobra.Command, args []string) CliInput {
	register, _ := cmd.Flags().GetBool("register")
	remove, _ := cmd.Flags().GetBool("remove")
	tag, _ := cmd.Flags().GetString("tag")

	input := CliInput {
		Register: register,
		Remove: remove,
		Tag: tag,
	}

	return input
}
