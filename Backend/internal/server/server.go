package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"reed_backend/internal/middleware"
	"reed_backend/internal/routes"
)

// NewRouter returns a chi router pre-wired with middleware and routes.
func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	// use default middlewares (RequestID or custom)
	r.Use(middleware.DefaultMiddlewares())

	// health endpoint
	r.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("ok"))
	})

	// register application routes
	routes.RegisterRoutes(r)

	return r
}
