package repo

import (
	"context"

	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/main/internal/dto"
	"github.com/vicdevcode/ystudent/main/internal/entity"
	"github.com/vicdevcode/ystudent/main/pkg/postgres"
)

type StudentRepo struct {
	*postgres.Postgres
}

func NewStudent(db *postgres.Postgres) *StudentRepo {
	return &StudentRepo{db}
}

func (r *StudentRepo) Create(ctx context.Context, data dto.CreateStudent) (*entity.Student, error) {
	student := &entity.Student{
		GroupID: data.GroupID,
		Leader:  data.Leader,
		User: entity.User{
			Firstname:  data.Firstname,
			Middlename: data.Middlename,
			Surname:    data.Surname,
			Email:      data.Email,
			RoleType:   entity.STUDENT,
		},
	}

	if err := r.WithContext(ctx).Create(student).Error; err != nil {
		return nil, err
	}
	return student, nil
}

func (r *StudentRepo) FindAll(ctx context.Context, page dto.Page) ([]entity.Student, error) {
	var students []entity.Student
	if err := r.WithContext(ctx).Limit(page.Count).Offset((page.Page - 1) * page.Count).Preload("User").Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (r *StudentRepo) FindOneByID(ctx context.Context, id uuid.UUID) (*entity.Student, error) {
	var student *entity.Student
	if err := r.WithContext(ctx).Preload("User").Where("id = ?", id).First(&student).Error; err != nil {
		return nil, err
	}
	return student, nil
}

func (r *StudentRepo) Update(ctx context.Context, data dto.UpdateStudent) (*entity.Student, error) {
	student := &entity.Student{ID: data.ID}
	if err := r.WithContext(ctx).Model(&student).Updates(entity.Student{GroupID: data.GroupID}).Error; err != nil {
		return nil, err
	}
	return student, nil
}

func (r *StudentRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return r.WithContext(ctx).Unscoped().Delete(&entity.Student{ID: id}).Error
}
