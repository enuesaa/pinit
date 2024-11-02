package usecase

import (
	"github.com/enuesaa/pinit/internal/service"
	"github.com/gofiber/fiber/v2"
)

func (ctl *ServeCtl) SettingUpdate(c *fiber.Ctx) error {
	type Request struct {
		OpenaiToken string `json:"openaiToken"`
	}
	type Response struct {}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	registrySrv := service.NewRegistrySrv(ctl.repos)

	creation := service.AppConfigCreation{
		OpenaiToken: req.OpenaiToken,
	}
	if err := registrySrv.Update(creation); err != nil {
		return err
	}

	return c.JSON(Response{})
}
