package service

import (
	"net/http"

	"github.com/changyoungkwon/gxample/internal/database"
	"github.com/changyoungkwon/gxample/internal/routes/service/ingredient"
	"github.com/changyoungkwon/gxample/internal/routes/service/recipe"
	"github.com/go-chi/chi/v5"
)

// NewRouter gives service api
func NewRouter() http.Handler {
	router := chi.NewRouter()
	db := database.DBConn()
	router.Mount("/ingredients",
		ingredient.NewIngredientRouter(database.NewIngredientStore(db)))
	router.Mount("/recipes",
		recipe.NewRecipeRouter(database.NewRecipeStore(db)))
	return router
}
