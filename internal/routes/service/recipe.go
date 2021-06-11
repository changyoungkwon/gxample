package service

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/changyoungkwon/gxample/internal/logging"
	"github.com/changyoungkwon/gxample/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
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
	router.Post("/", createRecipe(repo))
	router.Get("/", listRecipes(repo))
	return router
}

func saveMultipartFile(r *http.Request, key string) (string, error) {
	file, header, err := r.FormFile(key)
	// if file is missig, return error
	if err != nil {
		return "", err
	}
	defer file.Close()

	// set image path, then save
	dirname, _ := uuid.NewUUID()
	imageDir := path.Join(StaticRootPath, dirname.String())
	if err := os.MkdirAll(imageDir, os.FileMode(0755)); err != nil {
		return "", err
	}
	imagePath := path.Join(imageDir, header.Filename)
	out, err := os.Create(imagePath)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, file) // save file
	if err != nil {
		return "", err
	}
	return imagePath, nil
}

func parseImages(rcp *models.Recipe, req *http.Request) error {
	imagePath, err := saveMultipartFile(req, "recipe")
	if err != nil {
		if err == http.ErrMissingFile {
			logging.Logger.Warnf("missing file %v", err)
		} else {
			logging.Logger.Errorf("error while save multipart-images, %v", err)
			return err
		}
	}
	rcp.ImagePath = imagePath
	for i, step := range rcp.Steps {
		key := fmt.Sprintf("step_%d", int(step.Index))
		logging.Logger.Info("check %s", key)
		imagePath, err := saveMultipartFile(req, key)
		if err != nil {
			if err == http.ErrMissingFile {
				logging.Logger.Warnf("missing file %v", err)
				continue
			}
			logging.Logger.Errorf("error while save multipart-images, %v", err)
			return err
		}
		rcp.Steps[i].ImagePath = imagePath
	}
	return nil
}

// create handle post
func createRecipe(repo RecipeRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var data RecipeRequest

		// handles form
		err := req.ParseMultipartForm(32 << 20) // 32MB
		if err != nil {
			render.Render(w, req, ErrInvalidRequest(err))
		}

		// handles json
		data.JSON = []byte(req.FormValue("json"))
		recipe, err := data.ToRecipe()
		if err != nil {
			render.Render(w, req, ErrInvalidRequest(err))
			return
		}

		// parse images
		if err := parseImages(recipe, req); err != nil {
			render.Render(w, req, ErrUnknown(err))
			return
		}

		// save all images, and save into recpie
		if err := repo.Add(recipe); err != nil {
			render.Render(w, req, ErrInvalidRequest(err))
			return
		}
		dto, err := dtoFromRecipe(&recipe)
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

// Bind binds additional parameters on IngredientRequest after decode
func (i *RecipeRequest) Bind(r *http.Request) error {
	return nil
}

// Render renders additional paramters before encode and response
func (i *RecipeResponse) Render(w http.ResponseWriter, r *http.Request) error {
	i.IsClipped = false
	return nil
}
