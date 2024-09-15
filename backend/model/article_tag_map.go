package model

type ArticleTagMap struct {
	ID        uint    `gorm:"primaryKey"`
	Article   Article `gorm:"not null;foreignKey:ArticleId"`
	ArticleId uint    `gorm:"not null"`
	Tag       Tag     `gorm:"not null;foreignKey:TagId"`
	TagId     uint    `gorm:"not null"`
}
