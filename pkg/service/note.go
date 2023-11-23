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

func (srv *NoteService) IsTabelExist() (bool, error) {
	return srv.repos.Database.IsTableExist(&Note{})
}

func (srv *NoteService) CreateTable() error {
	return srv.repos.Database.CreateTable(&Note{})
}

func (srv *NoteService) List() []*Note {
	notes := make([]*Note, 0)
	srv.repos.Database.ListAll(&notes)
	return notes
}

func (srv *NoteService) Get(id uint) (*Note, error) {
	note := Note{
		ID: id,
	}
	if err := srv.repos.Database.WhereFirst(&note); err != nil {
		return nil, err
	}
	return &note, nil
}

func (srv *NoteService) Create(note Note) error {
	if err := srv.repos.Database.Create(&note); err != nil {
		return err
	}
	return nil
}

func (srv *NoteService) RunCreatePrompt() (*Note, error) {
	note := Note{}
	return srv.RunEditPrompt(note)
}

func (srv *NoteService) RunEditPrompt(note Note) (*Note, error) {
	content, err := srv.repos.Prompt.Ask("Content", note.Content)
	if err != nil {
		return nil, err
	}
	note.Content = content

	return &note, nil
}

func (srv *NoteService) Update(note Note) error {
	if err := srv.repos.Database.Update(&note); err != nil {
		return err
	}
	return nil
}

func (srv *NoteService) Delete(id uint) error {
	note, err := srv.Get(id)
	if err != nil {
		return err
	}
	if err := srv.repos.Database.Delete(&note); err != nil {
		return err
	}
	return nil
}
