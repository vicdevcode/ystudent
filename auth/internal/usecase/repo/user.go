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

// Create

func (r *UserRepo) Create(ctx context.Context, data dto.CreateUser) (*entity.User, error) {
	user := &entity.User{
		ID:         data.ID,
		Firstname:  data.Firstname,
		Middlename: data.Middlename,
		Surname:    data.Surname,
		Email:      data.Email,
		Password:   data.Password,
		RoleType:   data.Role,
	}
	if err := r.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Select

func (r *UserRepo) FindAll(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	if err := r.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepo) FindAllByIDs(ctx context.Context, ids []uuid.UUID) ([]entity.User, error) {
	var users []entity.User
	if err := r.WithContext(ctx).Find(&users, ids).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepo) FindOneByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	var user *entity.User
	if err := r.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepo) FindOneByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user *entity.User
	if err := r.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepo) FindOneByRefreshToken(
	ctx context.Context,
	refreshToken string,
) (*entity.User, error) {
	var user *entity.User
	if err := r.WithContext(ctx).Where("refresh_token = ?", refreshToken).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Delete

func (r *UserRepo) Delete(ctx context.Context, id string) error {
	return nil
}

// Update

func (r *UserRepo) UpdateRefreshToken(
	ctx context.Context,
	data dto.UpdateRefreshToken,
) (*entity.User, error) {
	user := &entity.User{
		ID: data.ID,
	}
	if err := r.WithContext(ctx).Model(&user).Update("refresh_token", data.RefreshToken).Error; err != nil {
		return nil, err
	}
	if err := r.WithContext(ctx).Where(user).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
