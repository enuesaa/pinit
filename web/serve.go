package web

import (
	"fmt"
	"mime"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (ctl *Controller) ServeStatic(c *fiber.Ctx) error {
	requestPath := c.Path() // like `/`

	path := fmt.Sprintf("dist%s", requestPath) // like `./`
	if strings.HasSuffix(path, "/") {
		path += "index.html"
	}

	f, err := Dist.ReadFile(path)
	if err != nil {
		return err
	}
	fileExt := filepath.Ext(path)
	mimeType := mime.TypeByExtension(fileExt)
	c.Set(fiber.HeaderContentType, mimeType)

	return c.SendString(string(f))
}
