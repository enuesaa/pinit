package usecase

import (
	"fmt"

	"github.com/enuesaa/pinit/internal/repository"
	"github.com/enuesaa/pinit/ui"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/google/uuid"
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

func (ctl *ServeCtl) Serve() error {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
	}))
	
	app.Get("/api/binders", ctl.BinderList)
	app.Post("/api/binders", ctl.BinderCreate)
	app.Put("/api/binders/:binderName", ctl.BinderUpdate)
	app.Delete("/api/binders/:binderName", ctl.BinderDelete)
	app.Get("/api/binders/:binderName/notes", ctl.NoteList)
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
	return ctl.Serve()
}
