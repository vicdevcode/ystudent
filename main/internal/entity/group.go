package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Group struct {
	ID           uuid.UUID      `json:"id"                      gorm:"uuid;default:gen_random_uuid();primarykey"`
	Name         string         `json:"name"                    gorm:"unique"`
	DepartmentID *uuid.UUID     `json:"department_id,omitempty" gorm:"uuid"`
	CuratorID    *uuid.UUID     `json:"curator_id,omitempty"    gorm:"uuid"`
	Students     []Student      `json:"students"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at,omitempty"    gorm:"index"`
}
