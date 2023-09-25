package service

import (
	"time"

	"github.com/enuesaa/pinit/pkg/repository"
)

type Note struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(255)"`
	Content   string    `gorm:"type:text"`
	Comment   string    `gorm:"type:text"`
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

func (srv *NoteService) List() []*Note {
	notes := make([]*Note, 0)
	srv.repos.Database.ListAll(&notes)
	return notes
}

func (srv *NoteService) Get(name string) (*Note, error) {
	note := Note{
		Name: name,
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

func (srv *NoteService) Update(note Note) error {
	if err := srv.repos.Database.Update(&note); err != nil {
		return err
	}
	return nil
}

func (srv *NoteService) Delete(name string) error {
	note, err := srv.Get(name)
	if err != nil {
		return err
	}
	if err := srv.repos.Database.Delete(&note); err != nil {
		return err
	}
	return nil
}
