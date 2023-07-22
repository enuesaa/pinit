package cli

import (
	"github.com/spf13/cobra"
)

type CliInput struct {
	Register bool
	Remove   bool
	Tags     []string
}

func (cli *CliInput) HasTagFlag() bool {
	return len(cli.Tags) > 0
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
	tags := make([]string, 0)
	if tag != "" {
		tags = append(tags, tag)
	}

	input := CliInput {
		Register: register,
		Remove: remove,
		Tags: tags,
	}

	return input
}
