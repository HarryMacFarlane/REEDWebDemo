package repository

import (
	"context"
	"errors"
	"sync"
	"time"

	"reed_backend/internal/models"
)

// ErrNotFound indicates an item was not found in the repository.
var ErrNotFound = errors.New("not found")

// InMemoryUserRepository is a simple in-memory implementation of UserRepository
// intended for development and tests.
type InMemoryUserRepository struct {
	mu    sync.RWMutex
	store map[string]*models.User
}

// NewInMemoryUserRepository creates a new in-memory repository.
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		store: make(map[string]*models.User),
	}
}

func (r *InMemoryUserRepository) Create(ctx context.Context, u *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if u == nil {
		return errors.New("nil user")
	}
	now := time.Now().UTC()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = now
	}
	u.UpdatedAt = now
	r.store[u.ID] = u
	return nil
}

func (r *InMemoryUserRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	u, ok := r.store[id]
	if !ok {
		return nil, ErrNotFound
	}
	return u, nil
}

func (r *InMemoryUserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, u := range r.store {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, ErrNotFound
}

func (r *InMemoryUserRepository) List(ctx context.Context) ([]*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]*models.User, 0, len(r.store))
	for _, u := range r.store {
		out = append(out, u)
	}
	return out, nil
}

func (r *InMemoryUserRepository) Update(ctx context.Context, u *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if u == nil {
		return errors.New("nil user")
	}
	existing, ok := r.store[u.ID]
	if !ok {
		return ErrNotFound
	}
	u.CreatedAt = existing.CreatedAt
	u.UpdatedAt = time.Now().UTC()
	r.store[u.ID] = u
	return nil
}

func (r *InMemoryUserRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.store[id]; !ok {
		return ErrNotFound
	}
	delete(r.store, id)
	return nil
}
