package database

import (
	"fmt"

	"github.com/changyoungkwon/gxample/internal/models"
	"gorm.io/gorm"
)

// RecipeCategoryStore is the wrapper class for the database
type RecipeCategoryStore struct {
	db *gorm.DB
}

// NewRecipeCategoryStore provide store abstraction
func NewRecipeCategoryStore(db *gorm.DB) *RecipeCategoryStore {
	return &RecipeCategoryStore{
		db: db,
	}
}

// Add create recipe, update i with automatically received value, then return err
func (s *RecipeCategoryStore) Add(i *models.RecipeCategory) error {
	result := s.db.Create(&i)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Get get recipe by key
func (s *RecipeCategoryStore) Get(key int) (*models.RecipeCategory, error) {
	i := &models.RecipeCategory{}
	if result := s.db.First(&i, key); result.Error != nil {
		return nil, result.Error
	}
	return i, nil
}

// List list all recipes
func (s *RecipeCategoryStore) List(names []string) ([]models.RecipeCategory, error) {
	var rcs []models.RecipeCategory
	if names == nil {
		names = []string{}
	}

	// iterate over list of names
	query := s.db
	for _, name := range names {
		query = query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	result := query.Find(&rcs)
	if result.Error != nil {
		return nil, result.Error
	}

	return rcs, nil
}
