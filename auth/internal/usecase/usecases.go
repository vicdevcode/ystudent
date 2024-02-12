package usecase

import (
	"github.com/vicdevcode/ystudent/auth/internal/usecase/repo"
	"github.com/vicdevcode/ystudent/auth/pkg/config"
	"github.com/vicdevcode/ystudent/auth/pkg/postgres"
)

type UseCases struct {
	UserUseCase    User
	StudentUseCase Student
	TeacherUseCase Teacher
	FacultyUseCase Faculty
	GroupUseCase   Group
}

func New(cfg *config.Config, db *postgres.Postgres) UseCases {
	t := cfg.ContextTimeout
	return UseCases{
		UserUseCase:    NewUser(repo.NewUser(db), t),
		StudentUseCase: NewStudent(repo.NewUser(db), repo.NewStudent(db), t),
		TeacherUseCase: NewTeacher(repo.NewUser(db), repo.NewTeacher(db), t),
		FacultyUseCase: NewFaculty(repo.NewFaculty(db), t),
		GroupUseCase:   NewGroup(repo.NewGroup(db), t),
	}
}
