package repo

import (
	"context"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
	"github.com/vicdevcode/ystudent/auth/pkg/postgres"
)

type UserRepo struct {
	*postgres.Postgres
}

func NewUser(db *postgres.Postgres) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) Create(ctx context.Context, data dto.CreateUser) (*entity.User, error) {
	return &entity.User{}, nil
}

func (r *UserRepo) FindAll(ctx context.Context) ([]entity.User, error) {
	return []entity.User{}, nil
}

func (r *UserRepo) Delete(ctx context.Context, id string) error {
	return nil
}
