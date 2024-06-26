package repo

import (
	"context"

	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/main/internal/dto"
	"github.com/vicdevcode/ystudent/main/internal/entity"
	"github.com/vicdevcode/ystudent/main/pkg/postgres"
)

type GroupRepo struct {
	*postgres.Postgres
}

func NewGroup(db *postgres.Postgres) *GroupRepo {
	return &GroupRepo{db}
}

func (r *GroupRepo) Create(ctx context.Context, data dto.CreateGroup) (*entity.Group, error) {
	departmentID, err := uuid.Parse(data.DepartmentID)
	if err != nil {
		return nil, err
	}
	curatorID, err := uuid.Parse(data.CuratorID)
	if err != nil {
		return nil, err
	}
	group := &entity.Group{Name: data.Name, DepartmentID: &departmentID, CuratorID: &curatorID}
	if err := r.WithContext(ctx).Create(group).Error; err != nil {
		return nil, err
	}

	return group, nil
}

func (r *GroupRepo) FindAll(ctx context.Context, page dto.Page) ([]entity.Group, error) {
	var groups []entity.Group
	if err := r.WithContext(ctx).Limit(page.Count).Offset((page.Page - 1) * page.Count).Preload("Students.User").Find(&groups).Error; err != nil {
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

func (r *GroupRepo) Update(ctx context.Context, data dto.UpdateGroup) (*entity.Group, error) {
	group := &entity.Group{ID: data.ID}
	if err := r.WithContext(ctx).Model(&group).Updates(entity.Group{Name: data.Name, DepartmentID: data.DepartmentID, CuratorID: data.CuratorID}).Error; err != nil {
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

func (r *GroupRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return r.WithContext(ctx).Unscoped().Delete(&entity.Group{ID: id}).Error
}
