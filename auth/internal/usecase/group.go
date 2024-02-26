package usecase

import (
	"context"
	"time"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
)

type GroupUseCase struct {
	repo       GroupRepo
	ctxTimeout time.Duration
}

func newGroup(r GroupRepo, t time.Duration) *GroupUseCase {
	return &GroupUseCase{
		repo:       r,
		ctxTimeout: t,
	}
}

func (uc *GroupUseCase) Create(c context.Context, data dto.CreateGroup) (*entity.Group, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	group, err := uc.repo.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (uc *GroupUseCase) FindAll(c context.Context) ([]entity.Group, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	groups, err := uc.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return groups, nil
}
