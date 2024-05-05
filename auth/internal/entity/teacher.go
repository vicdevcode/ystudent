package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Teacher struct {
	ID        uuid.UUID      `json:"id"                   gorm:"uuid;default:gen_random_uuid();primarykey"`
	Groups    []Group        `json:"groups,omitempty"     gorm:"foreignKey:CuratorID"`
	UserID    uuid.UUID      `json:"user_id"              gorm:"uuid"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
