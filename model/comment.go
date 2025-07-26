package model

import "time"

type Comment struct {
	ID        string `gorm:"primaryKey;size:12"`
	PostID    string `gorm:"size:36;not null"`
	UserID    string `gorm:"size:36;not null"`
	Nickname  string `gorm:"size:12;not null"`
	Content   string `gorm:"size:100;not null"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
}
