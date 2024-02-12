package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID         uuid.UUID      `json:"id" gorm:"uuid;default:gen_random_uuid();primarykey"`
	Firstname  string         `json:"firstname"`
	Middlename string         `json:"middlename,omitempty"`
	Surname    string         `json:"surname"`
	Email      string         `json:"email" gorm:"unique"`
	Password   string         `json:"password"`
	Student    *Student       `json:"student,omitempty"`
	Teacher    *Teacher       `json:"teacher,omitempty"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
