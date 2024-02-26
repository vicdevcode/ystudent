package usecase

import (
	"context"
	"time"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
)

type AdminUseCase struct {
	repo       AdminRepo
	ctxTimeout time.Duration
}

func newAdmin(r AdminRepo, t time.Duration) *AdminUseCase {
	return &AdminUseCase{
		repo:       r,
		ctxTimeout: t,
	}
}

func (uc *AdminUseCase) FindOne(c context.Context, login string) (*entity.Admin, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	admin, err := uc.repo.FindOne(ctx, login)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (uc *AdminUseCase) UpdateRefreshToken(c context.Context, data dto.UpdateRefreshToken) (*entity.Admin, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	admin, err := uc.repo.UpdateRefreshToken(ctx, data)
	if err != nil {
		return nil, err
	}
	return admin, err
}
