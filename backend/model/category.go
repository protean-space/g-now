package model

type Category struct {
	ID           uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	CategoryName string `json:"category_name" gorm:"not null"`
}
