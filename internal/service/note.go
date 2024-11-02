package service

import (
	"time"

	"github.com/enuesaa/pinit/internal/repository"
	"github.com/oklog/ulid/v2"
)

type Note struct {
	InternalBinderName string   `dynamo:"BinderName"`
	NoteName  string    `dynamo:"NoteName"`

	Publisher string    `dynamo:"Publisher"`
	Comment   string    `dynamo:"Comment"`
	Content   string    `dynamo:"Content"`
	CreatedAt time.Time `dynamo:"CreatedAt"`
	UpdatedAt time.Time `dynamo:"UpdatedAt"`
}
type NoteCreation struct {
	Publisher string    `dynamo:"Publisher"`
	Comment   string    `dynamo:"Comment"`
	Content   string    `dynamo:"Content"`
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

func (srv *NoteService) Create(creation NoteCreation) (string, error) {
	note := Note{
		InternalBinderName: srv.binderName,
		NoteName: ulid.Make().String(),
		Content: creation.Content,
		Comment: creation.Comment,
		Publisher: creation.Publisher,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := srv.repos.Db.Put(note); err != nil {
		return "", err
	}
	return note.NoteName, nil
}

func (srv *NoteService) Update(name string, creation NoteCreation) (string, error) {
	note, err := srv.Get(name)
	if err != nil {
		return "", err
	}
	note.Content = creation.Content
	note.Comment = creation.Comment
	note.UpdatedAt = time.Now()
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
