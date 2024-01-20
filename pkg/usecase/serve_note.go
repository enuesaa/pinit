package usecase

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type ListNotesItem struct {
	Id      string `json:"id"`
	Content string `json:"content"`
}

func (ctl *ServeCtl) ListNotes(c *fiber.Ctx) error {
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

type CreateNoteRequest struct {
	BinderId uint   `json:"binderId"`
	Content  string `json:"content"`
}

func (ctl *ServeCtl) CreateNote(c *fiber.Ctx) error {
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
	id, err := noteSrv.Create(note);
	if err != nil {
		return err
	}

	return c.JSON(ServeCreateResponse{ Id: id })
}
