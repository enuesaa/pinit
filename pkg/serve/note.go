package serve

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/usecase"
	"github.com/gofiber/fiber/v2"
)

type ListNotesItem struct {
	Id      string `json:"id"`
	Content string `json:"content"`
}

func (ctl *ServeCtl) ListNotes(c *fiber.Ctx) error {
	binderId, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	notes, err := usecase.ListBinderNotes(ctl.Repos, uint(binderId))
	if err != nil {
		return err
	}
	res := ApiListResponse[ListNotesItem]{
		Items: make([]ListNotesItem, 0),
	}
	for _, note := range notes {
		res.Items = append(res.Items, ListNotesItem{
			Id:      fmt.Sprintf("%d", note.ID),
			Content: note.Content,
		})
	}
	return c.JSON(res)
}
