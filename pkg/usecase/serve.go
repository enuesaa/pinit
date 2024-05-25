package usecase

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/ui"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pkg/browser"
)


func NewServeCtl(repos repository.Repos, port int) ServeCtl {
	return ServeCtl{
		repos: repos,
		port: port,
	}
}

type ServeCtl struct {
	repos repository.Repos
	port int
}

func (ctl *ServeCtl) CreateId() string {
	return uuid.New().String()
}

func (ctl *ServeCtl) Addr() string {
	return fmt.Sprintf("localhost:%d", ctl.port)
}

func (ctl *ServeCtl) Open() error {
	url := fmt.Sprintf("http://%s", ctl.Addr())
	return browser.OpenURL(url)
}

func (ctl *ServeCtl) Serve() error {
	app := fiber.New()
	app.Get("/api/binders", ctl.BinderList)
	app.Post("/api/binders", ctl.BinderCreate)
	app.Delete("/api/binders/:id", ctl.BinderDelete)
	app.Get("/api/binders/:id/notes", ctl.NoteList)
	app.Post("/api/notes", ctl.NoteCreate)
	app.Get("/api/actions", ctl.ActionList)
	app.Post("/api/chat", ctl.Chat)
	app.Post("/api/recog", ctl.Recog)
	app.Get("/*", ui.Serve)

	return app.Listen(ctl.Addr())
}
