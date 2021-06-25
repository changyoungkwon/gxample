package database

import (
	"github.com/changyoungkwon/gxample/internal/logging"
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
func (s *RecipeStore) Add(r *models.Recipe) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		// tags insertion
		for _, t := range r.Tags {
			result := tx.Where(&models.Tag{Name: t.Name}).FirstOrCreate(&t)
			if result.Error != nil {
				return result.Error
			}
		}
		result := tx.Debug().First(&r.RecipeCategory, r.RecipeCategoryID)
		if result.Error != nil {
			return result.Error
		}
		// recipe insertions
		logging.Infof("%v", r)
		result = tx.Debug().Omit("Ingredients", "Tags.*").Create(&r)
		if result.Error != nil {
			return result.Error
		}
		// ingredient insertions
		for _, q := range r.IngredientQuantities {
			q.RecipeID = r.ID
			result := tx.Create(&q)
			if result.Error != nil {
				return result.Error
			}
		}
		result = tx.Debug().Preload("Ingredients").Joins("Writer").Joins("RecipeCategory").First(&r)
		if result.Error != nil {
			logging.Errorf("error during bring joined tables, %v", result.Error)
			return result.Error
		}
		return nil
	})
	return err
}

// Get get ingredient by key
func (s *RecipeStore) Get(key int) (*models.Recipe, error) {
	i := &models.Recipe{}
	var iqs []*models.IngredientQuantity
	if result := s.db.Debug().Preload("Ingredients").Joins("RecipeCategory").Joins("Writer").First(&i, key); result.Error != nil {
		return nil, result.Error
	}
	result := s.db.Debug().Where("recipe_id = ?", i.ID).Find(&iqs)
	if result.Error != nil {
		return nil, result.Error
	}
	i.IngredientQuantities = iqs
	return i, nil
}

// List list all ingredients
func (s *RecipeStore) List() ([]models.Recipe, error) {
	var recipes []models.Recipe
	var iqs []*models.IngredientQuantity
	result := s.db.Debug().Preload("Ingredients").Joins("RecipeCategory").Joins("Writer").Find(&recipes)
	if result.Error != nil {
		return nil, result.Error
	}
	for i, r := range recipes {
		result = s.db.Debug().Where("recipe_id = ?", r.ID).Find(&iqs)
		if result.Error != nil {
			return nil, result.Error
		}
		recipes[i].IngredientQuantities = iqs
	}
	return recipes, nil
}
