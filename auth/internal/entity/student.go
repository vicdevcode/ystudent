package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	ID        uuid.UUID      `json:"id" gorm:"uuid;default:gen_random_uuid();primarykey"`
	GroupID   uuid.UUID      `json:"group_id" gorm:"uuid"`
	UserID    uuid.UUID      `json:"user_id" gorm:"uuid"`
	Leader    bool           `json:"leader"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
