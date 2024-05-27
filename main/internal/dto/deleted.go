package dto

import "github.com/google/uuid"

type Deleted struct {
	ID uuid.UUID `json:"id"`
}
