package repository

import (
	"g-now/model"

	"gorm.io/gorm"
)

type ICategoryRepository interface {
	GetAllCategories(category *[]model.Category) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &categoryRepository{db}
}

func (cu *categoryRepository) GetAllCategories(categories *[]model.Category) error {
	if err := cu.db.Order("id").Find(categories).Error; err != nil {
		return err
	}
	return nil
}
