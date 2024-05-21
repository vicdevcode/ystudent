package usecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
)

type (
	// User
	User interface {
		FindAll(context.Context) ([]entity.User, error)
		FindOne(context.Context, entity.User) (*entity.User, error)
		Create(context.Context, dto.CreateUser) (*entity.User, error)
		UpdateRefreshToken(context.Context, dto.UpdateRefreshToken) (*entity.User, error)
	}
	UserRepo interface {
		FindAll(context.Context) ([]entity.User, error)
		FindAllByIDs(context.Context, []uuid.UUID) ([]entity.User, error)
		FindOneByID(context.Context, uuid.UUID) (*entity.User, error)
		FindOneByEmail(context.Context, string) (*entity.User, error)
		FindOneByRefreshToken(context.Context, string) (*entity.User, error)
		Create(context.Context, dto.CreateUser) (*entity.User, error)
		Delete(context.Context, string) error
		UpdateRefreshToken(context.Context, dto.UpdateRefreshToken) (*entity.User, error)
	}
	// Hash
	Hash interface {
		HashPassword(string) (string, error)
		CheckPasswordHash(string, string) bool
	}
	// Jwt
	Jwt interface {
		CreateAccessToken(dto.AccessTokenPayload) (string, error)
		CreateRefreshToken(dto.RefreshTokenPayload) (string, error)
		IsTokenValid(string, bool) (bool, error)
		ExtractFromToken(string, string, bool) (string, error)
		CreateTokens(dto.AccessTokenPayload, dto.RefreshTokenPayload) (*dto.Tokens, error)
	}
)
