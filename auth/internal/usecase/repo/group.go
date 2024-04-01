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
	if err := r.WithContext(ctx).Preload("Students").Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

func (r *GroupRepo) FindOneByID(ctx context.Context, id uuid.UUID) (*entity.Group, error) {
	var group *entity.Group
	if err := r.WithContext(ctx).Where("id = ?", id).First(&group).Error; err != nil {
		return nil, err
	}
	return group, nil
}

func (r *GroupRepo) FindOneByName(ctx context.Context, name string) (*entity.Group, error) {
	var group *entity.Group
	if err := r.WithContext(ctx).Where("name = ?", name).First(&group).Error; err != nil {
		return nil, err
	}
	return group, nil
}

func (r *GroupRepo) UpdateCurator(
	ctx context.Context,
	data dto.UpdateGroupCurator,
) (*entity.Group, error) {
	group := &entity.Group{
		ID: data.ID,
	}
	if err := r.WithContext(ctx).Model(&group).Update("curator_id", data.CuratorID).Error; err != nil {
		return nil, err
	}
	return group, nil
}
