package dto

type User struct {
	Fio
	Student Student `json:"student,omitempty"`
}

type Fio struct {
	Firstname  string `json:"firstname" binding:"required,alpha"`
	Middlename string `json:"middlename" binding:"omitempty,alpha"`
	Surname    string `json:"surname" binding:"required,alpha"`
}

type CreateUser struct {
	Fio
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=30"`
}

type CreateUserWithStudent struct {
	CreateUser
	CreateStudent
}
