package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/main/internal/dto"
	"github.com/vicdevcode/ystudent/main/internal/entity"
)

type TeacherUseCase struct {
	teacherRepo TeacherRepo
	userRepo    UserRepo
	groupRepo   GroupRepo
	ctxTimeout  time.Duration
}

func newTeacher(tr TeacherRepo, ur UserRepo, gr GroupRepo, t time.Duration) *TeacherUseCase {
	return &TeacherUseCase{
		teacherRepo: tr,
		userRepo:    ur,
		groupRepo:   gr,
		ctxTimeout:  t,
	}
}

func (uc *TeacherUseCase) Create(
	c context.Context,
	data dto.CreateTeacher,
) (*entity.Teacher, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	return uc.teacherRepo.Create(ctx, data)
}

func (uc *TeacherUseCase) FindAll(c context.Context, page dto.Page) ([]entity.Teacher, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	return uc.teacherRepo.FindAll(ctx, page)
}

func (uc *TeacherUseCase) FindOne(
	c context.Context,
	data entity.Teacher,
) (*entity.Teacher, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	if uuid.Nil != data.ID {
		return uc.teacherRepo.FindOneByID(ctx, data.ID)
	}
	return nil, errors.New("record not found")
}

func (uc *TeacherUseCase) Delete(
	c context.Context,
	id uuid.UUID,
) error {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	teacher, err := uc.teacherRepo.FindOneByID(ctx, id)
	if err != nil {
		return err
	}

	if err := uc.teacherRepo.Delete(ctx, id); err != nil {
		return err
	}

	if err := uc.userRepo.Delete(ctx, teacher.User.ID); err != nil {
		return err
	}

	return nil
}

func (uc *TeacherUseCase) AddGroup(
	c context.Context,
	data dto.AddGroupToTeacher,
) (*entity.Teacher, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	return uc.teacherRepo.AddGroup(ctx, data)
}

func (uc *TeacherUseCase) DeleteGroup(
	c context.Context,
	data dto.DeleteGroupFromTeacher,
) (*entity.Teacher, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	group, err := uc.groupRepo.FindOneByID(ctx, data.GroupID)
	if err != nil {
		return nil, err
	}

	teacher, err := uc.teacherRepo.FindOneByID(ctx, data.TeacherID)
	if err != nil {
		return nil, err
	}

	return uc.teacherRepo.DeleteGroup(ctx, teacher, group)
}
