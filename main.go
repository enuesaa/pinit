package main

import (
	"log"
	"os"

	"github.com/enuesaa/pinit/pkg/cli"
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/spf13/cobra"
)

func init() {
	log.SetFlags(0)
}

func main() {
	dbPath := os.Getenv("PINIT_DB_PATH")
	repos := repository.New(dbPath)

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
		log.Fatalf("Error: %s", err.Error())
	}
}
