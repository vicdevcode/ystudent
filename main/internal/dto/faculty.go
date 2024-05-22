package dto

import "github.com/google/uuid"

type Faculty struct {
	Name string `json:"name" binding:"required"`
}

type CreateFaculty Faculty

type UpdateFaculty struct {
	ID   uuid.UUID `json:"id"   binding:"required,uuid"`
	Name string    `json:"name" binding:"required"`
}
