package repo

import (
	"context"

	"github.com/google/uuid"

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
	userId, err := uuid.Parse(data.UserID)
	if err != nil {
		return nil, err
	}

	groupId, err := uuid.Parse(data.GroupID)
	if err != nil {
		return nil, err
	}

	student := &entity.Student{
		Leader:  data.Leader,
		UserID:  userId,
		GroupID: groupId,
	}

	if err := r.WithContext(ctx).Create(student).Error; err != nil {
		return nil, err
	}
	return student, nil
}
