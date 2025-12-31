package services

import (
	"context"
	"errors"

	"reed_backend/internal/models"
	"reed_backend/internal/repository"
)

// UserServiceImpl is a concrete implementation of a user service.
type UserServiceImpl struct {
	repo repository.UserRepository
}

// NewUserService creates a new UserServiceImpl.
func NewUserService(repo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) Create(ctx context.Context, payload interface{}) (interface{}, error) {
	u, ok := payload.(*models.UserAccount)
	if !ok || u == nil {
		return nil, errors.New("invalid payload")
	}
	// caller should set ID (or generate here). For now require ID present.
	if u.ID == "" {
		return nil, errors.New("missing id")
	}
	if err := s.repo.Create(ctx, u); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *UserServiceImpl) GetByID(ctx context.Context, id string) (interface{}, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserServiceImpl) List(ctx context.Context) ([]interface{}, error) {
	us, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]interface{}, 0, len(us))
	for _, u := range us {
		out = append(out, u)
	}
	return out, nil
}

func (s *UserServiceImpl) Update(ctx context.Context, id string, payload interface{}) (interface{}, error) {
	u, ok := payload.(*models.UserAccount)
	if !ok || u == nil {
		return nil, errors.New("invalid payload")
	}
	if u.ID == "" {
		u.ID = id
	}
	if err := s.repo.Update(ctx, u); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *UserServiceImpl) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
