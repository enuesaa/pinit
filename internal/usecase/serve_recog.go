package usecase

import (
	"bytes"

	"github.com/enuesaa/pinit/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/oklog/ulid/v2"
)

func (ctl *ServeCtl) Recog(c *fiber.Ctx) error {
	type RecogResponse struct {
		Id   string `json:"id"`
		Text string `json:"text"`
	}

	body := c.BodyRaw()
	aiSrv := service.NewAiService(ctl.repos)

	id := ulid.Make().String()
	registrySrv := service.NewRegistrySrv(ctl.repos)
	config := registrySrv.Get()

	text, err := aiSrv.Speak(config.OpenaiToken, bytes.NewReader(body))
	if err != nil {
		return err
	}

	res := RecogResponse{
		Id:   id,
		Text: text,
	}
	return c.JSON(res)
}
