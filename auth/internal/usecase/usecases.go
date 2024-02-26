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
	HashUseCase    Hash
	AdminUseCase   Admin
	JwtUseCase     Jwt
}

func New(cfg *config.Config, db *postgres.Postgres) UseCases {
	t := cfg.ContextTimeout
	return UseCases{
		UserUseCase:    newUser(repo.NewUser(db), t),
		StudentUseCase: newStudent(repo.NewUser(db), repo.NewStudent(db), t),
		TeacherUseCase: newTeacher(repo.NewUser(db), repo.NewTeacher(db), t),
		FacultyUseCase: newFaculty(repo.NewFaculty(db), t),
		GroupUseCase:   newGroup(repo.NewGroup(db), t),
		HashUseCase:    newHash(),
		AdminUseCase:   newAdmin(repo.NewAdmin(db), t),
		JwtUseCase:     newJwt(JwtConfig(cfg.JWT)),
	}
}
