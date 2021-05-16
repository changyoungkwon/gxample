package models

import "gorm.io/gorm"

// Tag to specify the user
type Tag struct {
	gorm.Model
	Name    string    `gorm:"unique;not null"`
	Recipes []*Recipe `gorm:"many2many:recipe_tags"`
}
