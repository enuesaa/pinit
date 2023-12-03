package main

import (
	"log"

	"github.com/enuesaa/pinit/pkg/cli"
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/spf13/cobra"
)

func init() {
	log.SetFlags(0)

	repos := repository.NewRepos()
	configSrv := service.NewConfigSevice(repos)
	configSrv.Init()
}

func main() {
	app := &cobra.Command{
		Use:     "pinit",
		Short:   "A personal note-taking app",
		Version: "0.0.3",
	}

	repos := repository.NewRepos()
	app.AddCommand(cli.CreateConfigureCmd(repos))
	app.AddCommand(cli.CreateLsCmd(repos))
	app.AddCommand(cli.CreateCreateCmd(repos))
	app.AddCommand(cli.CreateWriteCmd(repos))
	app.AddCommand(cli.CreateDescribeCmd(repos))
	app.AddCommand(cli.CreateRmCmd(repos))
	app.AddCommand(cli.CreateServeCmd(repos))
	app.AddCommand(cli.CreateActionCmd(repos))

	// disable default
	app.SetHelpCommand(&cobra.Command{Hidden: true})
	app.CompletionOptions.DisableDefaultCmd = true
	app.SilenceUsage = true
	app.PersistentFlags().SortFlags = false
	app.PersistentFlags().BoolP("help", "", false, "Show help information")
	app.PersistentFlags().BoolP("version", "", false, "Show version")
	app.SetHelpTemplate(`{{.Short}}
{{if .HasAvailableSubCommands}}{{$cmds := .Commands}}{{if eq (len .Groups) 0}}
Available Commands:{{range $cmds}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{else}}{{range $group := .Groups}}

{{.Title}}{{range $cmds}}{{if (and (eq .GroupID $group.ID) (or .IsAvailableCommand (eq .Name "help")))}}
{{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{end}}
{{end}}
{{if .HasAvailableLocalFlags}}Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
{{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}
`)

	app.Execute()
}
