package usecase

import (
	"context"
	"time"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
)

type TeacherUseCase struct {
	teacherRepo TeacherRepo
	ctxTimeout  time.Duration
}

func newTeacher(tr TeacherRepo, t time.Duration) *TeacherUseCase {
	return &TeacherUseCase{
		teacherRepo: tr,
		ctxTimeout:  t,
	}
}

func (uc *TeacherUseCase) Create(c context.Context, data dto.CreateTeacher) (*entity.Teacher, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	teacher, err := uc.teacherRepo.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return teacher, nil
}
