package entity

import (
	"time"

	"gorm.io/gorm"
)

type Faculty struct {
	ID        uint           `json:"id"                   gorm:"primarykey"`
	Name      string         `json:"name"                 gorm:"unique"`
	Groups    []Group        `json:"groups,omitempty"     gorm:"foreignKey:FacultyID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
