package usecase

import (
	"fmt"
	"mime"
	"path/filepath"
	"strings"

	"github.com/enuesaa/pinit/web"
	"github.com/gofiber/fiber/v2"
)

func Serve(port int) error {
	app := fiber.New()
	app.Get("/*", func(c *fiber.Ctx) error {
		requestPath := c.Path() // like `/`

		path := fmt.Sprintf("dist%s", requestPath) // like `./`
		if strings.HasSuffix(path, "/") {
			path += "index.html"
		}

		fmt.Println(path)

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
