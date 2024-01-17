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
	return app.Listen(addr)
}
