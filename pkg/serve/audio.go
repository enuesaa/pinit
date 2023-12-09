package serve

import (
	"fmt"
	"github.com/enuesaa/pinit/pkg/repository"
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RecogResponse struct {
	Text string `json:"text"`
}

func (ctl *Controller) Recog(c *fiber.Ctx) error {
	body := c.BodyRaw()

	configSrv := service.NewConfigSevice(repository.NewRepos())
	config, err := configSrv.Read()
	if err != nil {
		return err
	}

	aiSrv := service.NewAiService(ctl.Repos)

	if err := ctl.Repos.Fs.CreateDir("tmp"); err != nil {
		return err
	}
	id := uuid.New().String()
	path := fmt.Sprintf("tmp/%s.wav", id)
	if err := ctl.Repos.Fs.Create(path, body); err != nil {
		return err
	}

	text, err := aiSrv.Speak(config.ChatgptToken, path)
	if err != nil {
		return err
	}

	return c.JSON(RecogResponse{Text: text})
}

