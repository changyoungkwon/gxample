package recipe

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/changyoungkwon/gxample/internal/logging"
	"github.com/changyoungkwon/gxample/internal/models"
	common "github.com/changyoungkwon/gxample/internal/routes/service/common"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

// Store abstracts repository
type Store interface {
	Add(*models.Recipe) error
	Get(k int) (*models.Recipe, error)
	List() ([]models.Recipe, error)
}

// NewRecipeRouter exports router for ingredient resource
func NewRecipeRouter(store Store) chi.Router {
	router := chi.NewRouter()
	router.Post("/", create(store))
	router.Get("/", list(store))
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
	imageDir := path.Join(common.StaticRootPath, dirname.String())
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
func create(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var data Request

		// handles form
		err := req.ParseMultipartForm(32 << 20) // 32MB
		if err != nil {
			render.Render(w, req, common.ErrInvalidRequest(err))
		}

		// handles json
		data.JSON = []byte(req.FormValue("json"))
		recipe, err := data.NewRecipe()
		if err != nil {
			render.Render(w, req, common.ErrInvalidRequest(err))
			return
		}

		// parse images
		if err := parseImages(recipe, req); err != nil {
			render.Render(w, req, common.ErrUnknown(err))
			return
		}

		// save all images, and save into recpie
		if err := store.Add(recipe); err != nil {
			render.Render(w, req, common.ErrInvalidRequest(err))
			return
		}
		render.Status(req, http.StatusCreated)
		render.Render(w, req, MapFrom(recipe))
	}
}

func list(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		recipes, err := store.List()
		if err != nil {
			render.Render(w, r, common.ErrUnknown(err))
			return
		}
		responses := make([]render.Renderer, 0, len(recipes))
		for i, rcp := range recipes {
			responses = append(responses, MapFrom(&rcp))
			logging.Logger.Infof("value: %v", responses[i])
		}
		render.Status(r, http.StatusOK)
		render.RenderList(w, r, responses)
	}
}
