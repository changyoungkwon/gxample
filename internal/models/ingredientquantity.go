package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// IngredientQuantity saves ingredient quantity in recipes
type IngredientQuantity struct {
	Quantity     datatypes.JSON `gorm:"type:json;" sql:"type:json;" json:"quantity"`
	RecipeID     uint           `gorm:"primaryKey" json:"-"`
	IngredientID uint           `gorm:"primaryKey" json:"ingredient_id"`
	Ingredient   Ingredient     `json:"-"`
}

// TableName declares table name of ingredeint quantity
func (i *IngredientQuantity) TableName() string {
	return "ingredient_quantity"
}

// BeforeCreate is a hook
func (IngredientQuantity) BeforeCreate(db *gorm.DB) error {
	// ...
	return nil
}
