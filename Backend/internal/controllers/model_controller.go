package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// ModelController handles HTTP requests for the Model resource.
type ModelController struct {
	// inject services/repositories here when implemented
}

// NewModelController creates a new controller instance.
func NewModelController() *ModelController {
	return &ModelController{}
}

// RegisterRoutes registers the model routes under the provided router.
func (c *ModelController) RegisterRoutes(r chi.Router) {
	r.Route("/models", func(r chi.Router) {
		r.Get("/", c.List)
		r.Post("/", c.Create)
		r.Get("/{id}", c.GetByID)
		r.Put("/{id}", c.Update)
		r.Delete("/{id}", c.Delete)
	})
}

// Handlers below are placeholders. Implement business logic using services.
func (c *ModelController) List(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
	w.WriteHeader(http.StatusNotImplemented)
}

func (c *ModelController) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
	w.WriteHeader(http.StatusNotImplemented)
}

func (c *ModelController) GetByID(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
	w.WriteHeader(http.StatusNotImplemented)
}

func (c *ModelController) Update(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
	w.WriteHeader(http.StatusNotImplemented)
}

func (c *ModelController) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
	w.WriteHeader(http.StatusNotImplemented)
}
