package routes

import (
	"net/http"
	"strings"

	"github.com/changyoungkwon/gxample/internal/logging"
	"github.com/changyoungkwon/gxample/internal/routes/health"
	"github.com/changyoungkwon/gxample/internal/routes/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// attachFileServer conveniently sets up a http.FileServer handler to serve
func attachFileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

// Router routes root
// @title cooker API
// @version 1.0
// @description cookerserver
// @contact.name changyoung
// @host cooker:3000
// @BasePath /v2/docs-api
func Router() http.Handler {
	router := chi.NewRouter()

	// middleware
	router.Use(middleware.RequestID) // add header x-request-id on r.context
	router.Use(middleware.RealIP)    // add real-ip on r.context
	router.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{
		Logger:  logging.Logger,
		NoColor: false,
	})) // logger
	router.Use(middleware.Recoverer) // recoverer logs the panics, returns 500

	// router: file server, docs, and service
	attachFileServer(router, "/static", http.Dir("static"))
	attachFileServer(router, "/v2", http.Dir("docs"))
	router.Get("/health", health.Handler)
	router.Mount("/api", service.NewRouter())
	router.Mount("/v2/api-docs", httpSwagger.Handler(
		httpSwagger.URL("/swagger/swagger.yaml"),
	))
	return cors.AllowAll().Handler(router)
}
