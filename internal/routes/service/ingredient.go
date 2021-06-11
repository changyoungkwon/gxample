// ingredient.go contains ingredienthandler
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
	router.Use(imageHandleMiddleware)
	router.Post("/", createIngredient(repo))
	router.Get("/", listIngredients(repo))
	return router
}

// create handles post
func createIngredient(repo IngredientRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var data ingredientRequestJSON

		// unmarshal form-data to ingredient
		err := json.Unmarshal([]byte(req.FormValue("json")), &data)
		if err != nil {
			render.Render(w, req, ErrInvalidRequest(err))
			return
		}
		ingredient := dtoToIngredient(&data)
		err = bindImagePathOnIngredient(ingredient, req.Context())
		if err := repo.Add(ingredient); err != nil {
			render.Render(w, req, ErrInvalidRequest(err))
			return
		}
		render.Status(req, http.StatusCreated)
		render.Render(w, req, dtoFromIngredient(ingredient))
	}
}

// list handles list
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

func bindImagePathOnIngredient(i *models.Ingredient, c context.Context) error {
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
