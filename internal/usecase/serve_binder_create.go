package usecase

import (
	"github.com/enuesaa/pinit/internal/service"
	"github.com/gofiber/fiber/v2"
)

func (ctl *ServeCtl) BinderCreate(c *fiber.Ctx) error {
	type Request struct {}
	type Response struct {
		Name string `json:"name"`
	}

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	binderSrv := service.NewBinderService(ctl.repos)

	name, err := binderSrv.Create()
	if err != nil {
		return err
	}
	return c.JSON(Response{Name: name})
}
