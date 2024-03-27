package dto

type Person struct {
	Firstname  string `json:"firstname"  binding:"required,alpha"`
	Middlename string `json:"middlename" binding:"omitempty,alpha"`
	Surname    string `json:"surname"    binding:"required,alpha"`
}

type CreatePerson Person
