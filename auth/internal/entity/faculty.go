package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Faculty struct {
	ID        uuid.UUID      `json:"id" gorm:"uuid;default:gen_random_uuid();primarykey"`
	Name      string         `json:"name" gorm:"unique"`
	Groups    []Group        `json:"groups,omitempty" gorm:"foreignKey:FacultyID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
