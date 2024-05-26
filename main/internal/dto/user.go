package dto

import (
	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/main/internal/entity"
)

type User struct {
	Fio
	Student Student `json:"student,omitempty"`
}

type Fio struct {
	Firstname  string `json:"firstname"  binding:"required,alphaunicode"`
	Middlename string `json:"middlename" binding:"omitempty,alphaunicode"`
	Surname    string `json:"surname"    binding:"required,alphaunicode"`
}

type CreateUser struct {
	Fio
	Email    string          `json:"email"     binding:"required,email"`
	RoleType entity.UserRole `json:"role_type"`
}

type UpdateUser struct {
	ID         uuid.UUID `json:"id"         binding:"required,uuid"`
	Firstname  string    `json:"firstname"  binding:"omitempty,alphaunicode"`
	Middlename string    `json:"middlename" binding:"omitempty,alphaunicode"`
	Surname    string    `json:"surname"    binding:"omitempty,alphaunicode"`
	Email      string    `json:"email"      binding:"omitempty,email"`
}

type UserResponse struct {
	*entity.User
	CUD
}

type ModeratorResponse struct {
	*entity.User
	Password string `json:"password"`
	CUD
}
