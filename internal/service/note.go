package service

import (
	"fmt"
	"time"

	"github.com/enuesaa/pinit/internal/repository"
	"github.com/google/uuid"
)

type Note struct {
	BinderName string   `dynamo:"BinderName"`
	NoteName  string    `dynamo:"NoteName"`

	Publisher string    `dynamo:"Publisher"`
	Comment   string    `dynamo:"Comment"`
	Content   string    `dynamo:"Content"`
	CreatedAt time.Time `dynamo:"CreatedAt"`
	UpdatedAt time.Time `dynamo:"UpdatedAt"`
}

type NoteService struct {
	binderName string
	repos repository.Repos
}

func NewNoteService(binderName string, repos repository.Repos) *NoteService {
	return &NoteService{
		binderName: binderName,
		repos: repos,
	}
}

func (srv *NoteService) List() ([]Note, error) {
	list := []Note{}
	if err := srv.repos.Db.List(srv.binderName, &list); err != nil {
		return list, err
	}
	return list, nil
}

func (srv *NoteService) Get(name string) (Note, error) {
	data := Note{}
	if err := srv.repos.Db.Get(srv.binderName, name, &data); err != nil {
		return data, err
	}
	return data, nil
}

func (srv *NoteService) Create(note Note) (string, error) {
	note.BinderName = srv.binderName
	note.NoteName = uuid.NewString()

	if err := srv.repos.Db.Put(note); err != nil {
		return "", err
	}
	return note.NoteName, nil
}

func (srv *NoteService) Update(note Note) (string, error) {
	if note.NoteName == "" {
		return "", fmt.Errorf("name is required")
	}
	if err := srv.repos.Db.Put(note); err != nil {
		return "", err
	}
	return note.NoteName, nil
}

func (srv *NoteService) Delete(noteName string) error {
	return srv.repos.Db.Delete(srv.binderName, noteName)
}

func (srv *NoteService) DeleteAllInBinder() error {
	list := []Note{}
	if err := srv.repos.Db.List(srv.binderName, &list); err != nil {
		return err
	}
	for _, data := range list {
		if err := srv.repos.Db.Delete(srv.binderName, data.NoteName); err != nil {
			return err
		}
	}
	return nil
}
