package cli

import (
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/spf13/cobra"
)

func CreateDescribeCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:  "describe <name>",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			noteSrv := service.NoteService{}
			noteSrv.Get("")
		},
	}

	return cmd
}
