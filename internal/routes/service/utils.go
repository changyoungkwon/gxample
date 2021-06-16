package service

import (
	"encoding/json"

	"github.com/changyoungkwon/gxample/internal/models"
	"gorm.io/gorm"
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
	// to response
	// tags
	tags := make([]string, 0, len(r.Tags))
	for _, t := range r.Tags {
		tags = append(tags, t.Name)
	}
	// ingerdient quantities
	quantities, err := dtoFromIQuantities(r.IngredientQuantities)
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
		Writer:               UserResponse{},
		Steps:                r.Steps,
		IsClipped:            false,
		Tags:                 tags,
	}, nil
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
func dtoFromIQuantities(qs []*models.IngredientQuantity) ([]*IngredientQuantityResponse, error) {
	var quantities []*IngredientQuantityResponse
	for _, q := range qs {
		var res IngredientQuantityResponse
		err := json.Unmarshal(q.Quantity, &res)
		if err != nil {
			return nil, err
		}
		quantities = append(quantities, &IngredientQuantityResponse{
			Ingredient: *dtoFromIngredient(&q.Ingredient),
			Quantity:   res,
		})
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
