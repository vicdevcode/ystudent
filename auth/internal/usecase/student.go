package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
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

func (uc *StudentUseCase) Create(c context.Context, data dto.CreateStudent) (*entity.Student, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	student, err := uc.studentRepo.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (uc *StudentUseCase) FindAll(c context.Context) ([]entity.User, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	students, err := uc.studentRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var ids []uuid.UUID
	for _, student := range students {
		ids = append(ids, student.UserID)
	}

	users, err := uc.userRepo.FindAllByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	return users, nil
}
