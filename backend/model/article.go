package model

import "time"

type Article struct {
	ID                uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UID               string `json:"uid" gorm:"not null;uniqueIndex"`
	Title             string `json:"title" gorm:"not null"`
	Contents          string `json:"contents" gorm:"not null"`
	ArticleUrl        string
	ArticleImageUrl   string
	PageView          uint
	PayloadJson       string
	SourcePublishedAt time.Time `json:"source_published_at" gorm:"not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Categories        []Category `json:"categories" gorm:"many2many:article_category_maps"`
	Tags              []Tag      `json:"tags" gorm:"many2many:article_tag_maps"`
}
