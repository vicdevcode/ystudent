package dto

import (
	"github.com/google/uuid"
)

type TokenPayload struct {
	ID    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}

type UpdateRefreshToken struct {
	ID           uuid.UUID `json:"id"`
	RefreshToken string    `json:"refresh_token"`
}
