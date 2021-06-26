package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/changyoungkwon/gxample/internal/logging"
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
	router.Group(func(r chi.Router) {
		r.Use(multipartJSONHandler)
		r.Post("/", createRecipe(repo))
	})
	router.Get("/", listRecipes(repo))
	router.Get("/{recipeID}", getRecipe(repo))
	return router
}

// Create godoc
// @Summary Upload the recipe
// @Description Upload single recipe. The name must be unique
// @Accept  mpfd
// @Produce json
// @Param file formData file false "image of recipe"
// @Param step_1 formData file false "image of step 1"
// @Param step_2 formData file false "image of step 2"
// @Param step_3 formData file false "image of step 3"
// @Param step_4 formData file false "image of step 4"
// @Param step_5 formData file false "image of step 5"
// @Param json formData string true "json structure of recipe"
// @Success 200 {object} service.RecipeResponse
// @Failure 400,404 {object} service.ErrResponse
// @Failure 500 {object} service.ErrResponse
// @Failure default {object} service.ErrResponse
// @Router /api/recipes [post]
func createRecipe(repo RecipeRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var data recipeRequestJSON

		rawdata, err := getMultipartJSON(req.Context())
		if err != nil {
			logging.Errorf("error during get multipart-json, %v", err)
			render.Render(w, req, ErrInvalidRequest(err))
			return
		}
		err = json.Unmarshal(rawdata, &data)
		if err != nil {
			logging.Errorf("error during parsing json, %v", err)
			render.Render(w, req, ErrInvalidRequest(err))
			return
		}

		recipe, err := dtoToRecipe(&data, "0")
		if err != nil {
			logging.Errorf("error during conversion from dto, %v", err)
			render.Render(w, req, ErrInvalidRequest(err))
			return
		}

		// fill imagepath
		err = bindImagePathOnRecipe(req.Context(), recipe)
		if err != nil {
			logging.Errorf("error during context binding, %v", err)
			render.Render(w, req, ErrUnknown(err))
			return
		}

		// save all images, and save into recpie
		if err := repo.Add(recipe); err != nil {
			logging.Errorf("error during add recipe, %v", err)
			render.Render(w, req, ErrInvalidRequest(err))
			return
		}

		dto, err := dtoFromRecipe(recipe)
		if err != nil {
			logging.Errorf("error rendering, %v", err)
			render.Render(w, req, ErrUnknown(err))
			return
		}
		render.Status(req, http.StatusCreated)
		render.Render(w, req, dto)
	}
}

// List godoc
// @Summary List all uploaded recipes
// @Description List all uploaded recipes
// @Accept  json
// @Produce json
// @Success 200 {array} service.RecipeResponse
// @Failure 400,404 {object} service.ErrResponse
// @Failure 500 {object} service.ErrResponse
// @Failure default {object} service.ErrResponse
// @Router /api/recipes [get]
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

// getRecipe godoc
// @Summary Get the detail of recipe
// @Description Get the detail of recipe
// @Accept  json
// @Produce json
// @Param recipeID path int true "recipeID"
// @Success 200 {object} service.RecipeResponse
// @Failure 400,404 {object} service.ErrResponse
// @Failure 500 {object} service.ErrResponse
// @Failure default {object} service.ErrResponse
// @Router /api/recipes/{recipeID} [get]
func getRecipe(repo RecipeRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		recipeID, err := strconv.Atoi(chi.URLParam(r, "recipeID"))
		if err != nil {
			logging.Errorf("error during convert URLParam recipeID, %v", err)
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}
		recipe, err := repo.Get(recipeID)
		if err != nil {
			logging.Errorf("error during get from repo, %v", err)
			render.Render(w, r, ErrUnknown(err))
			return
		}
		response, err := dtoFromRecipe(recipe)
		if err != nil {
			logging.Errorf("error while converting to response, %v", err)
			render.Render(w, r, ErrUnknown(err))
			return
		}
		render.Render(w, r, response)
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
