package dto

type Faculty struct {
	Name string `json:"name" binding:"required"`
}

type CreateFaculty Faculty
