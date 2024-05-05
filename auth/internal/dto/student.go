package dto

import (
	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/auth/internal/entity"
)

type Student struct {
	Leader  bool      `json:"leader"`
	GroupID uuid.UUID `json:"group_id" binding:"required,uuid"`
	UserID  uuid.UUID `json:"user_id"`
}

type CreateStudent Student

type StudentResponse struct {
	*entity.Student
	CUD
}
