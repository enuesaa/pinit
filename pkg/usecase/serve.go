package usecase

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/serve"
	"github.com/enuesaa/pinit/web"
	"github.com/gofiber/fiber/v2"
)

func Serve(repos repository.Repos, port int) error {
	ctl := serve.New(repos)

	app := fiber.New()
	app.Get("/api/binders", ctl.ListBinders)
	app.Get("/api/binders/:id/notes", ctl.ListNotes)
	app.Get("/api/actions", ctl.ListActions)
	app.Post("/api/chat", ctl.Chat)
	app.Post("/api/recog", ctl.Recog)
	app.Get("/*", web.Serve)

	addr := fmt.Sprintf(":%d", port)
	return app.Listen(addr)
}
