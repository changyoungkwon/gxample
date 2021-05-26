package database

import (
	"github.com/changyoungkwon/gxample/internal/models"
	"gorm.io/gorm"
)

// RecipeStore is the wrapper class for the database
type RecipeStore struct {
	db *gorm.DB
}

// NewRecipeStore provide store abstraction
func NewRecipeStore(db *gorm.DB) *RecipeStore {
	return &RecipeStore{
		db: db,
	}
}

// Add create ingredient, update i with automatically received value, then return err
func (s *RecipeStore) Add(i *models.Recipe) error {
	result := s.db.Create(&i)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Get get ingredient by key
func (s *RecipeStore) Get(key int) (*models.Recipe, error) {
	i := &models.Recipe{}
	if result := s.db.First(&i, key); result.Error != nil {
		return nil, result.Error
	}
	return i, nil
}

// List list all ingredients
func (s *RecipeStore) List() ([]models.Recipe, error) {
	var recipes []models.Recipe
	result := s.db.Find(&recipes)
	if result.Error != nil {
		return nil, result.Error
	}
	return recipes, nil
}
