package usecase

import (
	"github.com/enuesaa/pinit/internal/service"
	"github.com/gofiber/fiber/v2"
)

func (ctl *ServeCtl) Chat(c *fiber.Ctx) error {
	type ChatRequest struct {
		Message string `json:"message"`
	}
	type ChatResponse struct {
		Message string `json:"message"`
	}
	
	var req ChatRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	chatgptSrv := service.NewAiService(ctl.repos)
	registrySrv := service.NewRegistrySrv(ctl.repos)

	config := registrySrv.Get()

	res, err := chatgptSrv.Call(config.OpenaiToken, req.Message)
	if err != nil {
		return err
	}

	return c.JSON(ChatResponse{Message: res})
}
