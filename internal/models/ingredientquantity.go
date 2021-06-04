package models

import (
	"gorm.io/datatypes"
)

// IngredientQuantity saves ingredient quantity in recipes
type IngredientQuantity struct {
	RecipeID     uint           `gorm:"primaryKey" json:"-"`
	IngredientID uint           `gorm:"primaryKey" json:"ingredient_id"`
	Ingredient   Ingredient     `json:"-"`
	Quantity     datatypes.JSON `gorm:"type:json;" sql:"type:json;" json:"quantity"`
}

// TableName overrides default tablename
func (i *IngredientQuantity) TableName() string {
	return "ingredient_quantity"
}
