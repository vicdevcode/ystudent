package repo

import (
	"context"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
	"github.com/vicdevcode/ystudent/auth/pkg/postgres"
)

type StudentRepo struct {
	*postgres.Postgres
}

func NewStudent(db *postgres.Postgres) *StudentRepo {
	return &StudentRepo{db}
}

func (r *StudentRepo) Create(ctx context.Context, data dto.CreateStudent) (*entity.Student, error) {
	student := &entity.Student{
		Leader:  data.Leader,
		UserID:  data.UserID,
		GroupID: data.GroupID,
	}

	if err := r.WithContext(ctx).Create(student).Error; err != nil {
		return nil, err
	}
	return student, nil
}

func (r *StudentRepo) FindAll(ctx context.Context) ([]entity.Student, error) {
	var students []entity.Student
	if err := r.WithContext(ctx).Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}
