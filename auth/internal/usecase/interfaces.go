package usecase

import (
	"context"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
)

type (
	// Admin
	Admin interface {
		FindOne(context.Context, entity.Admin) (*entity.Admin, error)
		UpdateRefreshToken(context.Context, dto.UpdateRefreshToken) (*entity.Admin, error)
	}
	AdminRepo interface {
		FindOneByRefreshToken(context.Context, string) (*entity.Admin, error)
		FindOneByLogin(context.Context, string) (*entity.Admin, error)
		FindOneByID(context.Context, uint) (*entity.Admin, error)
		UpdateRefreshToken(context.Context, dto.UpdateRefreshToken) (*entity.Admin, error)
	}
	// User
	User interface {
		FindAll(context.Context) ([]entity.User, error)
		FindOne(context.Context, entity.User) (*entity.User, error)
		Create(context.Context, dto.CreateUser) (*entity.User, error)
		UpdateRefreshToken(context.Context, dto.UpdateRefreshToken) (*entity.User, error)
	}
	UserRepo interface {
		FindAll(context.Context) ([]entity.User, error)
		FindAllByIDs(context.Context, []uint) ([]entity.User, error)
		FindOneByID(context.Context, uint) (*entity.User, error)
		FindOneByEmail(context.Context, string) (*entity.User, error)
		FindOneByRefreshToken(context.Context, string) (*entity.User, error)
		Create(context.Context, dto.CreateUser) (*entity.User, error)
		Delete(context.Context, string) error
		UpdateRefreshToken(context.Context, dto.UpdateRefreshToken) (*entity.User, error)
	}
	// Student
	Student interface {
		Create(context.Context, dto.CreateStudent) (*entity.Student, error)
		FindAll(context.Context) ([]entity.User, error)
	}
	StudentRepo interface {
		Create(context.Context, dto.CreateStudent) (*entity.Student, error)
		FindAll(context.Context) ([]entity.Student, error)
	}
	// Teacher
	Teacher interface {
		Create(context.Context, dto.CreateTeacher) (*entity.Teacher, error)
		FindAll(context.Context) ([]entity.User, error)
	}
	TeacherRepo interface {
		Create(context.Context, dto.CreateTeacher) (*entity.Teacher, error)
		FindAll(context.Context) ([]entity.Teacher, error)
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
		FindOne(context.Context, entity.Group) (*entity.Group, error)
		FindAll(context.Context) ([]entity.Group, error)
		UpdateCurator(context.Context, dto.UpdateGroupCurator) (*entity.Group, error)
	}
	GroupRepo interface {
		Create(context.Context, dto.CreateGroup) (*entity.Group, error)
		FindOneByID(context.Context, uint) (*entity.Group, error)
		FindOneByName(context.Context, string) (*entity.Group, error)
		FindAll(context.Context) ([]entity.Group, error)
		UpdateCurator(context.Context, dto.UpdateGroupCurator) (*entity.Group, error)
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
