package usecase

import (
	"github.com/vicdevcode/ystudent/auth/internal/usecase/repo"
	"github.com/vicdevcode/ystudent/auth/pkg/config"
	"github.com/vicdevcode/ystudent/auth/pkg/postgres"
)

type UseCases struct {
	UserUseCase User
	HashUseCase Hash
	JwtUseCase  Jwt
}

func New(cfg *config.Config, db *postgres.Postgres) UseCases {
	t := cfg.ContextTimeout
	return UseCases{
		UserUseCase: newUser(repo.NewUser(db), t),
		JwtUseCase:  newJwt(JwtConfig(cfg.JWT)),
		HashUseCase: newHash(),
	}
}
