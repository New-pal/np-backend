package category

import "gorm.io/gorm"

type CategoryService struct {
	db *gorm.DB
}

func (s *CategoryService) create(name string, desc string, code string, order uint) (*Category, error) {
	cat := Category{Name: name, Description: desc, Code: code, Order: order}
	res := s.db.Create(&cat)
	return &cat, res.Error
}

type SubCategoryService struct {
	db *gorm.DB
}

func (s *SubCategoryService) create(
	categoryId uint,
	name string,
	desc string,
	code string,
	order uint,
) (*Subcategory, error) {
	subCat := Subcategory{CategoryId: categoryId, Name: name, Description: desc, Code: code, Order: order}
	res := s.db.Create(&subCat)
	return &subCat, res.Error
}

func NewCategoryService(db *gorm.DB) *CategoryService {
	return &CategoryService{db: db}
}
func NewSubcategoryService(db *gorm.DB) *SubCategoryService {
	return &SubCategoryService{db: db}
}
