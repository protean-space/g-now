package model

type ArticleCategoryMap struct {
	ID         uint     `gorm:"primaryKey"`
	Article    Article  `gorm:"foreignKey:ArticleId`
	ArticleId  uint     `gorm:"not null"`
	Category   Category `gorm:"not null;foreignKey:CategoryId"`
	CategoryId uint     `gorm:"not null"`
}
