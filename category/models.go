package category

import "gorm.io/gorm"

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Category{}, &Subcategory{})
}

type Category struct {
	gorm.Model
	Name        string `json:"name" gorm:"size:50;not null;index;unique"`
	Description string `json:"description" gorm:"size:255;"`
	Code        string `json:"code" gorm:"size:50;not null;index;unique"`
	Order       uint   `json:"order" gorm:"default=1"`
}

type Subcategory struct {
	gorm.Model
	CategoryId  uint
	Category    Category `json:"category" gorm:"constraint: OnDelete:CASCADE; default:null"`
	Name        string   `json:"name" gorm:"size:50; not null; index; unique"`
	Description string   `json:"description" gorm:"size:255"`
	Code        string   `json:"code" gorm:"size:50;not null; index;unique"`
	Order       uint     `json:"order" gorm:"default=1"`
}
