package repo

import (
	"context"

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

func (r *StudentRepo) FindAll(ctx context.Context) ([]entity.Student, error) {
	var students []entity.Student
	if err := r.WithContext(ctx).Preload("User").Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}
