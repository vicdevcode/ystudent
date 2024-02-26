package usecase

import (
	"context"
	"time"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
)

type FacultyUseCase struct {
	repo       FacultyRepo
	ctxTimeout time.Duration
}

func newFaculty(r FacultyRepo, t time.Duration) *FacultyUseCase {
	return &FacultyUseCase{
		repo:       r,
		ctxTimeout: t,
	}
}

func (uc *FacultyUseCase) Create(c context.Context, data dto.CreateFaculty) (*entity.Faculty, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	faculty, err := uc.repo.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return faculty, nil
}

func (uc *FacultyUseCase) FindAll(c context.Context) ([]entity.Faculty, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	faculties, err := uc.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return faculties, nil
}
