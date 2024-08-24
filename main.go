package main

import (
	"log"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/usecase"
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
		RunE: func(cmd *cobra.Command, args []string) error {
			port, _ := cmd.Flags().GetInt("port")
			isServe, _ := cmd.Flags().GetBool("serve")

			if !isServe {
				return cmd.Help()
			}

			if !usecase.DBIsExist(repos) {
				if err := usecase.DBSetup(repos); err != nil {
					return err
				}
			}
			if err := usecase.DBOpen(repos); err != nil {
				return err
			}
			defer usecase.DBClose(repos)

			return usecase.Serve(repos, port)
		},
	}
	app.Flags().Bool("serve", false, "Start pinit app")
	app.Flags().Int("port", 3000, "port. Default: 3000")

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
