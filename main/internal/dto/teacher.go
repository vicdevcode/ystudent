package dto

import (
	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/main/internal/entity"
)

type Teacher struct {
	Fio
	Email string `json:"email" binding:"required,email"`
}

type CreateTeacher Teacher

type UpdateTeacher struct {
	ID         uuid.UUID `json:"id"         binding:"required,uuid"`
	Firstname  string    `json:"firstname"  binding:"omitempty,alphaunicode"`
	Middlename string    `json:"middlename" binding:"omitempty,alphaunicode"`
	Surname    string    `json:"surname"    binding:"omitempty,alphaunicode"`
	Email      string    `json:"email"      binding:"omitempty,email"`
}

type UpdateTeacherBody struct {
	Firstname  string `json:"firstname"  binding:"omitempty,alphaunicode"`
	Middlename string `json:"middlename" binding:"omitempty,alphaunicode"`
	Surname    string `json:"surname"    binding:"omitempty,alphaunicode"`
	Email      string `json:"email"      binding:"omitempty,email"`
}

type AddGroupToTeacher struct {
	TeacherID uuid.UUID `json:"teacher_id" binding:"required,uuid"`
	GroupID   uuid.UUID `json:"group_id"   binding:"required,uuid"`
}

type DeleteGroupFromTeacher AddGroupToTeacher

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
