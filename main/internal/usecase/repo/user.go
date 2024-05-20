package repo

import (
	"context"

	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/main/internal/dto"
	"github.com/vicdevcode/ystudent/main/internal/entity"
	"github.com/vicdevcode/ystudent/main/pkg/postgres"
)

type UserRepo struct {
	*postgres.Postgres
}

func NewUser(db *postgres.Postgres) *UserRepo {
	return &UserRepo{db}
}

// Create

func (r *UserRepo) Create(ctx context.Context, data dto.CreateUser) (*entity.User, error) {
	user := &entity.User{
		Firstname:  data.Firstname,
		Middlename: data.Middlename,
		Surname:    data.Surname,
		Email:      data.Email,
		RoleType:   data.RoleType,
	}
	if err := r.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Select

func (r *UserRepo) FindAll(ctx context.Context, roleType entity.UserRole) ([]entity.User, error) {
	var users []entity.User
	if err := r.WithContext(ctx).Where("role_type = ?", roleType).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepo) FindOneByID(
	ctx context.Context,
	roleType entity.UserRole,
	id uuid.UUID,
) (*entity.User, error) {
	var user *entity.User
	if err := r.WithContext(ctx).Where("role_type = ? AND id = ?", roleType, id).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepo) FindOneByEmail(
	ctx context.Context,
	roleType entity.UserRole,
	email string,
) (*entity.User, error) {
	var user *entity.User
	if err := r.WithContext(ctx).Where("role_type = ? AND email = ?", roleType, email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Delete

func (r *UserRepo) Delete(ctx context.Context, id string) error {
	return nil
}

// Update
