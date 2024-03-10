package dto

import "github.com/vicdevcode/ystudent/auth/internal/entity"

type User struct {
	Fio
	Student Student `json:"student,omitempty"`
}

type Fio struct {
	Firstname  string `json:"firstname" binding:"required,alphaunicode"`
	Middlename string `json:"middlename" binding:"omitempty,alphaunicode"`
	Surname    string `json:"surname" binding:"required,alphaunicode"`
}

type CreateUser struct {
	Fio
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=30"`
}

type CreateUserAndStudent struct {
	CreateUser
	CreateStudent
}

type CreateUserAndTeacher struct {
	CreateTeacher
	CreateUser
}

type UserResponse struct {
	*entity.User
	CUD
	Password     interface{} `json:"password,omitempty"`
	RefreshToken interface{} `json:"refresh_token,omitempty"`
}
