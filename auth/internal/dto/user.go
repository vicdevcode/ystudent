package dto

import (
	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/auth/internal/entity"
)

type Fio struct {
	Firstname  string `json:"firstname"  binding:"required,alphaunicode"`
	Middlename string `json:"middlename" binding:"omitempty,alphaunicode"`
	Surname    string `json:"surname"    binding:"required,alphaunicode"`
}

type CreateUser struct {
	Fio
	ID       uuid.UUID       `json:"id"`
	Email    string          `json:"email"`
	Password string          `json:"password"`
	Role     entity.UserRole `json:"user_role"`
}

type UserResponse struct {
	*entity.User
	CUD
	Password     interface{} `json:"password,omitempty"`
	RefreshToken interface{} `json:"refresh_token,omitempty"`
}
