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
		FindOne(context.Context, uuid.UUID) (*entity.User, error)
		SignUp(context.Context, dto.CreateUser) (*entity.User, error)
	}
	UserRepo interface {
		FindAll(context.Context) ([]entity.User, error)
		FindOne(context.Context, uuid.UUID) (*entity.User, error)
		Create(context.Context, dto.CreateUser) (*entity.User, error)
		Delete(context.Context, string) error
	}
	// Student
	Student interface {
		SignUp(context.Context, dto.CreateUserWithStudent) (*entity.Student, error)
	}
	StudentRepo interface {
		Create(context.Context, dto.CreateStudent) (*entity.Student, error)
	}
)
