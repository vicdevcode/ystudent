package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/main/internal/dto"
	"github.com/vicdevcode/ystudent/main/internal/entity"
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

func (uc *FacultyUseCase) Create(
	c context.Context,
	data dto.CreateFaculty,
) (*entity.Faculty, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	faculty, err := uc.repo.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return faculty, nil
}

func (uc *FacultyUseCase) FindAll(c context.Context, page dto.Page) ([]entity.Faculty, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	faculties, err := uc.repo.FindAll(ctx, page)
	if err != nil {
		return nil, err
	}
	return faculties, nil
}

func (uc *FacultyUseCase) FindOne(c context.Context, data entity.Faculty) (*entity.Faculty, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	if uuid.Nil != data.ID {
		return uc.repo.FindOneByID(ctx, data.ID)
	}
	return nil, errors.New("record not found")
}

func (uc *FacultyUseCase) Update(
	c context.Context,
	data dto.UpdateFaculty,
) (*entity.Faculty, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	return uc.repo.Update(ctx, data)
}

func (uc *FacultyUseCase) Delete(
	c context.Context,
	id uuid.UUID,
) error {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	return uc.repo.Delete(ctx, id)
}
