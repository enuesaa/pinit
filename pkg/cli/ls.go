package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/usecase"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

func CreateLsCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "ls",
		Short: "List binders",
		Run: func(cmd *cobra.Command, args []string) {
			binders, err := usecase.ListBinders(repos)
			if err != nil {
				log.Fatalf("Error: %s\n", err.Error())
			}
			fmt.Printf("%d binder(s) found.\n", len(binders))
			if len(binders) == 0 {
				return
			}

			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{"ID", "NAME", "ARCHIVED AT", "CREATED AT"})
			for _, binder := range binders {
				archivedAt := ""
				if binder.ArchivedAt != nil {
					archivedAt = binder.ArchivedAt.Format("2006-01-02T15:04:05")
				}
				createdAt := binder.CreatedAt.Format("2006-01-02T15:04:05")
				t.AppendRow(table.Row{binder.ID, binder.Name, archivedAt, createdAt})
			}
			t.SetStyle(table.StyleLight)
			t.Render()
		},
	}

	return cmd
}
