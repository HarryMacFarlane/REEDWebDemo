package assemblers

import (
	"reed_backend/internal/models"
	"reed_backend/internal/transport/dto"
)

// ToUserDTO converts a domain User to a transport DTO, omitting sensitive fields.
func ToUserDTO(u *models.User) dto.UserOutput {
	if u == nil {
		return dto.UserOutput{}
	}
	return dto.UserOutput{
		ID:        u.ID,
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	}
}
