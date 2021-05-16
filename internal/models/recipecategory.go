package models

import "gorm.io/gorm"

// RecipeCategory application struct
type RecipeCategory struct {
	gorm.Model
	Name      string `gorm:"unique; not null"`
	ImagePath string
	Recipes   []Recipe
}
