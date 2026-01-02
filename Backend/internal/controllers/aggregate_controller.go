package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"reed_backend/internal/assemblers"
	"reed_backend/internal/models"
	"reed_backend/internal/services"
	t "reed_backend/internal/transport"
	"reed_backend/internal/transport/dto"
)

// AggregateController composes multiple models into composite DTOs.
type AggregateController struct {
	userSvc services.ModelService
}

// NewAggregateController creates an AggregateController instance.
func NewAggregateController(userSvc services.ModelService) *AggregateController {
	return &AggregateController{userSvc: userSvc}
}

// RegisterRoutes mounts aggregate routes on the router.
func (c *AggregateController) RegisterRoutes(r chi.Router) {
	r.Get("/aggregate/users/{id}", c.GetUserComposite)
}

// GetUserComposite fetches a user and returns a composed DTO.
func (c *AggregateController) GetUserComposite(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	got, err := c.userSvc.GetByID(r.Context(), id)
	if err != nil {
		t.WriteError(w, http.StatusNotFound, "user not found")
		return
	}

	// Expect the service to return the domain model *models.User
	if dm, ok := got.(*models.User); ok {
		userDTO := assemblers.ToUserDTO(dm)
		out := dto.UserCompositeOutput{User: userDTO}
		t.WriteJSON(w, http.StatusOK, out)
		return
	}

	// Fallback: return whatever the service gave us under Extra
	out := dto.UserCompositeOutput{Extra: got}
	t.WriteJSON(w, http.StatusOK, out)
}
