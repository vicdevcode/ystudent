package usecase

import (
	"context"
	"time"

	"github.com/vicdevcode/ystudent/main/internal/dto"
	"github.com/vicdevcode/ystudent/main/internal/entity"
)

type TeacherUseCase struct {
	teacherRepo TeacherRepo
	userRepo    UserRepo
	ctxTimeout  time.Duration
}

func newTeacher(tr TeacherRepo, ur UserRepo, t time.Duration) *TeacherUseCase {
	return &TeacherUseCase{
		teacherRepo: tr,
		userRepo:    ur,
		ctxTimeout:  t,
	}
}

func (uc *TeacherUseCase) Create(
	c context.Context,
	data dto.CreateTeacher,
) (*entity.Teacher, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	return uc.teacherRepo.Create(ctx, data)
}

func (uc *TeacherUseCase) FindAll(c context.Context, page dto.Page) ([]entity.Teacher, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	return uc.teacherRepo.FindAll(ctx, page)
}
