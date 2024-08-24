package usecase

import (
	"github.com/enuesaa/pinit/internal/service"
	"github.com/gofiber/fiber/v2"
)

func (ctl *ServeCtl) BinderDelete(c *fiber.Ctx) error {
	binderId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	noteSrv := service.NewNoteService(ctl.repos)
	binderSrv := service.NewBinderService(ctl.repos)

	if err := noteSrv.DeleteByBinderId(uint(binderId)); err != nil {
		return err
	}
	if err := binderSrv.Delete(uint(binderId)); err != nil {
		return err
	}

	return c.JSON(ServeDeleteResponse{})
}
