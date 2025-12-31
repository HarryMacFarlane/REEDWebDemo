package repository

import (
	"context"

	"reed_backend/internal/models"
)

// UserRepository defines typed storage operations for UserAccount.
type UserRepository interface {
	Create(ctx context.Context, u *models.UserAccount) error
	GetByID(ctx context.Context, id string) (*models.UserAccount, error)
	GetByEmail(ctx context.Context, email string) (*models.UserAccount, error)
	List(ctx context.Context) ([]*models.UserAccount, error)
	Update(ctx context.Context, u *models.UserAccount) error
	Delete(ctx context.Context, id string) error
}
