package serve

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type ChatRequest struct {
	Message string `json:"message"`
}
type ChatResponse struct {
	Message string `json:"message"`
}

func (ctl *ServeCtl) Chat(c *fiber.Ctx) error {
	var req ChatRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	chatgptSrv := service.NewAiService(ctl.Repos)
	res, err := chatgptSrv.Call(ctl.Repos.Config.ChatgptToken(), req.Message)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return err
	}
	return c.JSON(ChatResponse{Message: res})
}
