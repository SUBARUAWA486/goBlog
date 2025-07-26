package model

import "time"

type Post struct { 
	ID string `gorm:"primaryKey;size:36"`
	Title string `gorm:"not null;size:20"`
	Content string `gorm:"not null;size:300"`
	Cover string `gorm:"not null"`
	Views int `gorm:"default:0"`
	Stars int `gorm:"default:0"`
	UserID string `gorm:"not null;size:36"`
	Nickname string `gorm:"not null;size:12"`
	Avatar string
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
}