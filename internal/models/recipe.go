package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

// RecipeStep descripes the step
type RecipeStep struct {
	Index      int
	Desciption string
	Tip        string
	ImagePath  string
}

// RecipeSteps explain each step to make an recipe
type RecipeSteps []RecipeStep

// Recipe explains how to cook
type Recipe struct {
	gorm.Model
	Title            string `gorm:"not null"`
	ImagePath        string
	Ease             string `gorm:"not null"`
	PreparationTime  int    `gorm:"not null"`
	RecipeCategoryID int
	Ingredients      []*Ingredient `gorm:"many2many:ingredient_quantity;"`
	WriterID         string
	Steps            RecipeSteps `gorm:"type:json"`
	Tags             []*Tag      `gorm:"many2many:recipe_tags;"`
	Clippers         []*User     `gorm:"many2many:recipe_clippers;joinForeignKey:RecipeID;"`
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
