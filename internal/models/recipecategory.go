package models

import "gorm.io/gorm"

// RecipeCategory application struct
type RecipeCategory struct {
	gorm.Model
	Name      string   `gorm:"unique; not null" json:"name"`
	ImagePath string   `json:"image_path"`
	Recipes   []Recipe `json:"-"`
}
