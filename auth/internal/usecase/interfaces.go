package usecase

import (
	"context"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
)

type (
	// User
	User interface {
		FindAll(context.Context) ([]entity.User, error)
	}
	UserRepo interface {
		FindAll(context.Context) ([]entity.User, error)
		Create(context.Context, dto.CreateUser) (*entity.User, error)
		Delete(context.Context, string) error
	}
)
