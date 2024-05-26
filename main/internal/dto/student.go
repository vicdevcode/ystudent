package dto

import (
	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/main/internal/entity"
)

type Student struct {
	Fio
	Leader  bool      `json:"leader"`
	GroupID uuid.UUID `json:"group_id" binding:"required,uuid"`
	Email   string    `json:"email"    binding:"required,email"`
}

type CreateStudent Student

type UpdateStudent struct {
	ID         uuid.UUID `json:"id"         binding:"required,uuid"`
	Firstname  string    `json:"firstname"  binding:"omitempty,alphaunicode"`
	Middlename string    `json:"middlename" binding:"omitempty,alphaunicode"`
	Surname    string    `json:"surname"    binding:"omitempty,alphaunicode"`
	Email      string    `json:"email"      binding:"omitempty,email"`
	GroupID    uuid.UUID `json:"group_id"   binding:"omitempty,uuid"`
}

type UpdateStudentBody struct {
	Firstname  string    `json:"firstname"  binding:"omitempty,alphaunicode"`
	Middlename string    `json:"middlename" binding:"omitempty,alphaunicode"`
	Surname    string    `json:"surname"    binding:"omitempty,alphaunicode"`
	Email      string    `json:"email"      binding:"omitempty,email"`
	GroupID    uuid.UUID `json:"group_id"   binding:"omitempty,uuid"`
}

type StudentResponse struct {
	*entity.Student
	Password string `json:"password"`
	CUD
}

type StudentUserResponse struct {
	*entity.User
	GroupID  uuid.UUID `json:"group_id"`
	Password string    `json:"password"`
	CUD
}
