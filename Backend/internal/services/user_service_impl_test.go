package services

import (
	"context"
	"testing"

	"reed_backend/internal/models"
	"reed_backend/internal/repository"
)

func TestUserService_Create_Get(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()
	svc := NewUserService(repo)
	ctx := context.Background()

	u := &models.User{ID: "s1", Username: "serviceuser", FirstName: "Service", Email: "s@example.com"}
	created, err := svc.Create(ctx, u)
	if err != nil {
		t.Fatalf("service create failed: %v", err)
	}
	if created == nil {
		t.Fatalf("created is nil")
	}

	got, err := svc.GetByID(ctx, "s1")
	if err != nil {
		t.Fatalf("get failed: %v", err)
	}
	gu, ok := got.(*models.User)
	if !ok {
		t.Fatalf("unexpected type")
	}
	if gu.Email != "s@example.com" {
		t.Fatalf("unexpected email: %s", gu.Email)
	}
}
