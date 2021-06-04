package models

import (
	"gorm.io/gorm"
)

// Ingredient saves available ingredients
type Ingredient struct {
	gorm.Model
	Name      string    `gorm:"unique;not null" json:"name"`
	ImagePath string    `json:"image_path"` // must be form of /static/image/ingredient/{id}/*.jpg
	Recipes   []*Recipe `gorm:"many2many:ingredient_quantity;" json:"recipes"`
}
