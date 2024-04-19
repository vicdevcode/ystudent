package dto

type Teacher struct {
	UserID uint `json:"user_id" binding:"omitempty"`
}

type CreateTeacher Teacher
