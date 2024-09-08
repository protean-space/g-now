package model

import "time"

type Article struct {
	ID                uint   `gorm:"primaryKey"`
	Title             string `gorm:"not null"`
	Contents          string `gorm:"not null"`
	ArticleUrl        string
	ArticleImageUrl   string
	PageView          uint
	PayloadJson       string
	SourcePublishedAt time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
