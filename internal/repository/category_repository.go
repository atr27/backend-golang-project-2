package repository

import (
	"github.com/ispu-monitoring/backend/internal/model"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAll() ([]model.ISPUCategory, error) {
	var categories []model.ISPUCategory
	result := r.db.Order("min_value ASC").Find(&categories)
	return categories, result.Error
}

func (r *CategoryRepository) GetCategoryForISPU(ispu int) (*model.ISPUCategory, error) {
	var category model.ISPUCategory
	result := r.db.
		Where("min_value <= ?", ispu).
		Where("max_value IS NULL OR max_value >= ?", ispu).
		First(&category)
	return &category, result.Error
}
