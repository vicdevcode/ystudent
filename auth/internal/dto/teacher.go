package dto

import "github.com/google/uuid"

type Teacher struct {
	UserID uuid.UUID `json:"user_id" binding:"uuid,omitempty"`
}

type CreateTeacher Teacher
