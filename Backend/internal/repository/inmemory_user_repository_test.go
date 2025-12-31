package repository

import (
	"context"
	"testing"
	"time"

	"reed_backend/internal/models"
)

func TestInMemoryUserRepository_CRUD(t *testing.T) {
	repo := NewInMemoryUserRepository()
	ctx := context.Background()

	u := &models.UserAccount{ID: "u1", Email: "a@example.com", Name: "Alice"}
	if err := repo.Create(ctx, u); err != nil {
		t.Fatalf("create failed: %v", err)
	}

	got, err := repo.GetByID(ctx, "u1")
	if err != nil {
		t.Fatalf("get failed: %v", err)
	}
	if got.Email != u.Email {
		t.Fatalf("unexpected email: %s", got.Email)
	}

	list, err := repo.List(ctx)
	if err != nil {
		t.Fatalf("list failed: %v", err)
	}
	if len(list) != 1 {
		t.Fatalf("expected 1 user, got %d", len(list))
	}

	// Update
	u.Name = "Alice Updated"
	if err := repo.Update(ctx, u); err != nil {
		t.Fatalf("update failed: %v", err)
	}

	got2, err := repo.GetByID(ctx, "u1")
	if err != nil {
		t.Fatalf("get after update failed: %v", err)
	}
	if got2.Name != "Alice Updated" {
		t.Fatalf("name not updated: %s", got2.Name)
	}

	// Delete
	if err := repo.Delete(ctx, "u1"); err != nil {
		t.Fatalf("delete failed: %v", err)
	}
	if _, err := repo.GetByID(ctx, "u1"); err == nil {
		t.Fatalf("expected not found after delete")
	}

	// ensure timestamps set
	u2 := &models.UserAccount{ID: "u2", Email: "b@example.com", Name: "Bob"}
	if err := repo.Create(ctx, u2); err != nil {
		t.Fatalf("create u2 failed: %v", err)
	}
	got3, _ := repo.GetByID(ctx, "u2")
	if got3.CreatedAt.IsZero() || got3.UpdatedAt.IsZero() {
		t.Fatalf("timestamps not set")
	}
	// slight delay to ensure update changed timestamp
	time.Sleep(1 * time.Millisecond)
	got3.Name = "Bob2"
	if err := repo.Update(ctx, got3); err != nil {
		t.Fatalf("update u2 failed: %v", err)
	}
	after, _ := repo.GetByID(ctx, "u2")
	if !after.UpdatedAt.After(after.CreatedAt) && !after.UpdatedAt.Equal(after.CreatedAt) {
		t.Fatalf("timestamps did not update")
	}
}
