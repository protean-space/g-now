package model

type Tag struct {
	ID        uint    `gorm:"primaryKey"`
	TagName   string  `gorm:"not null"`
	Article   Article `gorm:"not null;foreignKey:ArticleId"`
	ArticleId uint    `gorm:"not null"`
}
