package repo

import (
	"context"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
	"github.com/vicdevcode/ystudent/auth/pkg/sqlite"
)

type FacultyRepo struct {
	*sqlite.SQLite
}

func NewFaculty(db *sqlite.SQLite) *FacultyRepo {
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
