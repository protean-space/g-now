package model

type ArticleCategoryMap struct {
	ID         uint     `gorm:"primaryKey;autoIncrement"`
	Article    Article  `gorm:"not null;foreignKey:ArticleId"`
	ArticleId  uint     `gorm:"not null;index:idx_article_category"`
	Category   Category `gorm:"not null;foreignKey:CategoryId"`
	CategoryId uint     `gorm:"not null;index:idx_article_category"`
}
