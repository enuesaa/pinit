package service

import (
	"time"
)

type Tag struct {
	ID        uint      `gorm:"primaryKey"`
	NoteId    uint      `gorm:"not null"`
	Name      string    `gorm:"type:varchar(255)"`
	Value     string    `gorm:"type:varchar(255)"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null;default:current_timestamp on update current_timestamp"`
}

type TagService struct{}

func (srv *TagService) List() []Tag {
	return make([]Tag, 0)
}

func (srv *TagService) Get() (Tag, error) {
	return *new(Tag), nil
}

func (srv *TagService) Create() error {
	return nil
}

func (srv *TagService) Update() error {
	return nil
}

func (srv *TagService) Remove() error {
	return nil
}
