package repo

import (
	"context"

	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/main/internal/dto"
	"github.com/vicdevcode/ystudent/main/internal/entity"
	"github.com/vicdevcode/ystudent/main/pkg/postgres"
)

type FacultyRepo struct {
	*postgres.Postgres
}

func NewFaculty(db *postgres.Postgres) *FacultyRepo {
	return &FacultyRepo{db}
}

func (r *FacultyRepo) Create(ctx context.Context, data dto.CreateFaculty) (*entity.Faculty, error) {
	faculty := &entity.Faculty{Name: data.Name}
	if err := r.WithContext(ctx).Create(faculty).Error; err != nil {
		return nil, err
	}

	return faculty, nil
}

func (r *FacultyRepo) FindAll(ctx context.Context) ([]entity.Faculty, error) {
	var faculties []entity.Faculty
	if err := r.WithContext(ctx).Find(&faculties).Error; err != nil {
		return nil, err
	}
	return faculties, nil
}

func (r *FacultyRepo) Update(ctx context.Context, data dto.UpdateFaculty) (*entity.Faculty, error) {
	faculty := &entity.Faculty{ID: data.ID}
	if err := r.WithContext(ctx).Model(faculty).Updates(entity.Faculty{
		Name: data.Name,
	}).Error; err != nil {
		return nil, err
	}
	return faculty, nil
}

func (r *FacultyRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return r.WithContext(ctx).Unscoped().Delete(&entity.Faculty{ID: id}).Error
}
