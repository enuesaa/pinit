package usecase

import (
	"github.com/enuesaa/pinit/internal/service"
	"github.com/gofiber/fiber/v2"
)

func (ctl *ServeCtl) NoteList(c *fiber.Ctx) error {
	binderName := c.Params("binderName")

	type ListNotesItem struct {
		Id      string `json:"id"`
		Content string `json:"content"`
	}

	res := NewServeListResponse[ListNotesItem]()

	noteSrv := service.NewNoteService(binderName, ctl.repos)
	notes, err := noteSrv.List()
	if err != nil {
		return err
	}
	for _, note := range notes {
		res.Items = append(res.Items, ListNotesItem{
			Id:      note.NoteName,
			Content: note.Content,
		})
	}

	return c.JSON(res)
}
