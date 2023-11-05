package service

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/enuesaa/pinit/admin"
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/gofiber/fiber/v2"
)

type AdminService struct {
	repos repository.Repos
}

func NewAdminService(repos repository.Repos) *AdminService {
	return &AdminService{
		repos: repos,
	}
}

func (srv *AdminService) Serve() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Get("/*", func(c *fiber.Ctx) error {
		requestPath := c.Path() // like `/`

		path := fmt.Sprintf("dist%s", requestPath) // like `./`
		if strings.HasSuffix(path, "/") {
			path += "index.html"
		}

		fmt.Println(path)

		f, err := admin.Dist.ReadFile(path)
		if err != nil {
			return err
		}

		mimeType := http.DetectContentType(f)
		c.Set(fiber.HeaderContentType, mimeType)

		return c.SendString(string(f))
	})

	app.Listen(":3000")
}
