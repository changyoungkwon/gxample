package service

import (
	"io"
	"time"

	"github.com/changyoungkwon/gxample/internal/models"
)

// IngredientRequest wraps request
type IngredientRequest struct {
	File io.Reader `form:"file" json:"file"`
	JSON []byte    `form:"json" json:"json"`
}

// ingredientReqeustJSON is represents IngredientRequest.JSON
type ingredientRequestJSON struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// IngredientResponse wraps response
type IngredientResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	ImagePath string `json:"image_path"`
}

// RecipeCategoryRequest body to post/update new recipe-category
type RecipeCategoryRequest struct {
	File io.Reader `form:"file" json:"file"`
	JSON []byte    `form:"json" json:"json"`
}

type recipeCategoryRequestJSON struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	ImagePath string `json:"image_path"`
}

// RecipeCategoryResponse body to response
type RecipeCategoryResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	ImagePath string `json:"image_path"`
}

// IngredientQuantityRequest body to post/update new recipe-category
type IngredientQuantityRequest struct {
	IngredientID uint        `json:"ingredient_id"`
	Quantity     interface{} `json:"quantity"`
}

// IngredientQuantityResponse body to response
type IngredientQuantityResponse struct {
	Ingredient IngredientResponse `json:"ingredient"`
	Quantity   interface{}        `json:"quantity" swaggertype:"object"`
}

// RecipeRequest wraps request
type RecipeRequest struct {
	File io.Reader `form:"file" json:"file"`
	JSON []byte    `form:"json" json:"json"`
}

type recipeRequestJSON struct {
	ID                   uint                         `json:"id"`
	Title                string                       `json:"title"`
	Ease                 string                       `json:"ease"`
	PreparationTime      int                          `json:"preparation_time"`
	RecipeCategoryID     int                          `json:"recipe_category_id"`
	IngredientQuantities []*IngredientQuantityRequest `json:"ingredients"`
	Steps                models.RecipeSteps           `json:"steps"`
	Tags                 []string                     `json:"tags"`
}

// RecipeResponse wraps response
type RecipeResponse struct {
	ID                   uint                          `json:"id"`
	UpdatedAt            time.Time                     `json:"updated_at"`
	CreatedAt            time.Time                     `json:"created_at"`
	Title                string                        `json:"title"`
	ImagePath            string                        `json:"image_path"`
	Ease                 string                        `json:"ease"`
	PreparationTime      int                           `json:"preparation_time"`
	RecipeCategory       RecipeCategoryResponse        `json:"recipe_category"`
	IngredientQuantities []*IngredientQuantityResponse `json:"ingredient_quantities"`
	Writer               UserResponse                  `json:"writer"`
	Steps                models.RecipeSteps            `json:"steps"`
	IsClipped            bool                          `json:"is_clipped"`
	Tags                 []string                      `json:"tags"`
}

// RecipeThumbResponse wraps lsit responses
type RecipeThumbResponse struct {
	ID              uint         `json:"id"`
	Title           string       `json:"title"`
	PreparationTime int          `json:"preparation_time"`
	Ease            string       `json:"ease"`
	ImagePath       string       `json:"image_path"`
	Writer          UserResponse `json:"writer"`
	IsClipped       bool         `json:"is_clipped"`
}

// UserResponse wraps response
type UserResponse struct {
	Name        string `json:"name"`
	ImagePath   string `json:"image_path"`
	Description string `json:"description"`
}
