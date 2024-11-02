package usecase

import (
	"github.com/enuesaa/pinit/internal/service"
	"github.com/gofiber/fiber/v2"
)

func (ctl *ServeCtl) SettingView(c *fiber.Ctx) error {
	type Response struct {
		OpenaiToken string `json:"openaiToken"`
	}

	registrySrv := service.NewRegistrySrv(ctl.repos)
	config := registrySrv.Get()

	res := Response{
		OpenaiToken: config.OpenaiToken,
	}
	return c.JSON(res)
}
