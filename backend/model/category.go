package model

type Category struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	CategoryName string `json:"category_name" gorm:"not null"`
}
