package main

import (
	"fmt"
	"log"

	"github.com/enuesaa/pinit/pkg/cli"
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/usecase"
	"github.com/enuesaa/pinit/web"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

func init() {
	log.SetFlags(0)
}

func main() {
	repos := repository.NewRepos()

	app := &cobra.Command{
		Use:     "pinit",
		Short:   "A personal note-taking app",
		Version: "0.0.8",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return usecase.OpenDb(repos)
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return usecase.CloseDb(repos)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			port, _ := cmd.Flags().GetInt("port")

			ctl := usecase.NewServeCtl(repos)
			app := fiber.New()
			app.Get("/api/binders", ctl.ListBinders)
			app.Post("/api/binders", ctl.CreateBinder)
			app.Delete("/api/binders/:id", ctl.DeleteBinder)
			app.Get("/api/binders/:id/notes", ctl.ListNotes)
			app.Post("/api/notes", ctl.CreateNote)
			app.Get("/api/actions", ctl.ListActions)
			app.Post("/api/chat", ctl.Chat)
			app.Post("/api/recog", ctl.Recog)
			app.Get("/*", web.Serve)

			addr := fmt.Sprintf("localhost:%d", port)
			url := fmt.Sprintf("http://%s", addr)
			if err := browser.OpenURL(url); err != nil {
				return err
			}
			return app.Listen(addr)
		},
	}
	app.Flags().Int("port", 3000, "port. Default: 3000")

	app.AddCommand(cli.CreateConfigureCmd(repos))

	// disable default
	app.SetHelpCommand(&cobra.Command{Hidden: true})
	app.CompletionOptions.DisableDefaultCmd = true
	app.SilenceUsage = true
	app.PersistentFlags().SortFlags = false
	app.PersistentFlags().BoolP("help", "", false, "Show help information")
	app.PersistentFlags().BoolP("version", "", false, "Show version")

	if err := app.Execute(); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
