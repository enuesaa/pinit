package usecase

import (
	"github.com/enuesaa/pinit/internal/service"
	"github.com/gofiber/fiber/v2"
)

func (ctl *ServeCtl) BinderUpdate(c *fiber.Ctx) error {
	type Request struct {
		Content string `json:"content"`
	}
	type Response struct{
		Name string `json:"name"`
	}

	binderName := c.Params("binderName")

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	noteSrv := service.NewNoteService(binderName, ctl.repos)

	histories, err := noteSrv.List()
	if err != nil || len(histories) > 0 {
		if histories[0].Content == req.Content {
			// Do not create a new note if the content has not been updated.
			return c.JSON(Response{ Name: binderName })
		}
	}

	note := service.NoteCreation{
		Comment: "",
		Content: req.Content,
		Publisher: "",
	}
	if _, err := noteSrv.Create(note); err != nil {
		return err
	}

	return c.JSON(Response{ Name: binderName })
}
