package usecase

import (
	"context"
	"time"

	"github.com/vicdevcode/ystudent/auth/internal/entity"
)

type UserUseCase struct {
	repo       UserRepo
	ctxTimeout time.Duration
}

func NewUser(r UserRepo, t time.Duration) *UserUseCase {
	return &UserUseCase{
		repo:       r,
		ctxTimeout: t,
	}
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
