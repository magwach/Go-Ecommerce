package dto

import (
	"go-ecommerce-app/internal/schema"
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Verified  bool      `json:"verified"`
	UserType  string    `json:"user_type"`
	CreatedAt time.Time `json:"created_at"`
}

func ToUserResponse(u schema.User) UserResponse {
	return UserResponse{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Phone:     u.Phone,
		Verified:  u.Verified,
		UserType:  u.UserType,
		CreatedAt: u.CreatedAt,
	}
}
