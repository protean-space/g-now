package model

type Tag struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	TagName string `json:"tag_name" gorm:"not null"`
}
