package usecase

import (
	"github.com/enuesaa/pinit/internal/service"
	"github.com/gofiber/fiber/v2"
)

func (ctl *ServeCtl) NoteList(c *fiber.Ctx) error {
	type Item struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}

	binderName := c.Params("binderName")

	noteSrv := service.NewNoteService(binderName, ctl.repos)
	notes, err := noteSrv.List()
	if err != nil {
		return err
	}

	res := NewServeListResponse[Item]()
	for _, note := range notes {
		res.Items = append(res.Items, Item{
			Name: note.NoteName,
			Content: note.Content,
		})
	}
	return c.JSON(res)
}
