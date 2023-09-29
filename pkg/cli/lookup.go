package cli

import (
	"fmt"
	"os"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

func CreateLookupCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:  "lookup <name>",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]

			// setup
			configSrv := service.NewConfigSevice(repos)
			databaseDsn, err := configSrv.ReadDatabaseDsn()
			if err != nil {
				return
			}
			repos.Database.WithDsn(databaseDsn)

			noteSrv := service.NewNoteService(repos)
			note, err := noteSrv.Get(name)
			if err != nil {
				fmt.Println(err)
				return
			}

			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{"ID", "NAME", "CONTENT", "COMMENT", "CREATED AT"})
			t.AppendRow(table.Row{note.ID, note.Name, note.Content, note.Comment, note.CreatedAt})
			t.SetStyle(table.StyleLight)
			t.Render()
		},
	}

	return cmd
}