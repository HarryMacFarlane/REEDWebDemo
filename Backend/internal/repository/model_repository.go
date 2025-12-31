package repository

import "context"

// ModelRepository defines storage operations for the Model resource.
// Implement this interface against your chosen persistence (SQL, NoSQL, in-memory).
type ModelRepository interface {
	Create(ctx context.Context, m interface{}) error
	GetByID(ctx context.Context, id string) (interface{}, error)
	List(ctx context.Context) ([]interface{}, error)
	Update(ctx context.Context, id string, m interface{}) error
	Delete(ctx context.Context, id string) error
}

// TODO: add concrete implementations (e.g., postgresRepository, memoryRepository)
