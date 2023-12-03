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
type ChatRequest struct {
	Message string `json:"message"`
}
type ChatResponse struct {
	Message string `json:"message"`
}
type Action struct {
	Name string `json:"name"`
	Template string `json:"template"`
}
type ActionResponse struct {
	Items []Action `json:"items"`
}

func Serve(repos repository.Repos, port int) error {
	app := fiber.New()

	app.Get("/api/binders", func(c *fiber.Ctx) error {
		binders := ListBinders(repos)
		return c.JSON(binders)
	})

	app.Get("/api/binder/:id/notes", func(c *fiber.Ctx) error {
		binderId, err := c.ParamsInt("id")
		if err != nil {
			return err
		}
		notes, err := ListBinderNotes(repos, uint(binderId))
		if err != nil {
			return err
		}
		return c.JSON(notes)
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

	app.Get("/api/actions", func(c *fiber.Ctx) error {
		res := ActionResponse{
			Items: make([]Action, 0),
		}
		actions, err := ListActions(repos)
		if err != nil {
			return err
		}
		for _, action := range actions {
			res.Items = append(res.Items, Action{
				Name: action.Name,
				Template: action.Template,
			})
		}
		return c.JSON(res)
	})

	app.Post("/api/chat", func(c *fiber.Ctx) error {
		var req ChatRequest
		if err := c.BodyParser(&req); err != nil {
			return err
		}

		configSrv := service.NewConfigSevice(repository.NewRepos())
		config, err := configSrv.Read()
		if err != nil {
			return err
		}

		chatgptSrv := service.NewAiService(repos)
		res, err := chatgptSrv.Call(config.ChatgptToken, req.Message)
		if err != nil {
			return err
		}
		return c.JSON(ChatResponse{ Message: res })
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
