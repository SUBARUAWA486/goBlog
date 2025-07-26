package model

import "time"

type Stars struct {
	ID        string    `gorm:"primaryKey"`
	UserID    string    `gorm:"index;not null"`
	PostID    string    `gorm:"index;not null"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
}
