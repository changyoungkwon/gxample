package recipe

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/changyoungkwon/gxample/internal/models"
)

// Request wraps request
type Request struct {
	File io.Reader `form:"file" json:"file"`
	JSON []byte    `form:"json" json:"json"`
}

// Response wraps response
type Response struct {
	ID                   uint                         `json:"id"`
	UpdatedAt            time.Time                    `json:"updated_at"`
	CreatedAt            time.Time                    `json:"created_at"`
	Title                string                       `json:"title"`
	ImagePath            string                       `json:"image_path"`
	Ease                 string                       `json:"ease"`
	PreparationTime      int                          `json:"preparation_time"`
	RecipeCategory       models.RecipeCategory        `json:"recipe_category"`
	IngredientQuantities []*models.IngredientQuantity `json:"ingredient_quantities"`
	WriterID             string                       `json:"writer_id"`
	Steps                models.RecipeSteps           `json:"steps"`
	IsClipped            bool                         `json:"is_clipped"`
	Tags                 []string                     `json:"tags"`
}

type requestJSON struct {
	Title                string                       `json:"title"`
	Ease                 string                       `json:"ease"`
	PreparationTime      int                          `json:"preparation_time"`
	RecipeCategoryID     int                          `json:"recipe_category_id"`
	WriterID             string                       `json:"writer_id"`
	IngredientQuantities []*models.IngredientQuantity `json:"ingredient_quantities"`
	Steps                []models.RecipeStep          `json:"steps"`
	Tags                 []string                     `json:"tags"`
}

// NewRecipe generates recipes based on request
func (i *Request) NewRecipe() (*models.Recipe, error) {
	tags := []*models.Tag{}
	var body requestJSON
	err := json.Unmarshal(i.JSON, &body)
	if err != nil {
		return nil, err
	}
	for _, t := range body.Tags {
		tags = append(tags, &models.Tag{Name: t})
	}

	return &models.Recipe{
		Title:                body.Title,
		Ease:                 body.Ease,
		PreparationTime:      body.PreparationTime,
		RecipeCategoryID:     body.RecipeCategoryID,
		WriterID:             body.WriterID,
		IngredientQuantities: body.IngredientQuantities,
		Steps:                body.Steps,
		Tags:                 tags,
	}, nil
}

// MapFrom generates based on recipes
func MapFrom(r *models.Recipe) *Response {
	tags := make([]string, 0, len(r.Tags))
	for _, t := range r.Tags {
		tags = append(tags, t.Name)
	}
	return &Response{
		ID:                   r.ID,
		CreatedAt:            r.CreatedAt,
		UpdatedAt:            r.UpdatedAt,
		Title:                r.Title,
		ImagePath:            r.ImagePath,
		Ease:                 r.Ease,
		PreparationTime:      r.PreparationTime,
		RecipeCategory:       r.RecipeCategory,
		IngredientQuantities: r.IngredientQuantities,
		WriterID:             r.WriterID,
		Steps:                r.Steps,
		IsClipped:            false,
		Tags:                 tags,
	}
}

// Bind binds additional parameters on IngredientRequest after decode
func (i *Request) Bind(r *http.Request) error {
	return nil
}

// Render renders additional paramters before encode and response
func (i *Response) Render(w http.ResponseWriter, r *http.Request) error {
	i.IsClipped = false
	return nil
}
