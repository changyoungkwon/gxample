package models

import (
	"gorm.io/datatypes"
)

// IngredientQuantity saves ingredient quantity in recipes
type IngredientQuantity struct {
	RecipeID     int            `gorm:"primaryKey"`
	IngredientID int            `gorm:"primaryKey"`
	Quantity     datatypes.JSON `gorm:"type:json;" sql:"type:json;"`
}

// TableName overrides default tablename
func (i *IngredientQuantity) TableName() string {
	return "ingredient_quantity"
}
