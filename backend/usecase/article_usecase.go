package usecase

import (
	"g-now/model"
	"g-now/repository"
)

type IArticleUsecase interface {
	GetAllArticles() ([]model.Article, error)
}

type articleUsecase struct {
	ar repository.IArticleRepository
}

func NewArticleUsecase(ar repository.IArticleRepository) IArticleUsecase {
	return &articleUsecase{ar}
}

func (au *articleUsecase) GetAllArticles() ([]model.Article, error) {
	articles := []model.Article{}
	if err := au.ar.GetAllArticles(&articles); err != nil {
		return nil, err
	}

	return articles, nil
}
