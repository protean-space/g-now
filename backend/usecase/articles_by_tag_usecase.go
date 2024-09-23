package usecase

import (
	"g-now/model"
	"g-now/repository"
)

type IArticlesByTagUsecase interface {
	GetArticlesByTag(tagName string) ([]model.Article, error)
}

type articlesByTagUsecase struct {
	atr repository.IArticlesByTagRepository
}

func NewArticlesByTagUsecase(atr repository.IArticlesByTagRepository) IArticlesByTagUsecase {
	return &articlesByTagUsecase{atr}
}

func (atu *articlesByTagUsecase) GetArticlesByTag(tagName string) ([]model.Article, error) {
	var articles []model.Article
	err := atu.atr.GetArticlesByTag(tagName, &articles)
	if err != nil {
		return nil, err
	}
	return articles, nil
}
