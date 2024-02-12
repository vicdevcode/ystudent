package dto

type Group struct {
	Name      string `json:"name" binding:"required"`
	FacultyID string `json:"faculty_id" binding:"required,uuid"`
}

type CreateGroup Group
