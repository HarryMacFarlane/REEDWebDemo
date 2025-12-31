package routes

import (
	"github.com/go-chi/chi/v5"

	"reed_backend/internal/controllers"
	"reed_backend/internal/repository"
	"reed_backend/internal/services"
)

// RegisterRoutes wires up controller routes onto the provided router.
// Controllers are responsible for their handler implementations.
func RegisterRoutes(r chi.Router) {
	// Create in-memory repo + service and register user routes for now.
	repo := repository.NewInMemoryUserRepository()
	svc := services.NewUserService(repo)
	userCtrl := controllers.NewUserController(svc)
	userCtrl.RegisterRoutes(r)

	// TODO: register other controllers (models, admin, etc.)
}
