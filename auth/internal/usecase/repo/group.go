package repo

import (
	"context"

	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
	"github.com/vicdevcode/ystudent/auth/pkg/postgres"
)

type GroupRepo struct {
	*postgres.Postgres
}

func NewGroup(db *postgres.Postgres) *GroupRepo {
	return &GroupRepo{db}
}

func (r *GroupRepo) Create(ctx context.Context, data dto.CreateGroup) (*entity.Group, error) {
	facultyId, err := uuid.Parse(data.FacultyID)
	if err != nil {
		return nil, err
	}
	group := &entity.Group{Name: data.Name, FacultyID: &facultyId}
	if err := r.WithContext(ctx).Create(group).Error; err != nil {
		return nil, err
	}

	return group, nil
}

func (r *GroupRepo) FindAll(ctx context.Context) ([]entity.Group, error) {
	var groups []entity.Group
	if err := r.WithContext(ctx).Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}
