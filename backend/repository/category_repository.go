package repository

import (
	"g-now/model"

	"gorm.io/gorm"
)

type ICategoryRepository interface {
	GetAllCategories(categories *[]model.Category) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &categoryRepository{db}
}

func (cr *categoryRepository) GetAllCategories(categories *[]model.Category) error {
	if err := cr.db.Order("id").Find(categories).Error; err != nil {
		return err
	}
	return nil
}
