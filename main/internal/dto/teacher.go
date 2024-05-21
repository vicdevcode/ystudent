package dto

import "github.com/vicdevcode/ystudent/main/internal/entity"

type Teacher struct {
	Fio
	Email string `json:"email" binding:"required,email"`
}

type CreateTeacher Teacher

type TeacherResponse struct {
	*entity.Teacher
	Password string `json:"password"`
	CUD
}

type TeacherUserResponse struct {
	*entity.User
	Password string `json:"password"`
	CUD
}
