package usecase

import (
	"context"
	"time"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
)

type TeacherUseCase struct {
	userRepo    UserRepo
	teacherRepo TeacherRepo
	ctxTimeout  time.Duration
}

func NewTeacher(ur UserRepo, tr TeacherRepo, t time.Duration) *TeacherUseCase {
	return &TeacherUseCase{
		userRepo:    ur,
		teacherRepo: tr,
		ctxTimeout:  t,
	}
}

func (uc *TeacherUseCase) SignUp(c context.Context, data dto.CreateUser) (*entity.Teacher, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	user, err := uc.userRepo.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	teacherData := dto.CreateTeacher{UserID: user.ID}
	teacher, err := uc.teacherRepo.Create(ctx, teacherData)
	return teacher, nil
}
