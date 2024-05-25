package cli

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/usecase"
	"github.com/enuesaa/pinit/web"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

func CreateServeCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "serve",
		Short: "Serve pinit app",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return usecase.DBOpen(repos)
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return usecase.DBClose(repos)
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
	cmd.Flags().Int("port", 3000, "port. Default: 3000")

	return cmd
}
