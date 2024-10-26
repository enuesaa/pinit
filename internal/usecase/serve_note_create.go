package usecase

import (
	"github.com/enuesaa/pinit/internal/service"
	"github.com/gofiber/fiber/v2"
)

func (ctl *ServeCtl) NoteCreate(c *fiber.Ctx) error {
	type Request struct {
		Content  string `json:"content"`
	}
	type Response struct {
		Name string `json:"name"`
	}

	binderName := c.Params("binderName")

	var req Request
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	noteSrv := service.NewNoteService(binderName, ctl.repos)

	note := service.NoteCreation{
		Comment:   "",
		Content:   req.Content,
		Publisher: "",
	}
	name, err := noteSrv.Create(note)
	if err != nil {
		return err
	}
	return c.JSON(Response{Name: name})
}
