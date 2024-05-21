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

type StudentResponse struct {
	*entity.Student
	Password string `json:"password"`
	CUD
}

type StudentUserResponse struct {
	*entity.User
	Password string `json:"password"`
	CUD
}
