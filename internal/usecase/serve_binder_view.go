package usecase

import (
	"github.com/enuesaa/pinit/internal/service"
	"github.com/gofiber/fiber/v2"
)

func (ctl *ServeCtl) BinderView(c *fiber.Ctx) error {
	type Item struct {
		Name       string `json:"name"`
		Content    string `json:"content"`
		ArchivedAt string `json:"archivedAt"`
		CreatedAt  string `json:"createdAt"`
		UpdatedAt  string `json:"updatedAt"`
	}
	binderName := c.Params("binderName")

	binderSrv := service.NewBinderService(ctl.repos)
	noteSrv := service.NewNoteService(binderName, ctl.repos)

	binder, err := binderSrv.Get(binderName)
	if err != nil {
		return err
	}

	var res Item
	res.Name = binderName
	res.ArchivedAt = ""
	res.Content = ""
	res.CreatedAt = binder.CreatedAt.String()
	res.UpdatedAt = binder.UpdatedAt.String()

	notes, err := noteSrv.List()
	if err != nil {
		return c.JSON(res)
	}
	if len(notes) == 0 {
		return c.JSON(res)
	}
	res.Content = notes[0].Content

	return c.JSON(res)
}
