package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

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

func (uc *AdminUseCase) FindOne(c context.Context, data entity.Admin) (*entity.Admin, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	if len(data.Login) != 0 {
		return uc.repo.FindOneByLogin(ctx, data.Login)
	} else if uuid.Nil != data.ID {
		return uc.repo.FindOneByID(ctx, data.ID)
	} else if len(data.RefreshToken) != 0 {
		return uc.repo.FindOneByRefreshToken(ctx, data.RefreshToken)
	}

	return nil, errors.New("record not found")
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
