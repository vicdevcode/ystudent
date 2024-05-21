package repo

import (
	"context"

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

func (r *TeacherRepo) FindAll(ctx context.Context) ([]entity.Teacher, error) {
	var teachers []entity.Teacher
	if err := r.WithContext(ctx).Preload("User").Find(&teachers).Error; err != nil {
		return nil, err
	}
	return teachers, nil
}