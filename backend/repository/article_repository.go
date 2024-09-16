package repository

import (
	"g-now/model"

	"gorm.io/gorm"
)

type IArticleRepository interface {
	GetAllArticles(articles *[]model.Article) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) IArticleRepository {
	return &articleRepository{db}
}

func (ar *articleRepository) GetAllArticles(article *[]model.Article) error {
	if err := ar.db.Preload("Categories").Preload("Tags").Find(article).Error; err != nil {
		return err
	}
	return nil
}
