package main

import (
	"github.com/enuesaa/pinit/internal/repository"
	"github.com/enuesaa/pinit/internal/usecase"
	"github.com/spf13/cobra"
)

func main() {
	repos := repository.New()

	app := &cobra.Command{
		Use:     "pinit",
		Short:   "A personal note-taking app",
		Version: "0.0.11",
		RunE: func(cmd *cobra.Command, args []string) error {
			port, _ := cmd.Flags().GetInt("port")
			isServe, _ := cmd.Flags().GetBool("serve")

			if !isServe {
				return cmd.Help()
			}

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
