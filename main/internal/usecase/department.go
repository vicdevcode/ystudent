package usecase

import (
	"context"
	"time"

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

func (uc *DepartmentUseCase) FindAll(c context.Context) ([]entity.Department, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	return uc.departmentRepo.FindAll(ctx)
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
