package usecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/main/internal/dto"
	"github.com/vicdevcode/ystudent/main/internal/entity"
)

type (
	// User
	User interface {
		Create(context.Context, dto.CreateUser) (*entity.User, error)
		FindAll(context.Context, entity.UserRole) ([]entity.User, error)
		FindOne(context.Context, entity.UserRole, entity.User) (*entity.User, error)
	}
	UserRepo interface {
		Create(context.Context, dto.CreateUser) (*entity.User, error)
		FindAll(context.Context, entity.UserRole) ([]entity.User, error)
		FindOneByID(context.Context, entity.UserRole, uuid.UUID) (*entity.User, error)
		FindOneByEmail(context.Context, entity.UserRole, string) (*entity.User, error)
		Delete(context.Context, string) error
	}
	// Student
	Student interface {
		Create(context.Context, dto.CreateStudent) (*entity.Student, error)
		FindAll(context.Context) ([]entity.Student, error)
	}
	StudentRepo interface {
		Create(context.Context, dto.CreateStudent) (*entity.Student, error)
		FindAll(context.Context) ([]entity.Student, error)
	}
	// Teacher
	Teacher interface {
		Create(context.Context, dto.CreateTeacher) (*entity.Teacher, error)
		FindAll(context.Context) ([]entity.Teacher, error)
	}
	TeacherRepo interface {
		Create(context.Context, dto.CreateTeacher) (*entity.Teacher, error)
		FindAll(context.Context) ([]entity.Teacher, error)
	}
	// Employee
	Employee interface {
		Create(context.Context, dto.CreateEmployee) (*entity.Employee, error)
		FindAll(context.Context) ([]entity.Employee, error)
		FindOne(context.Context, entity.Employee) (*entity.Employee, error)
	}
	EmployeeRepo interface {
		Create(context.Context, dto.CreateEmployee) (*entity.Employee, error)
		FindAll(context.Context) ([]entity.Employee, error)
		FindOneByID(context.Context, uuid.UUID) (*entity.Employee, error)
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
	// Department
	Department interface {
		Create(context.Context, dto.CreateDepartment) (*entity.Department, error)
		FindAll(context.Context) ([]entity.Department, error)
		AddEmployee(context.Context, dto.AddEmployeeToDepartment) (*entity.Department, error)
		DeleteEmployee(
			context.Context,
			dto.DeleteEmployeeFromDepartment,
		) (*entity.Department, error)
	}
	DepartmentRepo interface {
		Create(context.Context, dto.CreateDepartment) (*entity.Department, error)
		FindAll(context.Context) ([]entity.Department, error)
		FindOneByID(context.Context, uuid.UUID) (*entity.Department, error)
		AddEmployee(context.Context, dto.AddEmployeeToDepartment) (*entity.Department, error)
		DeleteEmployee(
			context.Context,
			*entity.Department,
			*entity.Employee,
		) (*entity.Department, error)
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
		FindOneByID(context.Context, uuid.UUID) (*entity.Group, error)
		FindOneByName(context.Context, string) (*entity.Group, error)
		FindAll(context.Context) ([]entity.Group, error)
		UpdateCurator(context.Context, dto.UpdateGroupCurator) (*entity.Group, error)
	}
	// Jwt
	Jwt interface {
		ExtractFromToken(string, string, bool) (string, error)
	}
)