package assemblers

import (
	"testing"
	"time"

	"reed_backend/internal/models"
)

func TestToUserDTO(t *testing.T) {
	now := time.Now().UTC()
	u := &models.User{
		ID:        "u1",
		Username:  "bob",
		FirstName: "Bob",
		LastName:  "Smith",
		Email:     "bob@example.com",
		CreatedAt: now,
	}
	dto := ToUserDTO(u)
	if dto.ID != "u1" {
		t.Fatalf("unexpected id: %s", dto.ID)
	}
	if dto.Username != "bob" {
		t.Fatalf("unexpected username: %s", dto.Username)
	}
	if !dto.CreatedAt.Equal(now) {
		t.Fatalf("createdAt mismatch")
	}
}
