package models

import (
	"gorm.io/gorm"
)

// Ingredient supports ingredient type
type Ingredient struct {
	gorm.Model
	Name          string
	ThumbnailPath *string
}
