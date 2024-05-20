package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Employee struct {
	ID          uuid.UUID      `json:"id"                    gorm:"uuid;default:gen_random_uuid();primarykey"`
	Departments []*Department  `json:"departments,omitempty" gorm:"many2many:department_employees;"`
	User        User           `json:"user"                  gorm:"polymorphic:Role;polymorphicValue:EMPLOYEE"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty"  gorm:"index"`
}
