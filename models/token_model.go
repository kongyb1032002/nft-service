package models

import "time"

type Token struct {
	ID        uint      `gorm:"primaryKey"`
	Owner     string    `gorm:"size:255;not null"`
	TokenID   int64     `gorm:"not null"`
	TokenURI  string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
