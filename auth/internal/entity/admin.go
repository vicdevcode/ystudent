package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	ID           uuid.UUID      `json:"id"                   gorm:"uuid;default:gen_random_uuid();primarykey"`
	Login        string         `json:"login"                gorm:"unique"`
	Password     string         `json:"password"`
	RefreshToken string         `json:"refresh_token"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
