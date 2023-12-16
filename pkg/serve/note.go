package serve

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/usecase"
	"github.com/gofiber/fiber/v2"
)

type ListNotesResponseItemNote struct {
	Id         string `json:"id"`
	Content    string `json:"content"`
}
type ListNotesResponse struct {
	Items []ListNotesResponseItemNote `json:"items"`
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
	res := ListNotesResponse {
		Items: make([]ListNotesResponseItemNote, 0),
	}
	for _, note := range notes {
		res.Items = append(res.Items, ListNotesResponseItemNote{
			Id: fmt.Sprintf("%d", note.ID),
			Content: note.Content,
		})
	}
	return c.JSON(res)
}
