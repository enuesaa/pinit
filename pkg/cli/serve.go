package cli

import (
	"fmt"
	"log"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/usecase"
	"github.com/enuesaa/pinit/web"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

func CreateServeCmd(repos repository.Repos) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "serve",
		Short: "Launch the web console",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return usecase.OpenDb(repos)
		},
		PostRunE: func(cmd *cobra.Command, args []string) error {
			return usecase.CloseDb(repos)
		},
		Run: func(cmd *cobra.Command, args []string) {
			port, _ := cmd.Flags().GetInt("port")

			ctl := usecase.NewServeCtl(repos)
			app := fiber.New()
			app.Get("/api/binders", ctl.ListBinders)
			app.Post("/api/binders", ctl.CreateBinder)
			app.Delete("/api/binders/:id", ctl.DeleteBinder)
			app.Get("/api/binders/:id/notes", ctl.ListNotes)
			//TODO change endpoint
			app.Post("/api/notes", ctl.CreateNote)
			app.Get("/api/actions", ctl.ListActions)
			app.Post("/api/chat", ctl.Chat)
			app.Post("/api/recog", ctl.Recog)
			app.Get("/*", web.Serve)
		
			addr := fmt.Sprintf(":%d", port)
			if err := app.Listen(addr); err != nil {
				log.Fatalf("Error: %s", err.Error())
			}
		},
	}
	cmd.Flags().Int("port", 3000, "port. Default: 3000")

	return cmd
}
