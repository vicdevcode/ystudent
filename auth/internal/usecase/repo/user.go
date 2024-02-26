package repo

import (
	"context"

	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
	"github.com/vicdevcode/ystudent/auth/pkg/postgres"
)

type UserRepo struct {
	*postgres.Postgres
}

func NewUser(db *postgres.Postgres) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) Create(ctx context.Context, data dto.CreateUser) (*entity.User, error) {
	user := &entity.User{
		Firstname:  data.Firstname,
		Middlename: data.Middlename,
		Surname:    data.Surname,
		Email:      data.Email,
		Password:   data.Password,
	}
	if err := r.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepo) FindOne(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	var user *entity.User

	if err := r.WithContext(ctx).First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepo) FindAll(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	if err := r.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepo) Delete(ctx context.Context, id string) error {
	return nil
}
