package entity

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID        uint           `json:"id"                   gorm:"primarykey"`
	GroupID   uint           `json:"group_id"`
	UserID    uint           `json:"user_id"`
	Leader    bool           `json:"leader"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
