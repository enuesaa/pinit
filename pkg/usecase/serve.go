package usecase

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/web"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func NewServeCtl(repos repository.Repos) ServeCtl {
	return ServeCtl{
		repos: repos,
	}
}

type ServeCtl struct {
	repos repository.Repos
}

func (ctl *ServeCtl) CreateId() string {
	return uuid.New().String()
}

func Serve(repos repository.Repos, port int) error {
	ctl := NewServeCtl(repos)

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
