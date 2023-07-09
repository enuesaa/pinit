package command

import (
	"github.com/spf13/cobra"
)

type CliInput struct {
	Register string
	Remove   string
	Tag      string
}

func (cli *CliInput) HasRegisterFlag() bool {
	return cli.Register != ""
}

func (cli *CliInput) HasTagFlag() bool {
	return cli.Tag != ""
}

func (cli *CliInput) HasRemoveFlag() bool {
	return cli.Remove != ""
}

func (cli *CliInput) IsOperationAmbiguous() bool {
	return cli.HasRegisterFlag() && cli.HasRemoveFlag()
}

func (cli *CliInput) HasFilename() bool {
	return false
}

func ParseArgs(cmd *cobra.Command, args []string) CliInput {
	register, _ := cmd.Flags().GetString("register")
	remove, _ := cmd.Flags().GetString("remove")
	tag, _ := cmd.Flags().GetString("tag")

	input := CliInput {
		Register: register,
		Remove: remove,
		Tag: tag,
	}

	return input
}
