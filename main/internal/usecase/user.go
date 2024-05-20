package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/main/internal/dto"
	"github.com/vicdevcode/ystudent/main/internal/entity"
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

func (uc *UserUseCase) Create(c context.Context, data dto.CreateUser) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	user, err := uc.repo.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserUseCase) FindAll(c context.Context, role entity.UserRole) ([]entity.User, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	users, err := uc.repo.FindAll(ctx, role)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uc *UserUseCase) FindOne(
	c context.Context,
	role entity.UserRole,
	data entity.User,
) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	if len(data.Email) != 0 {
		return uc.repo.FindOneByEmail(ctx, role, data.Email)
	} else if uuid.Nil != data.ID {
		return uc.repo.FindOneByID(ctx, role, data.ID)
	}
	return nil, errors.New("record not found")
}
