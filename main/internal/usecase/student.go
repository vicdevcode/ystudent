package usecase

import (
	"context"
	"time"

	"github.com/vicdevcode/ystudent/main/internal/dto"
	"github.com/vicdevcode/ystudent/main/internal/entity"
)

type StudentUseCase struct {
	studentRepo StudentRepo
	userRepo    UserRepo
	ctxTimeout  time.Duration
}

func newStudent(sr StudentRepo, ur UserRepo, t time.Duration) *StudentUseCase {
	return &StudentUseCase{
		studentRepo: sr,
		userRepo:    ur,
		ctxTimeout:  t,
	}
}

func (uc *StudentUseCase) Create(
	c context.Context,
	data dto.CreateStudent,
) (*entity.Student, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	return uc.studentRepo.Create(ctx, data)
}

func (uc *StudentUseCase) FindAll(c context.Context) ([]entity.Student, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	return uc.studentRepo.FindAll(ctx)
}
