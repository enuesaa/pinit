package cli

import (
	"fmt"
	"os"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

func CreateLsCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use: "ls",
		Short: "list notes.",
		Run: func(cmd *cobra.Command, args []string) {
			configSrv := service.NewConfigSevice(repos)
			if err := configSrv.ConfigureDatabaseDsn(); err != nil {
				fmt.Printf("Error: %s", err.Error())
				return
			}

			noteSrv := service.NewNoteService(repos)
			notes := noteSrv.List()
			fmt.Printf("%d note(s) found.\n", len(notes))
			if len(notes) == 0 {
				return
			}

			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{"ID", "NAME", "CONTENT", "COMMENT", "CREATED AT"})
			for _, note := range notes {
				t.AppendRow(table.Row{note.ID, note.Name, note.Content, note.Comment, note.CreatedAt})
			}
			t.SetStyle(table.StyleLight)
			t.Render()
		},
	}
	// cmd.Flags().StringArrayP("filter", "", []string{}, "search vavlue")

	return cmd
}
