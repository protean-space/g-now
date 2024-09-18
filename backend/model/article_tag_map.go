package model

type ArticleTagMap struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	Article   Article `gorm:"not null;foreignKey:ArticleId"`
	ArticleId uint    `gorm:"not null;index:idx_article_tag"`
	Tag       Tag     `gorm:"not null;foreignKey:TagId"`
	TagId     uint    `gorm:"not null;index:idx_article_tag"`
}
