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

// RecipeCategoryRepo abstract recipeCategory-db
type RecipeCategoryRepo interface {
	Add(i *models.RecipeCategory) error
	Get(k int) (*models.RecipeCategory, error)
	List() ([]models.RecipeCategory, error)
}

// NewRecipeCategoryRouter provies routers related to resource recipeCategory
func NewRecipeCategoryRouter(repo RecipeCategoryRepo) chi.Router {
	router := chi.NewRouter()
	router.Use(imageHandleMiddleware)
	router.Post("/", createRecipeCategory(repo))
	router.Get("/", listRecipeCategories(repo))
	return router
}

// create handles post
func createRecipeCategory(repo RecipeCategoryRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var data recipeCategoryRequestJSON

		// unmarshal form-data to recipeCategory
		err := json.Unmarshal([]byte(req.FormValue("json")), &data)
		if err != nil {
			render.Render(w, req, ErrInvalidRequest(err))
			return
		}
		recipeCategory := dtoToRecipeCategory(&data)
		err = bindImagePathOnRecipeCategory(req.Context(), recipeCategory)
		if err := repo.Add(recipeCategory); err != nil {
			render.Render(w, req, ErrInvalidRequest(err))
			return
		}
		render.Status(req, http.StatusCreated)
		render.Render(w, req, dtoFromRecipeCategory(recipeCategory))
	}
}

// list handles list
func listRecipeCategories(repo RecipeCategoryRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		recipeCategories, err := repo.List()
		if err != nil {
			render.Render(w, r, ErrUnknown(err))
			return
		}
		responses := make([]render.Renderer, 0, len(recipeCategories))
		for _, i := range recipeCategories {
			responses = append(responses, dtoFromRecipeCategory(&i))
		}
		render.Status(r, http.StatusOK)
		render.RenderList(w, r, responses)
	}
}

func bindImagePathOnRecipeCategory(c context.Context, i *models.RecipeCategory) error {
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
func (i *RecipeCategoryRequest) Bind(r *http.Request) error {
	return nil
}

// Render renders additional paramters before encode and response
func (i *RecipeCategoryResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
