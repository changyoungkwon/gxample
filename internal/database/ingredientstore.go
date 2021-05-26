package database

import (
	"github.com/changyoungkwon/gxample/internal/models"
	"gorm.io/gorm"
)

// IngredientStore is the wrapper class for the database
type IngredientStore struct {
	db *gorm.DB
}

// NewIngredientStore provide store abstraction
func NewIngredientStore(db *gorm.DB) *IngredientStore {
	return &IngredientStore{
		db: db,
	}
}

// Add create ingredient, update i with automatically received value, then return err
func (s *IngredientStore) Add(i *models.Ingredient) error {
	result := s.db.Create(&i)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Get get ingredient by key
func (s *IngredientStore) Get(key int) (*models.Ingredient, error) {
	i := &models.Ingredient{}
	if result := s.db.First(&i, key); result.Error != nil {
		return nil, result.Error
	}
	return i, nil
}

// List list all ingredients
func (s *IngredientStore) List() ([]models.Ingredient, error) {
	var ingredients []models.Ingredient
	result := s.db.Find(&ingredients)
	if result.Error != nil {
		return nil, result.Error
	}
	return ingredients, nil
}
