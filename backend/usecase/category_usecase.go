package usecase

import (
	"g-now/model"
	"g-now/repository"
)

type ICategoryUsecase interface {
	GetAllCategories() ([]model.Category, error)
}

type categoryUsecase struct {
	cr repository.ICategoryRepository
}

func NewCategoryUsecase(cr repository.ICategoryRepository) ICategoryUsecase {
	return &categoryUsecase{cr}
}

func (cu *categoryUsecase) GetAllCategories() ([]model.Category, error) {
	categories := []model.Category{}
	if err := cu.cr.GetAllCategories(&categories); err != nil {
		return nil, err
	}

	return categories, nil
}
