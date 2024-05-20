package entity

import (
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole string

const (
	ADMIN     UserRole = "ADMIN"
	STUDENT   UserRole = "STUDENT"
	TEACHER   UserRole = "TEACHER"
	EMPLOYEE  UserRole = "EMPLOYEE"
	MODERATOR UserRole = "MODERATOR"
)

func (ut *UserRole) Scan(value UserRole) error {
	*ut = value
	return nil
}

func (ut UserRole) Value() (driver.Value, error) {
	return string(ut), nil
}

type User struct {
	ID         uuid.UUID      `json:"id"                   gorm:"uuid;default:gen_random_uuid();primarykey"`
	Firstname  string         `json:"firstname"`
	Middlename string         `json:"middlename,omitempty"`
	Surname    string         `json:"surname"`
	Email      string         `json:"email"                gorm:"unique"`
	RoleType   UserRole       `json:"role"                 gorm:"type:user_role"`
	RoleID     *uuid.UUID     `json:"role_id,omitempty"    gorm:"uuid"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
