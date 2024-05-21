package dto

import "github.com/vicdevcode/ystudent/main/internal/entity"

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

type UserResponse struct {
	*entity.User
	CUD
}

type ModeratorResponse struct {
	*entity.User
	Password string `json:"password"`
	CUD
}
