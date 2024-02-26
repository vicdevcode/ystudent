package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
)

type UserUseCase struct {
	repo       UserRepo
	ctxTimeout time.Duration
}

func newUser(r UserRepo, t time.Duration) *UserUseCase {
	return &UserUseCase{
		repo:       r,
		ctxTimeout: t,
	}
}

func (uc *UserUseCase) SignUp(c context.Context, data dto.CreateUser) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	user, err := uc.repo.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserUseCase) FindAll(c context.Context) ([]entity.User, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	users, err := uc.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uc *UserUseCase) FindOne(c context.Context, id uuid.UUID) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	user, err := uc.repo.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
