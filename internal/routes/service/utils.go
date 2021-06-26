package service

import (
	"encoding/json"
	"math/rand"
	"sync"
	"time"

	"github.com/changyoungkwon/gxample/internal/models"
	"gorm.io/gorm"
)

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	once        sync.Once
)

// toRecipe generates recipes based on request
func dtoToRecipe(body *recipeRequestJSON, userID string) (*models.Recipe, error) {
	// dto to entities
	// tags
	tags := []*models.Tag{}
	for _, t := range body.Tags {
		tags = append(tags, &models.Tag{Name: t})
	}
	// quantities
	iqs, err := dtoToIQuantities(body.IngredientQuantities, body.ID)
	if err != nil {
		return nil, err
	}

	return &models.Recipe{
		Model:                gorm.Model{ID: body.ID},
		Title:                body.Title,
		Ease:                 body.Ease,
		PreparationTime:      body.PreparationTime,
		RecipeCategoryID:     body.RecipeCategoryID,
		WriterID:             userID,
		IngredientQuantities: iqs,
		Steps:                body.Steps,
		Tags:                 tags,
	}, nil
}

// FromRecipe generates based on recipes
func dtoFromRecipe(r *models.Recipe) (*RecipeResponse, error) {
	// tags
	tags := make([]string, 0, len(r.Tags))
	for _, t := range r.Tags {
		tags = append(tags, t.Name)
	}
	// ingerdient quantities
	quantities, err := dtoFromIQuantities(r)
	if err != nil {
		return nil, err
	}
	return &RecipeResponse{
		ID:                   r.ID,
		CreatedAt:            r.CreatedAt,
		UpdatedAt:            r.UpdatedAt,
		Title:                r.Title,
		ImagePath:            r.ImagePath,
		Ease:                 r.Ease,
		PreparationTime:      r.PreparationTime,
		RecipeCategory:       *dtoFromRecipeCategory(&r.RecipeCategory),
		IngredientQuantities: quantities,
		Writer:               *dtoToUserResponse(r.Writer),
		Steps:                r.Steps,
		IsClipped:            false,
		Tags:                 tags,
	}, nil
}

// thumbFromRecipe thumbnail generates based on recipe
func thumbFromRecipe(r *models.Recipe) *RecipeThumbResponse {
	return &RecipeThumbResponse{
		ID:              r.ID,
		Title:           r.Title,
		ImagePath:       r.ImagePath,
		Ease:            r.Ease,
		PreparationTime: r.PreparationTime,
		Writer:          *dtoToUserResponse(r.Writer),
		IsClipped:       false,
	}
}

// ToIngredient binds request to entity
func dtoToIngredient(i *ingredientRequestJSON) *models.Ingredient {
	return &models.Ingredient{
		Name: i.Name,
	}
}

// FromIngredient binds entity to response
func dtoFromIngredient(i *models.Ingredient) *IngredientResponse {
	return &IngredientResponse{
		ID:        i.ID,
		Name:      i.Name,
		ImagePath: i.ImagePath,
	}
}

// toIngredientQuantity binds request to entity
func dtoToIQuantities(rs []*IngredientQuantityRequest, recipeID uint) ([]*models.IngredientQuantity, error) {
	var quantities []*models.IngredientQuantity
	for _, r := range rs {
		bytes, err := json.Marshal(r.Quantity)
		if err != nil {
			return nil, err
		}
		quantities = append(quantities, &models.IngredientQuantity{
			RecipeID:     recipeID,
			IngredientID: r.IngredientID,
			Quantity:     bytes,
		})
	}
	return quantities, nil
}

// FromIngredientQuantity binds entity to response
func dtoFromIQuantities(rcp *models.Recipe) ([]*IngredientQuantityResponse, error) {
	var quantities []*IngredientQuantityResponse
	qs := rcp.IngredientQuantities
	for _, q := range qs {
		var res IngredientQuantityResponse
		err := json.Unmarshal(q.Quantity, &res.Quantity)
		if err != nil {
			return nil, err
		}
		for _, i := range rcp.Ingredients {
			if q.IngredientID == i.ID {
				res.Ingredient = *dtoFromIngredient(i)
				quantities = append(quantities, &res)
				break
			}
		}
	}
	return quantities, nil
}

// ToRecipeCategory binds request to entity
func dtoToRecipeCategory(r *recipeCategoryRequestJSON) *models.RecipeCategory {
	return &models.RecipeCategory{
		Model:     gorm.Model{ID: r.ID},
		Name:      r.Name,
		ImagePath: r.ImagePath,
	}
}

// FromIngredient binds entity to response
func dtoFromRecipeCategory(i *models.RecipeCategory) *RecipeCategoryResponse {
	return &RecipeCategoryResponse{
		ID:        i.ID,
		Name:      i.Name,
		ImagePath: i.ImagePath,
	}
}

func dtoToUserResponse(u models.User) *UserResponse {
	return &UserResponse{
		Name:        u.Name,
		ImagePath:   u.ImagePath,
		Description: u.Description,
	}
}

// RandString generates random string with fixted length
func RandString(n int) string {
	once.Do(func() {
		rand.Seed(time.Now().UnixNano())
	})

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
