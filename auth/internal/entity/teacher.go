package entity

import (
	"time"

	"gorm.io/gorm"
)

type Teacher struct {
	ID        uint           `json:"id"                   gorm:"primarykey"`
	Groups    []Group        `json:"groups,omitempty"     gorm:"foreignKey:CuratorID"`
	UserID    uint           `json:"user_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
