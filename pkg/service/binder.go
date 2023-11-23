package service

import (
	"time"
)

type Binder struct {
	ID         uint      `gorm:"primaryKey"`
	Name       string    `gorm:"type:varchar(255)"`
	Category   string    `gorm:"type:varchar(255)"`
	ArchivedAt time.Time `gorm:"type:timestamp"` 
	CreatedAt  time.Time `gorm:"type:timestamp;not null;default:current_timestamp"`
	UpdatedAt  time.Time `gorm:"type:timestamp;not null;default:current_timestamp on update current_timestamp"`
}