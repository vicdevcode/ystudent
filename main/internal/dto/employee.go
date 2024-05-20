package dto

type Employee struct {
	Fio
	Email string `json:"email" binding:"required,email"`
}

type CreateEmployee Employee
