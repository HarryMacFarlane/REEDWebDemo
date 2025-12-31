package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"reed_backend/internal/models"
	"reed_backend/internal/services"
)

// UserController handles HTTP requests related to UserAccount.
type UserController struct {
	svc services.ModelService
}

// NewUserController creates a controller backed by the provided service.
func NewUserController(svc services.ModelService) *UserController {
	return &UserController{svc: svc}
}

// RegisterRoutes registers user routes under /users
func (c *UserController) RegisterRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", c.List)
		r.Post("/", c.Create)
		r.Get("/{id}", c.GetByID)
		r.Put("/{id}", c.Update)
		r.Delete("/{id}", c.Delete)
	})
}

func (c *UserController) List(w http.ResponseWriter, r *http.Request) {
	out, err := c.svc.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(out)
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var u models.UserAccount
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	if u.ID == "" {
		u.ID = time.Now().UTC().Format("20060102150405.000000000")
	}
	created, err := c.svc.Create(r.Context(), &u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(created)
}

func (c *UserController) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	got, err := c.svc.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(got)
}

func (c *UserController) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var u models.UserAccount
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	// ensure ID matches
	u.ID = id
	updated, err := c.svc.Update(r.Context(), id, &u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(updated)
}

func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := c.svc.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
