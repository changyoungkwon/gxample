package service

import (
	"net/http"

	"github.com/changyoungkwon/gxample/internal/database"
	"github.com/changyoungkwon/gxample/internal/routes/service/ingredient"
	"github.com/go-chi/chi/v5"
)

// Router gives service api
func Router() http.Handler {
	router := chi.NewRouter()
	db := database.DBConn()
	router.Mount("/ingredients",
		ingredient.NewIngredientRouter(database.NewIngredientStore(db)))
	return router
}
