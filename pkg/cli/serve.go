package cli

import (
	"fmt"
	"log"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/web"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

func CreateServeCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "serve",
		Short: "Launch the web console",
		Run: func(cmd *cobra.Command, args []string) {
			port, _ := cmd.Flags().GetInt("port")

			app := fiber.New()
			ctl := web.Controller{
				Repos: repos,
			}
			app.Get("/api/binders", ctl.ListBinders)
			app.Get("/api/binders/:id/notes", ctl.ListNotes)
			app.Get("/api/config", ctl.GetConfig)
			app.Get("/api/actions", ctl.ListActions)
			app.Post("/api/chat", ctl.Chat)
			app.Get("/*", ctl.ServeStatic)

			addr := fmt.Sprintf(":%d", port)
			if err := app.Listen(addr); err != nil {
				log.Printf("Error: %s\n", err.Error())
				return
			}
		},
	}
	cmd.Flags().Int("port", 3000, "port. Default: 3000")

	return cmd
}
