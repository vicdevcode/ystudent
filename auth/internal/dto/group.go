package dto

import "github.com/google/uuid"

type Group struct {
	Name      string `json:"name" binding:"required"`
	FacultyID string `json:"faculty_id" binding:"required,uuid"`
}

type CreateGroup Group

type UpdateGroupCurator struct {
	ID        uuid.UUID `json:"id" binding:"uuid,required"`
	CuratorID uuid.UUID `json:"curator_id" binding:"required,uuid"`
}
