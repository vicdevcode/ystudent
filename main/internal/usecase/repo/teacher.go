package repo

import (
	"context"

	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/main/internal/dto"
	"github.com/vicdevcode/ystudent/main/internal/entity"
	"github.com/vicdevcode/ystudent/main/pkg/postgres"
)

type TeacherRepo struct {
	*postgres.Postgres
}

func NewTeacher(db *postgres.Postgres) *TeacherRepo {
	return &TeacherRepo{db}
}

func (r *TeacherRepo) Create(ctx context.Context, data dto.CreateTeacher) (*entity.Teacher, error) {
	teacher := &entity.Teacher{
		User: entity.User{
			Email:      data.Email,
			Firstname:  data.Firstname,
			Middlename: data.Middlename,
			Surname:    data.Surname,
			RoleType:   entity.TEACHER,
		},
	}

	if err := r.WithContext(ctx).Create(teacher).Error; err != nil {
		return nil, err
	}

	return teacher, nil
}

func (r *TeacherRepo) FindAll(ctx context.Context, page dto.Page) ([]entity.Teacher, error) {
	var teachers []entity.Teacher
	if err := r.WithContext(ctx).Limit(page.Count).Offset((page.Page - 1) * page.Count).Preload("User").Preload("Groups").Find(&teachers).Error; err != nil {
		return nil, err
	}
	return teachers, nil
}

func (r *TeacherRepo) FindOneByID(ctx context.Context, id uuid.UUID) (*entity.Teacher, error) {
	var teacher *entity.Teacher
	if err := r.WithContext(ctx).Preload("User").Where("id = ?", id).First(&teacher).Error; err != nil {
		return nil, err
	}
	return teacher, nil
}

func (r *TeacherRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return r.WithContext(ctx).Unscoped().Delete(&entity.Teacher{ID: id}).Error
}

func (r *TeacherRepo) AddGroup(
	ctx context.Context,
	data dto.AddGroupToTeacher,
) (*entity.Teacher, error) {
	teacher := &entity.Teacher{
		ID: data.TeacherID,
	}
	if err := r.WithContext(ctx).Model(&teacher).Association("Groups").Append(&entity.Teacher{
		ID: data.GroupID,
	}); err != nil {
		return nil, err
	}

	var err error

	teacher, err = r.FindOneByID(ctx, data.TeacherID)
	if err != nil {
		return nil, err
	}

	return teacher, nil
}

func (r *TeacherRepo) DeleteGroup(
	ctx context.Context,
	teacher *entity.Teacher,
	group *entity.Group,
) (*entity.Teacher, error) {
	if err := r.WithContext(ctx).Model(&teacher).Association("Groups").Delete(&group); err != nil {
		return nil, err
	}

	return teacher, nil
}
