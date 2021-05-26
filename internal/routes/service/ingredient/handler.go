package ingredient

import (
	"net/http"

	"github.com/changyoungkwon/gxample/internal/models"
	serviceError "github.com/changyoungkwon/gxample/internal/routes/service/error"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// Store abstracts repository
type Store interface {
	Add(i *models.Ingredient) error
	Get(k int) (*models.Ingredient, error)
	List() ([]models.Ingredient, error)
}

// NewIngredientRouter provies routers related to resource ingredient
func NewIngredientRouter(store Store) chi.Router {
	router := chi.NewRouter()
	router.Post("/", create(store))
	router.Get("/", list(store))
	return router
}

// create handles post
func create(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var data Request
		if err := render.Bind(req, &data); err != nil {
			render.Render(w, req, serviceError.ErrInvalidRequest(err))
			return
		}
		ingredient := data.Ingredient
		if err := store.Add(ingredient); err != nil {
			render.Render(w, req, serviceError.ErrInvalidRequest(err))
			return
		}
		render.Status(req, http.StatusCreated)
		render.Render(w, req, &Response{
			Ingredient: ingredient,
		})
	}
}

// list handles list
func list(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ingredients, err := store.List()
		if err != nil {
			render.Render(w, r, serviceError.ErrUnknown(err))
			return
		}
		responses := make([]render.Renderer, len(ingredients))
		for _, i := range ingredients {
			responses = append(responses, &Response{
				Ingredient: &i,
			})
		}
		render.Status(r, http.StatusOK)
		render.RenderList(w, r, responses)
	}
}
