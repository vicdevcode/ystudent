package dto

import "github.com/google/uuid"

type TokenPayload struct {
	ID    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
	Role  string `json:"role,omitempty"`
}

type AccessTokenPayload struct {
	ID    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
	Role  string `json:"role,omitempty"`
}

type RefreshTokenPayload struct {
	ID    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
	Role  string `json:"role,omitempty"`
}

type UpdateRefreshToken struct {
	ID           uuid.UUID `json:"id"`
	RefreshToken string    `json:"refresh_token"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
