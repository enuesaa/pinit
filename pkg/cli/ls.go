package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/spf13/cobra"
)

func CreateLsCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use: "ls",
		Run: func(cmd *cobra.Command, args []string) {
			noteSrv := service.NewNoteService(repos)
			notes := noteSrv.List()
			for _, note := range notes {
				fmt.Printf("%+v\n", note)
			}
		},
	}
	cmd.Flags().StringArrayP("filter", "", []string{}, "search vavlue")

	return cmd
}
