package dto

import (
	"github.com/vicdevcode/ystudent/auth/internal/entity"
)

type Student struct {
	Leader  bool `json:"leader"`
	GroupID uint `json:"group_id" binding:"required"`
	UserID  uint `json:"user_id"`
}

type CreateStudent Student

type StudentResponse struct {
	*entity.Student
	CUD
}
