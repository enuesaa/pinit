package serve

import (
	"fmt"
	"log"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/web"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func NewServeCtl(repos repository.Repos) ServeCtl {
	return ServeCtl{
		Repos: repos,
	}
}

type ServeCtl struct {
	Repos repository.Repos
}

func (ctl *ServeCtl) CreateId() string {
	return uuid.New().String()
}

func (ctl *ServeCtl) Serve(port int) {
	app := fiber.New()
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
