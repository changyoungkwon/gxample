package recipe

import (
	"net/http"

	"github.com/changyoungkwon/gxample/internal/models"
	serviceError "github.com/changyoungkwon/gxample/internal/routes/service/error"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
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

// create handle post
func create(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var data Request
		if err := render.Bind(req, &data); err != nil {
			render.Render(w, req, serviceError.ErrInvalidRequest(err))
			return
		}
		recipe := data.Recipe
		if err := store.Add(recipe); err != nil {
			render.Render(w, req, serviceError.ErrInvalidRequest(err))
			return
		}
		render.Status(req, http.StatusCreated)
		render.Render(w, req, &Response{
			Recipe: recipe,
		})
	}
}

func list(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		recipes, err := store.List()
		if err != nil {
			render.Render(w, r, serviceError.ErrUnknown(err))
			return
		}
		responses := make([]render.Renderer, len(recipes))
		for _, rcp := range recipes {
			responses = append(responses, &Response{
				Recipe: &rcp,
			})
		}
		render.Status(r, http.StatusOK)
		render.RenderList(w, r, responses)
	}
}
