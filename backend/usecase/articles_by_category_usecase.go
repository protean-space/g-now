package usecase

import (
	"g-now/model"
	"g-now/repository"
)

type IArticlesByCategoryUsecase interface {
	GetArticlesByCategory(categoryID int) ([]model.Article, error)
}

type articlesByCategoryUsecase struct {
	acr repository.IArticlesByCategoryRepository
}

func NewArticlesByCategoryUsecase(acr repository.IArticlesByCategoryRepository) IArticlesByCategoryUsecase {
	return &articlesByCategoryUsecase{acr}
}

func (acu *articlesByCategoryUsecase) GetArticlesByCategory(categoryID int) ([]model.Article, error) {
	var articles []model.Article
	err := acu.acr.GetArticlesByCategory(categoryID, &articles)
	if err != nil {
		return nil, err
	}
	return articles, nil
}
