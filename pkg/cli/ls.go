package cli

import (
	"fmt"
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
			binders := usecase.ListBinders(repos)
			fmt.Printf("%d binder(s) found.\n", len(binders))
			if len(binders) == 0 {
				return
			}

			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{"ID", "NAME", "CREATED AT"})
			for _, binder := range binders {
				t.AppendRow(table.Row{binder.ID, binder.Name, binder.CreatedAt})
			}
			t.SetStyle(table.StyleLight)
			t.Render()
		},
	}

	return cmd
}
