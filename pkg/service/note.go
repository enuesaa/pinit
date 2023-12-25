package service

import (
	"time"

	"github.com/enuesaa/pinit/pkg/repository"
)

type Note struct {
	ID        uint      `gorm:"primaryKey"`
	BinderId  uint      `gorm:"type:integer"`
	Publisher string    `gorm:"type:varchar(255)"`
	Comment   string    `gorm:"type:text"`
	Content   string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null;default:current_timestamp on update current_timestamp"`
}

type NoteService struct {
	repos repository.Repos
}

func NewNoteService(repos repository.Repos) *NoteService {
	return &NoteService{
		repos: repos,
	}
}

func (srv *NoteService) IsTableExist() (bool, error) {
	return srv.repos.Database.IsTableExist("notes")
}

func (srv *NoteService) CreateTable() error {
	return srv.repos.Database.CreateTable(&Note{})
}

func (srv *NoteService) List() ([]Note, error) {
	notes := make([]Note, 0)
	if err := srv.repos.Database.ListAll(&notes); err != nil {
		return notes, err
	}
	return notes, nil
}

func (srv *NoteService) Get(id uint) (Note, error) {
	var note Note
	if err := srv.repos.Database.WhereFirst(&note, "id = ?", id); err != nil {
		return note, err
	}
	return note, nil
}

// TODO return list response.
func (srv *NoteService) GetFirstByBinderId(binderId uint) (Note, error) {
	var note Note
	if err := srv.repos.Database.WhereFirst(&note, "binder_id = ?", binderId); err != nil {
		return note, err
	}
	return note, nil
}

// TODO: refactor
func (srv *NoteService) ListByBinderId(binderId uint) ([]Note, error) {
	notes, err := srv.List()
	if err != nil {
		return make([]Note, 0), err
	}
	list := make([]Note, 0)
	for _, note := range notes {
		if note.BinderId == binderId {
			list = append(list, note)
		}
	}
	return list, nil
}

func (srv *NoteService) Create(note *Note) error {
	return srv.repos.Database.Create(note)
}

func (srv *NoteService) RunPrompt(note *Note) error {
	content, err := srv.repos.Prompt.Ask("Content", note.Content)
	if err != nil {
		return err
	}
	note.Content = content
	return nil
}

func (srv *NoteService) Update(note *Note) error {
	return srv.repos.Database.Update(note)
}

func (srv *NoteService) Delete(id uint) error {
	note, err := srv.Get(id)
	if err != nil {
		return err
	}
	return srv.repos.Database.Delete(&note)
}
