package repo

import (
	"context"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
	"github.com/vicdevcode/ystudent/auth/pkg/postgres"
)

type TeacherRepo struct {
	*postgres.Postgres
}

func NewTeacher(db *postgres.Postgres) *TeacherRepo {
	return &TeacherRepo{db}
}

func (r *TeacherRepo) Create(ctx context.Context, data dto.CreateTeacher) (*entity.Teacher, error) {
	teacher := &entity.Teacher{
		UserID: data.UserID,
	}

	if err := r.WithContext(ctx).Create(teacher).Error; err != nil {
		return nil, err
	}

	return teacher, nil
}
