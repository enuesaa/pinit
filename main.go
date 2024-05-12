package main

import (
	"log"

	"github.com/enuesaa/pinit/pkg/cli"
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/spf13/cobra"
)

func init() {
	log.SetFlags(0)
}

func main() {
	start()

	app := &cobra.Command{
		Use:     "pinit",
		Short:   "A personal note-taking app",
		Version: "0.0.8",
	}

	repos := repository.NewRepos()
	app.AddCommand(cli.CreateConfigureCmd(repos))
	app.AddCommand(cli.CreateLsCmd(repos))
	app.AddCommand(cli.CreateCreateCmd(repos))
	app.AddCommand(cli.CreateWriteCmd(repos))
	app.AddCommand(cli.CreateLookupCmd(repos))
	app.AddCommand(cli.CreateRmCmd(repos))
	app.AddCommand(cli.CreateServeCmd(repos))

	// disable default
	app.SetHelpCommand(&cobra.Command{Hidden: true})
	app.CompletionOptions.DisableDefaultCmd = true
	app.SilenceUsage = true
	app.PersistentFlags().SortFlags = false
	app.PersistentFlags().BoolP("help", "", false, "Show help information")
	app.PersistentFlags().BoolP("version", "", false, "Show version")

	app.Execute()
}
