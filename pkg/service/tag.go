package service

import (
	"time"
)

type Tag struct {
	ID           uint    `gorm:"primaryKey"`
	NoteId       uint    `gorm:"not null"`
	Name         string  `gorm:"type:varchar(255)"`
	Value        string  `gorm:"type:varchar(255)"`
	CreatedAt    time.Time `gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt    time.Time `gorm:"type:timestamp;not null;default:current_timestamp on update current_timestamp"`
}
