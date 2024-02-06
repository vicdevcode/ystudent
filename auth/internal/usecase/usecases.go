package usecase

import (
	"github.com/vicdevcode/ystudent/auth/internal/usecase/repo"
	"github.com/vicdevcode/ystudent/auth/pkg/config"
	"github.com/vicdevcode/ystudent/auth/pkg/postgres"
)

type UseCases struct {
	UserUseCase User
}

func New(cfg *config.Config, db *postgres.Postgres) UseCases {
	return UseCases{
		UserUseCase: NewUser(repo.NewUser(db), cfg.ContextTimeout),
	}
}
