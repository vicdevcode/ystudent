package usecase

import (
	"context"
	"time"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
)

type StudentUseCase struct {
	userRepo    UserRepo
	studentRepo StudentRepo
	ctxTimeout  time.Duration
}

func newStudent(ur UserRepo, sr StudentRepo, t time.Duration) *StudentUseCase {
	return &StudentUseCase{
		userRepo:    ur,
		studentRepo: sr,
		ctxTimeout:  t,
	}
}

func (uc *StudentUseCase) SignUp(c context.Context, data dto.CreateUserAndStudent) (*entity.Student, error) {
	ctx, cancel := context.WithTimeout(c, uc.ctxTimeout)
	defer cancel()
	user, err := uc.userRepo.Create(ctx, data.CreateUser)
	if err != nil {
		return nil, err
	}
	studentData := dto.CreateStudent{
		UserID:  user.ID.String(),
		GroupID: data.GroupID,
		Leader:  data.Leader,
	}
	student, err := uc.studentRepo.Create(ctx, studentData)
	return student, nil
}
