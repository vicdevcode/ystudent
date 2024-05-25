package dto

import (
	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/main/internal/entity"
)

type Group struct {
	Name         string `json:"name"          binding:"required"`
	DepartmentID string `json:"department_id" binding:"required,uuid"`
	CuratorID    string `json:"curator_id"    binding:"required,uuid"`
}

type CreateGroup Group

type FindAllGroupResponse struct {
	DepartmentName string `json:"department_name"`
	CuratorFio     string `json:"curator_fio"`
	*entity.Group
}

type UpdateGroupBody struct {
	Name         string    `json:"name"          binding:"omitempty"`
	DepartmentID uuid.UUID `json:"department_id" binding:"omitempty,uuid"`
	CuratorID    uuid.UUID `json:"curator_id"    binding:"omitempty,uuid"`
}

type UpdateGroup struct {
	ID           uuid.UUID  `json:"id"`
	Name         string     `json:"name"`
	DepartmentID *uuid.UUID `json:"department_id"`
	CuratorID    *uuid.UUID `json:"curator_id"`
}

type UpdateGroupCurator struct {
	ID        uuid.UUID `json:"id"         binding:"required,uuid"`
	CuratorID uuid.UUID `json:"curator_id" binding:"required,uuid"`
}
