package serve

import (
	"fmt"
	"path/filepath"

	"github.com/enuesaa/pinit/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type RecogResponse struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

func (ctl *ServeCtl) Recog(c *fiber.Ctx) error {
	body := c.BodyRaw()
	aiSrv := service.NewAiService(ctl.Repos)

	id := ctl.CreateId()
	path := fmt.Sprintf("tmp/%s.wav", id)
	if err := ctl.Repos.Fs.CreateDir(filepath.Dir(path)); err != nil {
		return err
	}
	if err := ctl.Repos.Fs.Create(path, body); err != nil {
		return err
	}

	registrySrv := service.NewRegistrySrv(ctl.Repos)
	text, err := aiSrv.Speak(registrySrv.GetOpenAiApiToken(), path)
	if err != nil {
		return err
	}

	res := RecogResponse{
		Id:   id,
		Text: text,
	}
	return c.JSON(res)
}
