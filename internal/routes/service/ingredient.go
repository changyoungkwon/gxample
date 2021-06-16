package service

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/changyoungkwon/gxample/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// IngredientRepo abstract ingredient-db
type IngredientRepo interface {
	Add(i *models.Ingredient) error
	Get(k int) (*models.Ingredient, error)
	List() ([]models.Ingredient, error)
}

// NewIngredientRouter provies routers related to resource ingredient
func NewIngredientRouter(repo IngredientRepo) chi.Router {
	router := chi.NewRouter()
	router.Group(func(r chi.Router) {
		r.Use(multipartJSONHandler)
		r.Post("/", createIngredient(repo))
	})
	router.Get("/", listIngredients(repo))
	return router
}

// Create godoc
// @Summary Upload an ingredient
// @Description Upload single ingredient. The name must be unique
// @Accept  mpfd
// @Produce json
// @Param file formData file false "image of ingredient"
// @Param json formData string true "json structure of ingredient"
// @Success 200 {object} service.IngredientResponse
// @Failure 400,404 {object} service.ErrResponse
// @Failure 500 {object} service.ErrResponse
// @Failure default {object} service.ErrResponse
// @Router /api/ingredients [post]
func createIngredient(repo IngredientRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var data ingredientRequestJSON
		// unmarshal form-data to ingredient
		rawdata, err := getMultipartJSON(req.Context())
		if err != nil {
			render.Render(w, req, ErrInvalidRequest(err))
			return
		}
		err = json.Unmarshal(rawdata, &data)
		if err != nil {
			render.Render(w, req, ErrInvalidRequest(err))
			return
		}
		ingredient := dtoToIngredient(&data)
		err = bindImagePathOnIngredient(req.Context(), ingredient)
		if err := repo.Add(ingredient); err != nil {
			render.Render(w, req, ErrInvalidRequest(err))
			return
		}
		render.Status(req, http.StatusCreated)
		render.Render(w, req, dtoFromIngredient(ingredient))
	}
}

// List godoc
// @Summary List all uploaded ingredients
// @Description List all uploaded ingredients
// @Accept  json
// @Produce json
// @Success 200 {array} service.IngredientResponse
// @Failure 400,404 {object} service.ErrResponse
// @Failure 500 {object} service.ErrResponse
// @Failure default {object} service.ErrResponse
// @Router /api/ingredients [get]
func listIngredients(repo IngredientRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ingredients, err := repo.List()
		if err != nil {
			render.Render(w, r, ErrUnknown(err))
			return
		}
		responses := make([]render.Renderer, 0, len(ingredients))
		for _, i := range ingredients {
			responses = append(responses, dtoFromIngredient(&i))
		}
		render.Status(r, http.StatusOK)
		render.RenderList(w, r, responses)
	}
}

func bindImagePathOnIngredient(c context.Context, i *models.Ingredient) error {
	imagePaths, ok := c.Value(imageHandleKey).(map[string]string)
	if !ok {
		return errors.New("no imagehandle middleware gives wrong")
	}
	if filepath, ok := imagePaths["file"]; !ok {
		i.ImagePath = ""
	} else {
		i.ImagePath = filepath
	}
	return nil
}

// Bind binds additional parameters on Request after decode
func (i *IngredientRequest) Bind(r *http.Request) error {
	return nil
}

// Render renders additional paramters before encode and response
func (i *IngredientResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
