package usecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
)

type (
	// Admin
	Admin interface {
		FindOne(context.Context, string) (*entity.Admin, error)
		UpdateRefreshToken(context.Context, dto.UpdateRefreshToken) (*entity.Admin, error)
	}
	AdminRepo interface {
		FindOne(context.Context, string) (*entity.Admin, error)
		UpdateRefreshToken(context.Context, dto.UpdateRefreshToken) (*entity.Admin, error)
	}
	// Jwt
	Jwt interface {
		CreateToken(dto.TokenPayload, bool) (string, error)
		IsAuthorized(string) (bool, error)
		ExtractIDFromToken(string) (string, error)
	}
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
	// Hash
	Hash interface {
		HashPassword(string) (string, error)
		CheckPasswordHash(string, string) bool
	}
	// Student
	Student interface {
		SignUp(context.Context, dto.CreateUserAndStudent) (*entity.Student, error)
	}
	StudentRepo interface {
		Create(context.Context, dto.CreateStudent) (*entity.Student, error)
	}
	// Teacher
	Teacher interface {
		SignUp(context.Context, dto.CreateUser) (*entity.Teacher, error)
	}
	TeacherRepo interface {
		Create(context.Context, dto.CreateTeacher) (*entity.Teacher, error)
	}
	// Faculty
	Faculty interface {
		Create(context.Context, dto.CreateFaculty) (*entity.Faculty, error)
		FindAll(context.Context) ([]entity.Faculty, error)
	}
	FacultyRepo interface {
		Create(context.Context, dto.CreateFaculty) (*entity.Faculty, error)
		FindAll(context.Context) ([]entity.Faculty, error)
	}
	// Group
	Group interface {
		Create(context.Context, dto.CreateGroup) (*entity.Group, error)
		FindAll(context.Context) ([]entity.Group, error)
	}
	GroupRepo interface {
		Create(context.Context, dto.CreateGroup) (*entity.Group, error)
		FindAll(context.Context) ([]entity.Group, error)
	}
)
