package dto

type CUD struct {
	CreatedAt interface{} `json:"created_at,omitempty"`
	UpdatedAt interface{} `json:"updated_at,omitempty"`
	DeletedAt interface{} `json:"deleted_at,omitempty"`
}
