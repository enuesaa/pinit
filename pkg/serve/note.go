package serve

import (
	"github.com/enuesaa/pinit/pkg/usecase"
	"github.com/gofiber/fiber/v2"
)

func (ctl *ServeCtl) ListNotes(c *fiber.Ctx) error {
	binderId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	notes, err := usecase.ListBinderNotes(ctl.Repos, uint(binderId))
	if err != nil {
		return err
	}
	return c.JSON(notes)
}