package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Department struct {
	ID        uuid.UUID      `json:"id"                   gorm:"uuid;default:gen_random_uuid();primarykey"`
	Name      string         `json:"name"                 gorm:"unique"`
	FacultyID *uuid.UUID     `json:"faculty_id,omitempty" gorm:"uuid"`
	Employees []*Employee    `json:"employees,omitempty"  gorm:"many2many:department_employees;"`
	Groups    []Group        `json:"groups,omitempty"     gorm:"foreignKey:DepartmentID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
