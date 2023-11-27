package usecase

import (
	"fmt"
	"mime"
	"path/filepath"
	"strings"

	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/enuesaa/pinit/web"
	"github.com/gofiber/fiber/v2"
)

type ConfigResponse struct {
	Token string `json:"token"`
}

func Serve(port int) error {
	app := fiber.New()

	app.Get("/api/binders", func(c *fiber.Ctx) error {
		return nil
	})

	app.Get("/api/config", func(c *fiber.Ctx) error {
		configSrv := service.NewConfigSevice(repository.NewRepos())
		config, err := configSrv.Read()
		if err != nil {
			return err
		}
		res := ConfigResponse{
			Token: config.ChatgptToken,
		}
		return c.JSON(res)
	})
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
