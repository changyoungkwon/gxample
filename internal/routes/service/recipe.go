package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/changyoungkwon/gxample/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// RecipeRepo abstracts repository
type RecipeRepo interface {
	Add(*models.Recipe) error
	Get(k int) (*models.Recipe, error)
	List() ([]models.Recipe, error)
}

// NewRecipeRouter exports router for ingredient resource
func NewRecipeRouter(repo RecipeRepo) chi.Router {
	router := chi.NewRouter()
	router.Use(imageHandleMiddleware)
	router.Post("/", createRecipe(repo))
	router.Get("/", listRecipes(repo))
	return router
}

// create handle post
func createRecipe(repo RecipeRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var data recipeRequestJSON

		// unmarshal form-data to recipe
		err := json.Unmarshal([]byte(req.FormValue("json")), &data)
		if err != nil {
			render.Render(w, req, ErrInvalidRequest(err))
			return
		}
		recipe, err := dtoToRecipe(&data, "")
		if err != nil {
			render.Render(w, req, ErrInvalidRequest(err))
			return
		}

		// fill imagepath
		err = bindImagePathOnRecipe(req.Context(), recipe)
		if err != nil {
			render.Render(w, req, ErrUnknown(err))
			return
		}

		// save all images, and save into recpie
		if err := repo.Add(recipe); err != nil {
			render.Render(w, req, ErrInvalidRequest(err))
			return
		}

		dto, err := dtoFromRecipe(recipe)
		if err != nil {
			render.Render(w, req, ErrUnknown(err))
			return
		}
		render.Status(req, http.StatusCreated)
		render.Render(w, req, dto)
	}
}

func listRecipes(repo RecipeRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		recipes, err := repo.List()
		if err != nil {
			render.Render(w, r, ErrUnknown(err))
			return
		}
		responses := make([]render.Renderer, 0, len(recipes))
		for _, rcp := range recipes {
			dto, err := dtoFromRecipe(&rcp)
			if err != nil {
				render.Render(w, r, ErrUnknown(err))
				return
			}
			responses = append(responses, dto)
		}
		render.Status(r, http.StatusOK)
		render.RenderList(w, r, responses)
	}
}

// bindImagePathOnRecipe bind imagepaths in context to recipe
func bindImagePathOnRecipe(c context.Context, r *models.Recipe) error {
	imagePaths, ok := c.Value(imageHandleKey).(map[string]string)
	if !ok {
		return errors.New("no imagehandle middleware gives wrong")
	}
	if mainpath, ok := imagePaths["file"]; !ok {
		r.ImagePath = ""
	} else {
		r.ImagePath = mainpath
	}
	for i, step := range r.Steps {
		key := fmt.Sprintf("step_%d", step.Index)
		if steppath, ok := imagePaths[key]; !ok {
			r.Steps[i].ImagePath = ""
		} else {
			r.Steps[i].ImagePath = steppath
		}
	}
	return nil
}

// Bind binds additional parameters on IngredientRequest after decode
func (i *RecipeRequest) Bind(r *http.Request) error {
	return nil
}

// Render renders additional paramters before encode and response
func (i *RecipeResponse) Render(w http.ResponseWriter, r *http.Request) error {
	i.IsClipped = false
	return nil
}
