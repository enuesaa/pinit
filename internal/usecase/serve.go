package usecase

import (
	"fmt"

	"github.com/enuesaa/pinit/internal/repository"
	"github.com/enuesaa/pinit/ui"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pkg/browser"
)

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

func Serve(repos repository.Repos, port int) error {
	ctl := ServeCtl{
		repos: repos,
		port: port,
	}

	if err := ctl.Open(); err != nil {
		// ignore this err because this is not critical for app.
		ctl.repos.Log.Info("failed to open url because `%s`", err.Error())
	}

	return ctl.Serve()
}
