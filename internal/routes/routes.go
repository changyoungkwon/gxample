package routes

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/changyoungkwon/gxample/internal/config"
	"github.com/changyoungkwon/gxample/internal/logging"
	"github.com/changyoungkwon/gxample/internal/routes/health"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
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

	// router
	attachFileServer(router, "/static", http.Dir("static"))
	router.Get("/health", health.Handler)
	router.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(config.Get())
	})
	return cors.AllowAll().Handler(router)
}
