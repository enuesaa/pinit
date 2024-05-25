package usecase

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/service"
	"github.com/gofiber/fiber/v2"
)

func (ctl *ServeCtl) NoteList(c *fiber.Ctx) error {
	type ListNotesItem struct {
		Id      string `json:"id"`
		Content string `json:"content"`
	}

	res := NewServeListResponse[ListNotesItem]()
	binderId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	noteSrv := service.NewNoteService(ctl.repos)
	notes, err := noteSrv.ListByBinderId(uint(binderId))
	if err != nil {
		return err
	}
	for _, note := range notes {
		res.Items = append(res.Items, ListNotesItem{
			Id:      fmt.Sprintf("%d", note.ID),
			Content: note.Content,
		})
	}
	return c.JSON(res)
}
