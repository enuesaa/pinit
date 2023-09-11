package service

import (
	"time"
)

type Note struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(255)"`
	Content   string    `gorm:"type:text"`
	Comment   string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null;default:current_timestamp on update current_timestamp"`
}

type NoteService struct {}

func (srv *NoteService) list() []Note {
	return make([]Note, 0)
}

func (srv *NoteService) get() (Note, error) {
	return *new(Note), nil
}

func (srv *NoteService) create() error {
	return nil
}

func (srv *NoteService) update() error {
	return nil
}

func (srv *NoteService) remove() error {
	return nil
}
