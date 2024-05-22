package repo

import (
	"context"

	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/main/internal/dto"
	"github.com/vicdevcode/ystudent/main/internal/entity"
	"github.com/vicdevcode/ystudent/main/pkg/postgres"
)

type DepartmentRepo struct {
	*postgres.Postgres
}

func NewDepartment(db *postgres.Postgres) *DepartmentRepo {
	return &DepartmentRepo{db}
}

func (r *DepartmentRepo) Create(
	ctx context.Context,
	data dto.CreateDepartment,
) (*entity.Department, error) {
	department := &entity.Department{
		Name:      data.Name,
		FacultyID: data.FacultyID,
	}
	if err := r.WithContext(ctx).Create(department).Error; err != nil {
		return nil, err
	}
	return department, nil
}

func (r *DepartmentRepo) FindAll(ctx context.Context) ([]entity.Department, error) {
	var departments []entity.Department
	if err := r.WithContext(ctx).Preload("Employees.User").Preload("Groups.Students.User").Find(&departments).Error; err != nil {
		return nil, err
	}
	return departments, nil
}

func (r *DepartmentRepo) FindOneByID(
	ctx context.Context,
	id uuid.UUID,
) (*entity.Department, error) {
	var department *entity.Department
	if err := r.WithContext(ctx).Preload("Employees.User").Preload("Groups.Students.User").Where("id = ?", id).First(&department).Error; err != nil {
		return nil, err
	}
	return department, nil
}

func (r *DepartmentRepo) Update(
	ctx context.Context,
	data dto.UpdateDepartment,
) (*entity.Department, error) {
	department := &entity.Department{ID: data.ID}
	if err := r.WithContext(ctx).Model(department).Updates(entity.Department{
		Name:      data.Name,
		FacultyID: data.FacultyID,
	}).Error; err != nil {
		return nil, err
	}
	return department, nil
}

func (r *DepartmentRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return r.WithContext(ctx).Unscoped().Delete(&entity.Department{ID: id}).Error
}

func (r *DepartmentRepo) AddEmployee(
	ctx context.Context,
	data dto.AddEmployeeToDepartment,
) (*entity.Department, error) {
	department := &entity.Department{
		ID: data.DepartmentID,
	}
	if err := r.WithContext(ctx).Model(&department).Association("Employees").Append(&entity.Employee{
		ID: data.EmployeeID,
	}); err != nil {
		return nil, err
	}

	var err error

	department, err = r.FindOneByID(ctx, data.DepartmentID)
	if err != nil {
		return nil, err
	}

	return department, nil
}

func (r *DepartmentRepo) DeleteEmployee(
	ctx context.Context,
	department *entity.Department,
	employee *entity.Employee,
) (*entity.Department, error) {
	if err := r.WithContext(ctx).Model(&department).Association("Employees").Delete(&employee); err != nil {
		return nil, err
	}

	return department, nil
}
