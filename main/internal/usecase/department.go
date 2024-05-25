package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/main/internal/dto"
	"github.com/vicdevcode/ystudent/main/internal/entity"
)

type DepartmentUseCase struct {
	departmentRepo DepartmentRepo
	employeeRepo   EmployeeRepo
	ctxTimeout     time.Duration
}

func newDepartment(dr DepartmentRepo, er EmployeeRepo, t time.Duration) *DepartmentUseCase {
	return &DepartmentUseCase{
		departmentRepo: dr,
		employeeRepo:   er,
		ctxTimeout:     t,
	}
}

func (uc *DepartmentUseCase) Create(
	c context.Context,
	data dto.CreateDepartment,
) (*entity.Department, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	return uc.departmentRepo.Create(ctx, data)
}

func (uc *DepartmentUseCase) FindAll(
	c context.Context,
	page dto.Page,
) ([]entity.Department, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	return uc.departmentRepo.FindAll(ctx, page)
}

func (uc *DepartmentUseCase) FindOne(
	c context.Context,
	data entity.Department,
) (*entity.Department, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	if uuid.Nil != data.ID {
		return uc.departmentRepo.FindOneByID(ctx, data.ID)
	}

	return nil, errors.New("record not found")
}

func (uc *DepartmentUseCase) Update(
	c context.Context,
	data dto.UpdateDepartment,
) (*entity.Department, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	return uc.departmentRepo.Update(ctx, data)
}

func (uc *DepartmentUseCase) Delete(
	c context.Context,
	id uuid.UUID,
) error {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	return uc.departmentRepo.Delete(ctx, id)
}

func (uc *DepartmentUseCase) AddEmployee(
	c context.Context,
	data dto.AddEmployeeToDepartment,
) (*entity.Department, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	return uc.departmentRepo.AddEmployee(ctx, data)
}

func (uc *DepartmentUseCase) DeleteEmployee(
	c context.Context,
	data dto.DeleteEmployeeFromDepartment,
) (*entity.Department, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	employee, err := uc.employeeRepo.FindOneByID(ctx, data.EmployeeID)
	if err != nil {
		return nil, err
	}

	department, err := uc.departmentRepo.FindOneByID(ctx, data.DepartmentID)
	if err != nil {
		return nil, err
	}

	return uc.departmentRepo.DeleteEmployee(ctx, department, employee)
}
