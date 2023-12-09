package serve

import (
	"fmt"
	"log"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/web"
	"github.com/gofiber/fiber/v2"
)

func Serve(repos repository.Repos, port int) {
	app := fiber.New()
	ctl := Controller{
		Repos: repos,
	}
	app.Get("/api/binders", ctl.ListBinders)
	app.Get("/api/binders/:id/notes", ctl.ListNotes)
	app.Get("/api/actions", ctl.ListActions)
	app.Post("/api/chat", ctl.Chat)
	app.Post("/api/recog", ctl.Recog)
	app.Get("/*", web.Serve)

	addr := fmt.Sprintf(":%d", port)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
