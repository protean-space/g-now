package repository

import (
	"g-now/model"

	"gorm.io/gorm"
)

type IArticlesByTagRepository interface {
	GetArticlesByTag(tagName string, articles *[]model.Article) error
}

type articlesByTagRepository struct {
	db *gorm.DB
}

func NewArticlesByTagRepository(db *gorm.DB) IArticlesByTagRepository {
	return &articlesByTagRepository{db}
}

func (atc *articlesByTagRepository) GetArticlesByTag(tagName string, articles *[]model.Article) error {
	if err := atc.db.
		Table("articles ar").
		Select("ar.*, tags.*").
		Joins("left join article_tag_maps atm on atm.article_id = ar.id").
		Joins("left join tags on tags.id = atm.tag_id").
		Where("tags.tag_name LIKE ?", "%"+tagName+"%").
		Order("ar.id desc"). // idで降順にソート
		Preload("Tags").     // Tagsリレーションをプリロード
		Find(articles).Error; err != nil {
		return err
	}
	return nil
}
