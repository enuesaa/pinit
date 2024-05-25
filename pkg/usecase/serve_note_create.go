package usecase

import (
	"github.com/enuesaa/pinit/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func (ctl *ServeCtl) NoteCreate(c *fiber.Ctx) error {
	type CreateNoteRequest struct {
		BinderId uint   `json:"binderId"`
		Content  string `json:"content"`
	}

	var req CreateNoteRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	note := service.Note{
		BinderId:  req.BinderId,
		Comment:   "",
		Content:   req.Content,
		Publisher: "",
	}

	noteSrv := service.NewNoteService(ctl.repos)
	id, err := noteSrv.Create(note)
	if err != nil {
		return err
	}

	return c.JSON(ServeCreateResponse{Id: id})
}
