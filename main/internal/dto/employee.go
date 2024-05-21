package dto

import "github.com/vicdevcode/ystudent/main/internal/entity"

type Employee struct {
	Fio
	Email string `json:"email" binding:"required,email"`
}

type CreateEmployee Employee

type EmployeeResponse struct {
	*entity.Employee
	Password string `json:"password"`
	CUD
}

type EmployeeUserResponse struct {
	*entity.User
	Password string `json:"password"`
	CUD
}
