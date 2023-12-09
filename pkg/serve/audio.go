package serve

import (
	"fmt"
	"os"

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

	if err := os.MkdirAll("tmp", os.ModePerm); err != nil {
		return err
	}

	id := uuid.New().String()
	path := fmt.Sprintf("tmp/%s.wav", id)

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.Write(body); err != nil {
		return err
	}

	text, err := aiSrv.Speak(config.ChatgptToken, path)
	if err != nil {
		return err
	}

	return c.JSON(RecogResponse{Text: text})
}

