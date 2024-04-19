package entity

import (
	"time"

	"gorm.io/gorm"
)

type Group struct {
	ID        uint           `json:"id"                   gorm:"primarykey"`
	Name      string         `json:"name"                 gorm:"unique"`
	FacultyID *uint          `json:"faculty_id,omitempty"`
	CuratorID *uint          `json:"curator_id,omitempty"`
	Students  []Student      `json:"students"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
