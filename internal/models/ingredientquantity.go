package models

import (
	"gorm.io/datatypes"
)

// IngredientQuantity saves ingredient quantity in recipes
type IngredientQuantity struct {
	RecipeID     uint       `gorm:"primaryKey" json:"-"`
	IngredientID uint       `gorm:"primaryKey" json:"ingredient_id"`
	Ingredient   Ingredient `json:"-"`
	AnyIdea      string
	Quantity     datatypes.JSON `gorm:"type:json;" sql:"type:json;" json:"quantity"`
}
