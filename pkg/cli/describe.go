package cli

import (
	"fmt"
	"os"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/jedib0t/go-pretty/v6/table"
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
				t := table.NewWriter()
				t.SetOutputMirror(os.Stdout)
				t.AppendHeader(table.Row{"ID", "NAME", "CONTENT", "COMMENT", "CREATED AT"})
				t.AppendRow(table.Row{note.ID, note.Name, note.Content, note.Comment, note.CreatedAt})
				t.SetStyle(table.StyleLight)
				t.Render()
			}
		},
	}

	return cmd
}
