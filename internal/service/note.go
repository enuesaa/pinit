package service

import (
	"fmt"
	"time"

	"github.com/enuesaa/pinit/internal/repository"
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
	repos repository.Repos
}

func NewNoteService(repos repository.Repos) *NoteService {
	return &NoteService{
		repos: repos,
	}
}

func (srv *NoteService) List() ([]Note, error) {
	list := []Note{}
	if err := srv.repos.Db.List("BinderName", "bidnername", &list); err != nil {
		return list, err
	}
	return list, nil
}

func (srv *NoteService) Get(name string) (Note, error) {
	data := Note{}
	if err := srv.repos.Db.Get("BinderName", "bidnername", &data); err != nil {
		return data, err
	}
	return data, nil
}

// TODO return list response.
func (srv *NoteService) GetFirstByBinderId(binderId uint) (Note, error) {
	return srv.Get("a")
}

func (srv *NoteService) ListByBinderId(binderId uint) ([]Note, error) {
	return srv.List()
}

func (srv *NoteService) Create(note Note) (uint, error) {
	return 0, fmt.Errorf("not implemented")
}

func (srv *NoteService) Update(note Note) error {
	return fmt.Errorf("not implemented")
}

func (srv *NoteService) Delete(id uint) error {
	return fmt.Errorf("not implemented")
}

func (srv *NoteService) DeleteByBinderId(id uint) error {
	return fmt.Errorf("not implemented")
}
