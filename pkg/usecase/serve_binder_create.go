package usecase

import (
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func (ctl *ServeCtl) BinderCreate(c *fiber.Ctx) error {
	type CreateBinderRequest struct {
		Name string `json:"name"`
	}

	var req CreateBinderRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	binderSrv := service.NewBinderService(ctl.repos)

	binder := service.Binder{
		Name:     req.Name,
		Category: "",
	}
	id, err := binderSrv.Create(binder)
	if err != nil {
		return err
	}

	return c.JSON(ServeCreateResponse{Id: id})
}
