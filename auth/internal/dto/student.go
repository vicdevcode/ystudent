package dto

type Student struct {
	Leader  bool   `json:"leader"`
	GroupID string `json:"group_id" binding:"required,uuid"`
	UserID  string `json:"user_id"`
}

type CreateStudent Student
