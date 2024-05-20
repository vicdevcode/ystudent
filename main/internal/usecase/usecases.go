package usecase

import (
	"github.com/vicdevcode/ystudent/main/internal/usecase/repo"
	"github.com/vicdevcode/ystudent/main/pkg/config"
	"github.com/vicdevcode/ystudent/main/pkg/postgres"
)

type UseCases struct {
	UserUseCase       User
	StudentUseCase    Student
	TeacherUseCase    Teacher
	EmployeeUseCase   Employee
	FacultyUseCase    Faculty
	GroupUseCase      Group
	DepartmentUseCase Department
	JwtUseCase        Jwt
}

func New(cfg *config.Config, db *postgres.Postgres) UseCases {
	t := cfg.ContextTimeout
	return UseCases{
		UserUseCase:       newUser(repo.NewUser(db), t),
		StudentUseCase:    newStudent(repo.NewStudent(db), repo.NewUser(db), t),
		TeacherUseCase:    newTeacher(repo.NewTeacher(db), repo.NewUser(db), t),
		EmployeeUseCase:   newEmployee(repo.NewEmployee(db), t),
		FacultyUseCase:    newFaculty(repo.NewFaculty(db), t),
		GroupUseCase:      newGroup(repo.NewGroup(db), t),
		DepartmentUseCase: newDepartment(repo.NewDepartment(db), repo.NewEmployee(db), t),
		JwtUseCase:        newJwt(JwtConfig(cfg.JWT)),
	}
}
