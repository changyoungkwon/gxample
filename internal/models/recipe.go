package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

// RecipeStep descripes the step
type RecipeStep struct {
	Index       int    `json:"index"`
	Description string `json:"description"`
	Tip         string `json:"tip"`
	ImagePath   string `json:"image_path"`
}

// RecipeSteps explain each step to make an recipe
type RecipeSteps []RecipeStep

// Recipe explains how to cook
type Recipe struct {
	gorm.Model
	Title                string                `gorm:"not null" json:"title"`
	ImagePath            string                `json:"image_path"`
	Ease                 string                `gorm:"not null" json:"ease"`
	PreparationTime      int                   `gorm:"not null" json:"preparation_time"`
	RecipeCategoryID     int                   `json:"-"`
	RecipeCategory       RecipeCategory        `json:"recipe_category"`
	IngredientQuantities []*IngredientQuantity `gorm:"-" json:"ingredient_quantities"`
	Ingredients          []*Ingredient         `gorm:"many2many:ingredient_quantity;" json:"-"`
	WriterID             string                `json:"writer_id"`
	Steps                RecipeSteps           `gorm:"type:json" json:"steps"`
	Tags                 []*Tag                `gorm:"many2many:recipe_tags;constraint:OnDelete:CASCADE" json:"-"`
	Clippers             []*User               `gorm:"many2many:recipe_clippers;joinForeignKey:RecipeID;" json:"-"`
}

// Value for custom type
func (steps RecipeSteps) Value() (driver.Value, error) {
	if len(steps) == 0 {
		return nil, nil
	}
	bytes, err := json.Marshal(steps)
	return string(bytes), err
}

// Scan for custom type
func (steps *RecipeSteps) Scan(v interface{}) error {
	bytes, ok := v.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal json value: %s", v)
	}
	err := json.Unmarshal(bytes, &steps)
	return err
}
