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
	note := service.Note{
		Comment:   "",
		Content:   req.Content,
		Publisher: "",
	}

	noteSrv := service.NewNoteService(binderName, ctl.repos)
	name, err := noteSrv.Create(note)
	if err != nil {
		return err
	}
	return c.JSON(Response{Name: name})
}
