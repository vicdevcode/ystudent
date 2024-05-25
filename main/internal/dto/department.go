package dto

import (
	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/main/internal/entity"
)

type Department struct {
	Name      string     `json:"name"       binding:"required"`
	FacultyID *uuid.UUID `json:"faculty_id" binding:"required,uuid"`
}

type FindAllDepartmentResponse struct {
	FacultyName string `json:"faculty_name"`
	*entity.Department
}

type CreateDepartment Department

type UpdateDepartmentBody struct {
	Name      string     `json:"name"       binding:"omitempty"`
	FacultyID *uuid.UUID `json:"faculty_id" binding:"omitempty,uuid"`
}

type UpdateDepartment struct {
	ID        uuid.UUID  `json:"id"         binding:"required,uuid"`
	Name      string     `json:"name"       binding:"required"`
	FacultyID *uuid.UUID `json:"faculty_id" binding:"omitempty,uuid"`
}

type AddEmployeeToDepartment struct {
	DepartmentID uuid.UUID `json:"department_id" binding:"required,uuid"`
	EmployeeID   uuid.UUID `json:"employee_id"   binding:"required,uuid"`
}

type DeleteEmployeeFromDepartment AddEmployeeToDepartment
