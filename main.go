package main

import (
	"log"

	"github.com/enuesaa/pinit/pkg/cli"
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/spf13/cobra"
)

func main() {
	repos, err := repository.New()
	if err != nil {
		log.Panic("failed to start up app")
	}

	app := &cobra.Command{
		Use:     "pinit",
		Short:   "A personal note-taking app",
		Version: "0.0.10",
	}
	app.AddCommand(cli.CreateInitCmd(repos))
	app.AddCommand(cli.CreateServeCmd(repos))

	// disable default
	app.SetHelpCommand(&cobra.Command{Hidden: true})
	app.CompletionOptions.DisableDefaultCmd = true
	app.SilenceUsage = true
	app.PersistentFlags().SortFlags = false
	app.PersistentFlags().BoolP("help", "", false, "Show help information")
	app.PersistentFlags().BoolP("version", "", false, "Show version")
	app.SilenceErrors = true

	if err := app.Execute(); err != nil {
		repos.Log.Err(err)
	}
}
