package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/main/internal/dto"
	"github.com/vicdevcode/ystudent/main/internal/entity"
)

type EmployeeUseCase struct {
	repo       EmployeeRepo
	ctxTimeout time.Duration
}

func newEmployee(r EmployeeRepo, t time.Duration) *EmployeeUseCase {
	return &EmployeeUseCase{
		repo:       r,
		ctxTimeout: t,
	}
}

func (uc *EmployeeUseCase) Create(
	c context.Context,
	data dto.CreateEmployee,
) (*entity.Employee, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	return uc.repo.Create(ctx, data)
}

func (uc *EmployeeUseCase) FindAll(c context.Context) ([]entity.Employee, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	return uc.repo.FindAll(ctx)
}

func (uc *EmployeeUseCase) FindOne(
	c context.Context,
	data entity.Employee,
) (*entity.Employee, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	if uuid.Nil != data.ID {
		return uc.repo.FindOneByID(ctx, data.ID)
	}

	return nil, errors.New("record not found")
}
