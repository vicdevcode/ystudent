package usecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/main/internal/dto"
	"github.com/vicdevcode/ystudent/main/internal/entity"
)

type (
	// Faculty
	Faculty interface {
		Create(context.Context, dto.CreateFaculty) (*entity.Faculty, error)
		FindAll(context.Context, dto.Page) ([]entity.Faculty, error)
		FindOne(context.Context, entity.Faculty) (*entity.Faculty, error)
		Update(context.Context, dto.UpdateFaculty) (*entity.Faculty, error)
		Delete(context.Context, uuid.UUID) error
	}
	FacultyRepo interface {
		Create(context.Context, dto.CreateFaculty) (*entity.Faculty, error)
		FindAll(context.Context, dto.Page) ([]entity.Faculty, error)
		FindOneByID(context.Context, uuid.UUID) (*entity.Faculty, error)
		Update(context.Context, dto.UpdateFaculty) (*entity.Faculty, error)
		Delete(context.Context, uuid.UUID) error
	}
	// Department
	Department interface {
		Create(context.Context, dto.CreateDepartment) (*entity.Department, error)
		FindAll(context.Context, dto.Page) ([]entity.Department, error)
		FindOne(context.Context, entity.Department) (*entity.Department, error)
		Update(context.Context, dto.UpdateDepartment) (*entity.Department, error)
		Delete(context.Context, uuid.UUID) error
		AddEmployee(context.Context, dto.AddEmployeeToDepartment) (*entity.Department, error)
		DeleteEmployee(
			context.Context,
			dto.DeleteEmployeeFromDepartment,
		) (*entity.Department, error)
	}
	DepartmentRepo interface {
		Create(context.Context, dto.CreateDepartment) (*entity.Department, error)
		FindAll(context.Context, dto.Page) ([]entity.Department, error)
		FindOneByID(context.Context, uuid.UUID) (*entity.Department, error)
		Update(context.Context, dto.UpdateDepartment) (*entity.Department, error)
		Delete(context.Context, uuid.UUID) error
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
		FindAll(context.Context, dto.Page) ([]entity.Group, error)
		Update(context.Context, dto.UpdateGroup) (*entity.Group, error)
		UpdateCurator(context.Context, dto.UpdateGroupCurator) (*entity.Group, error)
		Delete(context.Context, uuid.UUID) error
	}
	GroupRepo interface {
		Create(context.Context, dto.CreateGroup) (*entity.Group, error)
		FindOneByID(context.Context, uuid.UUID) (*entity.Group, error)
		FindOneByName(context.Context, string) (*entity.Group, error)
		FindAll(context.Context, dto.Page) ([]entity.Group, error)
		Update(context.Context, dto.UpdateGroup) (*entity.Group, error)
		UpdateCurator(context.Context, dto.UpdateGroupCurator) (*entity.Group, error)
		Delete(context.Context, uuid.UUID) error
	}
	// User
	User interface {
		Create(context.Context, dto.CreateUser) (*entity.User, error)
		FindAll(context.Context, entity.UserRole, dto.Page) ([]entity.User, error)
		FindOne(context.Context, entity.UserRole, entity.User) (*entity.User, error)
		Update(context.Context, dto.UpdateUser) (*entity.User, error)
		Delete(context.Context, uuid.UUID) error
	}
	UserRepo interface {
		Create(context.Context, dto.CreateUser) (*entity.User, error)
		FindAll(context.Context, entity.UserRole, dto.Page) ([]entity.User, error)
		FindOneByID(context.Context, entity.UserRole, uuid.UUID) (*entity.User, error)
		FindOneByEmail(context.Context, entity.UserRole, string) (*entity.User, error)
		Update(context.Context, dto.UpdateUser) (*entity.User, error)
		Delete(context.Context, uuid.UUID) error
	}
	// Student
	Student interface {
		Create(context.Context, dto.CreateStudent) (*entity.Student, error)
		FindAll(context.Context, dto.Page) ([]entity.Student, error)
		FindOne(context.Context, entity.Student) (*entity.Student, error)
		Update(context.Context, dto.UpdateStudent) (*entity.Student, error)
		Delete(context.Context, uuid.UUID) error
	}
	StudentRepo interface {
		Create(context.Context, dto.CreateStudent) (*entity.Student, error)
		FindAll(context.Context, dto.Page) ([]entity.Student, error)
		FindOneByID(context.Context, uuid.UUID) (*entity.Student, error)
		Update(context.Context, dto.UpdateStudent) (*entity.Student, error)
		Delete(context.Context, uuid.UUID) error
	}
	// Teacher
	Teacher interface {
		Create(context.Context, dto.CreateTeacher) (*entity.Teacher, error)
		FindAll(context.Context, dto.Page) ([]entity.Teacher, error)
		FindOne(context.Context, entity.Teacher) (*entity.Teacher, error)
		Delete(context.Context, uuid.UUID) error
		AddGroup(context.Context, dto.AddGroupToTeacher) (*entity.Teacher, error)
		DeleteGroup(context.Context, dto.DeleteGroupFromTeacher) (*entity.Teacher, error)
	}
	TeacherRepo interface {
		Create(context.Context, dto.CreateTeacher) (*entity.Teacher, error)
		FindAll(context.Context, dto.Page) ([]entity.Teacher, error)
		FindOneByID(context.Context, uuid.UUID) (*entity.Teacher, error)
		Delete(context.Context, uuid.UUID) error
		AddGroup(context.Context, dto.AddGroupToTeacher) (*entity.Teacher, error)
		DeleteGroup(context.Context,
			*entity.Teacher,
			*entity.Group,
		) (*entity.Teacher, error)
	}
	// Employee
	Employee interface {
		Create(context.Context, dto.CreateEmployee) (*entity.Employee, error)
		FindAll(context.Context, dto.Page) ([]entity.Employee, error)
		FindOne(context.Context, entity.Employee) (*entity.Employee, error)
		Delete(context.Context, uuid.UUID) error
	}
	EmployeeRepo interface {
		Create(context.Context, dto.CreateEmployee) (*entity.Employee, error)
		FindAll(context.Context, dto.Page) ([]entity.Employee, error)
		FindOneByID(context.Context, uuid.UUID) (*entity.Employee, error)
		Delete(context.Context, uuid.UUID) error
	}
	// Jwt
	Jwt interface {
		ExtractFromToken(string, string, bool) (string, error)
	}
)
