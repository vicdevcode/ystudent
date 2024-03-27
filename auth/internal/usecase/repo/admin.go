package repo

import (
	"context"

	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
	"github.com/vicdevcode/ystudent/auth/pkg/postgres"
)

type AdminRepo struct {
	*postgres.Postgres
}

func NewAdmin(db *postgres.Postgres) *AdminRepo {
	return &AdminRepo{db}
}

func (r *AdminRepo) FindOneByLogin(ctx context.Context, login string) (*entity.Admin, error) {
	var admin *entity.Admin
	if err := r.WithContext(ctx).Where("login = ?", login).First(&admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *AdminRepo) FindOneByRefreshToken(
	ctx context.Context,
	refreshToken string,
) (*entity.Admin, error) {
	var admin *entity.Admin
	if err := r.WithContext(ctx).Where("refresh_token = ?", refreshToken).First(&admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *AdminRepo) FindOneByID(ctx context.Context, id uuid.UUID) (*entity.Admin, error) {
	var admin *entity.Admin
	if err := r.WithContext(ctx).Where("id = ?", id).First(&admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *AdminRepo) UpdateRefreshToken(
	ctx context.Context,
	data dto.UpdateRefreshToken,
) (*entity.Admin, error) {
	admin := &entity.Admin{
		ID: data.ID,
	}
	if err := r.WithContext(ctx).Model(&admin).Update("refresh_token", data.RefreshToken).Error; err != nil {
		return nil, err
	}
	if err := r.WithContext(ctx).Where("id = ?", data.ID).First(&admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}
