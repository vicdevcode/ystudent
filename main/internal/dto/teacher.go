package dto

type Teacher struct {
	Fio
	Email string `json:"email" binding:"required,email"`
}

type CreateTeacher Teacher
