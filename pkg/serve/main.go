package serve

import (
	"fmt"
	"mime"
	"path/filepath"
	"strings"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/web"
	"github.com/gofiber/fiber/v2"
)

func Serve(repos repository.Repos, port int) error {
	app := fiber.New()

	ctl := Controller{repos: repos}
	app.Get("/api/binders", ctl.ListBinders)
	app.Get("/api/binders/:id/notes", ctl.ListNotes)
	app.Get("/api/config", ctl.GetConfig)
	app.Get("/api/actions", ctl.ListActions)
	app.Post("/api/chat", ctl.Chat)
	app.Get("/*", func(c *fiber.Ctx) error {
		requestPath := c.Path() // like `/`

		path := fmt.Sprintf("dist%s", requestPath) // like `./`
		if strings.HasSuffix(path, "/") {
			path += "index.html"
		}

		f, err := web.Dist.ReadFile(path)
		if err != nil {
			return err
		}
		fileExt := filepath.Ext(path)
		mimeType := mime.TypeByExtension(fileExt)
		c.Set(fiber.HeaderContentType, mimeType)

		return c.SendString(string(f))
	})

	addr := fmt.Sprintf(":%d", port)
	return app.Listen(addr)
}
