package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/spf13/cobra"
)

func CreateDescribeCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:  "describe <name>",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			noteSrv := service.NewNoteService(repos)
			note, err := noteSrv.Get(name)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%+v", note)
			}
		},
	}

	return cmd
}
