package category

import (
	"errors"
	"github.com/New-pal/np-backend/core"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func (r *CategoryRepository) getAll() ([]*Category, error) {
	var categories []*Category
	err := r.db.Order("categories.order").Find(&categories).Error
	return categories, err
}
func (r *CategoryRepository) getById(id string) (*Category, error) {
	var category *Category
	err := r.db.First(&category, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, core.ErrRecordNotFound
	}
	return category, err
}

type SubcategoryRepository struct {
	db *gorm.DB
}

func (r *SubcategoryRepository) getAll() ([]*Subcategory, error) {
	var subcategories []*Subcategory
	err := r.db.Order("subcategories.order").Find(&subcategories).Error
	return subcategories, err
}
func (r *SubcategoryRepository) getById(id string) (*Subcategory, error) {
	var subcategory *Subcategory
	err := r.db.First(&subcategory, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, core.ErrRecordNotFound
	}
	return subcategory, err
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}
func NewSubcategoryRepository(db *gorm.DB) *SubcategoryRepository {
	return &SubcategoryRepository{db: db}
}
