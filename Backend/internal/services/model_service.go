package services

import "context"

// ModelService defines business logic operations for Models.
// Wire this to repository implementations when available.
type ModelService interface {
	Create(ctx context.Context, payload interface{}) (interface{}, error)
	GetByID(ctx context.Context, id string) (interface{}, error)
	List(ctx context.Context) ([]interface{}, error)
	Update(ctx context.Context, id string, payload interface{}) (interface{}, error)
	Delete(ctx context.Context, id string) error
}

// TODO: add a concrete implementation that depends on repository.ModelRepository
