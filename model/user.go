package model

import "time"

type User struct {
	ID        string `gorm:"primaryKey;size:36"`
	Account   string `gorm:"uniqueIndex;not null;size:9"`
	Password  string `gorm:"not null;size:12"`
	Nickname  string `gorm:"unique;not null;size:12"`
	Avatar    string `gorm:"not null"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
}
