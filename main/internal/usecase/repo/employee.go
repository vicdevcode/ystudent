package repo

import (
	"context"

	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/main/internal/dto"
	"github.com/vicdevcode/ystudent/main/internal/entity"
	"github.com/vicdevcode/ystudent/main/pkg/postgres"
)

type EmployeeRepo struct {
	*postgres.Postgres
}

func NewEmployee(db *postgres.Postgres) *EmployeeRepo {
	return &EmployeeRepo{db}
}

func (r *EmployeeRepo) Create(
	ctx context.Context,
	data dto.CreateEmployee,
) (*entity.Employee, error) {
	employee := &entity.Employee{
		User: entity.User{
			Firstname:  data.Firstname,
			Middlename: data.Middlename,
			Surname:    data.Surname,
			Email:      data.Email,
		},
	}
	if err := r.WithContext(ctx).Create(employee).Error; err != nil {
		return nil, err
	}
	return employee, nil
}

func (r *EmployeeRepo) FindAll(ctx context.Context, page dto.Page) ([]entity.Employee, error) {
	var employees []entity.Employee

	if err := r.WithContext(ctx).Limit(page.Count).Offset((page.Page - 1) * page.Count).Preload("User").Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

func (r *EmployeeRepo) FindOneByID(ctx context.Context, id uuid.UUID) (*entity.Employee, error) {
	var employee *entity.Employee

	if err := r.WithContext(ctx).Where("id = ?", id).First(&employee).Error; err != nil {
		return nil, err
	}

	return employee, nil
}

func (r *EmployeeRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return r.WithContext(ctx).Unscoped().Delete(&entity.Employee{ID: id}).Error
}
