package dto

import "github.com/google/uuid"

type Teacher struct {
	UserID uuid.UUID `json:"user_id" binding:"omitempty,uuid"`
}

type CreateTeacher Teacher
