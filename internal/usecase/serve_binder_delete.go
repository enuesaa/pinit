package usecase

import (
	"github.com/enuesaa/pinit/internal/service"
	"github.com/gofiber/fiber/v2"
)

func (ctl *ServeCtl) BinderDelete(c *fiber.Ctx) error {
	binderName := c.Params("binderName")

	noteSrv := service.NewNoteService(binderName, ctl.repos)
	binderSrv := service.NewBinderService(ctl.repos)

	if err := noteSrv.DeleteAllInBinder(); err != nil {
		return err
	}
	if err := binderSrv.Delete(binderName); err != nil {
		return err
	}

	return c.JSON(ServeDeleteResponse{})
}
