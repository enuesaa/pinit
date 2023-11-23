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
		Use:   "lookup <name>",
		Short: "lookup a note.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]

			configSrv := service.NewConfigSevice(repos)
			if err := configSrv.Init(); err != nil {
				fmt.Printf("Error: %s", err.Error())
				return
			}

			binderSrv := service.NewBinderService(repos)
			binder, err := binderSrv.Get(name)
			if err != nil {
				fmt.Printf("Error: %s", err.Error())
				return
			}

			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{"ID", "NAME", "CREATED AT"})
			t.AppendRow(table.Row{binder.ID, binder.Name, binder.CreatedAt})
			t.SetStyle(table.StyleLight)
			t.Render()
		},
	}

	return cmd
}
