package usecase

import (
	"golang.org/x/crypto/bcrypt"
)

type HashUseCase struct{}

func newHash() *HashUseCase {
	return &HashUseCase{}
}

func (h *HashUseCase) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (h *HashUseCase) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
