package repository

import (
	"g-now/model"

	"gorm.io/gorm"
)

type IArticlesByCategoryRepository interface {
	GetArticlesByCategory(categoryID int, articles *[]model.Article) error
}

type articlesByCategoryRepository struct {
	db *gorm.DB
}

func NewArticlesByCategoryRepository(db *gorm.DB) IArticlesByCategoryRepository {
	return &articlesByCategoryRepository{db}
}

func (acr *articlesByCategoryRepository) GetArticlesByCategory(categoryID int, articles *[]model.Article) error {
	if err := acr.db.Joins("inner join article_category_maps acm on acm.article_id = articles.id").
		Joins("inner join categories ca on ca.id = acm.category_id").
		Where("ca.id = ?", categoryID).
		Preload("Tags").
		Find(articles).Error; err != nil {
		return err
	}
	return nil
}
