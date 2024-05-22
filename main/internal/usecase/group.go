package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/main/internal/dto"
	"github.com/vicdevcode/ystudent/main/internal/entity"
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

func (uc *GroupUseCase) FindOne(c context.Context, data entity.Group) (*entity.Group, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	if len(data.Name) != 0 {
		return uc.repo.FindOneByName(ctx, data.Name)
	} else if uuid.Nil != data.ID {
		return uc.repo.FindOneByID(ctx, data.ID)
	}

	return nil, errors.New("record not found")
}

func (uc *GroupUseCase) Update(
	c context.Context,
	data dto.UpdateGroup,
) (*entity.Group, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	return uc.repo.Update(ctx, data)
}

func (uc *GroupUseCase) UpdateCurator(
	c context.Context,
	data dto.UpdateGroupCurator,
) (*entity.Group, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	group, err := uc.repo.UpdateCurator(ctx, data)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (uc *GroupUseCase) Delete(
	c context.Context,
	id uuid.UUID,
) error {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	return uc.repo.Delete(ctx, id)
}
