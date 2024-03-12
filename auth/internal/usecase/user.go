package usecase

import (
	"context"
	"errors"
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

func (uc *UserUseCase) Create(c context.Context, data dto.CreateUser) (*entity.User, error) {
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

func (uc *UserUseCase) FindOne(c context.Context, data entity.User) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	var user *entity.User
	var err error

	if len(data.Email) != 0 {
		user, err = uc.repo.FindOneByEmail(ctx, data.Email)
	} else if uuid.Nil != data.ID {
		user, err = uc.repo.FindOneByID(ctx, data.ID)
	} else if len(data.RefreshToken) != 0 {
		user, err = uc.repo.FindOneByRefreshToken(ctx, data.RefreshToken)
	} else {
		// TODO: Придумать текст ошибкы
		err = errors.New("record not found")
	}

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserUseCase) UpdateRefreshToken(c context.Context, data dto.UpdateRefreshToken) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	user, err := uc.repo.UpdateRefreshToken(ctx, data)
	if err != nil {
		return nil, err
	}
	return user, err
}
