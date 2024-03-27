package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
)

type TeacherUseCase struct {
	teacherRepo TeacherRepo
	userRepo    UserRepo
	ctxTimeout  time.Duration
}

func newTeacher(tr TeacherRepo, ur UserRepo, t time.Duration) *TeacherUseCase {
	return &TeacherUseCase{
		teacherRepo: tr,
		userRepo:    ur,
		ctxTimeout:  t,
	}
}

func (uc *TeacherUseCase) Create(
	c context.Context,
	data dto.CreateTeacher,
) (*entity.Teacher, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	teacher, err := uc.teacherRepo.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

func (uc *TeacherUseCase) FindAll(c context.Context) ([]entity.User, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()

	students, err := uc.teacherRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var ids []uuid.UUID
	for _, student := range students {
		ids = append(ids, student.UserID)
	}

	users, err := uc.userRepo.FindAllByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	return users, nil
}
