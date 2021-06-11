package service

import (
	"net/http"

	"github.com/changyoungkwon/gxample/internal/database"
	"github.com/go-chi/chi/v5"
)

// NewRouter gives service api
func NewRouter() http.Handler {
	router := chi.NewRouter()
	db := database.DBConn()
	router.Mount("/ingredients",
		NewIngredientRouter(database.NewIngredientStore(db)))
	router.Mount("/recipes",
		NewRecipeRouter(database.NewRecipeStore(db)))
	router.Mount("/recipe-categories",
		NewRecipeCategoryRouter(database.NewRecipeCategoryStore(db)))
	return router
}
