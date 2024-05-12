package usecase

import (
	"fmt"

	"github.com/enuesaa/pinit/pkg/service"
)

type ListNotesItem struct {
	Id      string `json:"id"`
	Content string `json:"content"`
}

func (ctl *ServeCtl) ListNotes(binderId int) ([]ListNotesItem, error) {
	res := NewServeListResponse[ListNotesItem]()
	noteSrv := service.NewNoteService(ctl.repos)
	notes, err := noteSrv.ListByBinderId(uint(binderId))
	if err != nil {
		return res.Items, err
	}
	for _, note := range notes {
		res.Items = append(res.Items, ListNotesItem{
			Id:      fmt.Sprintf("%d", note.ID),
			Content: note.Content,
		})
	}
	return res.Items, nil
}

type CreateNoteRequest struct {
	BinderId uint   `json:"binderId"`
	Content  string `json:"content"`
}

func (ctl *ServeCtl) CreateNote(req CreateNoteRequest) (ServeCreateResponse, error) {
	note := service.Note{
		BinderId:  req.BinderId,
		Comment:   "",
		Content:   req.Content,
		Publisher: "",
	}

	noteSrv := service.NewNoteService(ctl.repos)
	id, err := noteSrv.Create(note)
	if err != nil {
		return ServeCreateResponse{}, err
	}

	return ServeCreateResponse{Id: id}, nil
}
