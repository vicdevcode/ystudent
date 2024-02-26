package usecase

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
)

type JwtConfig struct {
	Secret           string
	AccessExpiresAt  time.Duration
	RefreshExpiresAt time.Duration
}

type JwtUseCase struct {
	JwtConfig
}

func newJwt(c JwtConfig) *JwtUseCase {
	return &JwtUseCase{c}
}

type CustomClaims struct {
	dto.TokenPayload
	jwt.RegisteredClaims
}

func (j *JwtUseCase) CreateToken(payload dto.TokenPayload, isAccessToken bool) (string, error) {
	var expiresAt time.Duration
	if isAccessToken {
		expiresAt = j.AccessExpiresAt
	} else {
		expiresAt = j.RefreshExpiresAt
	}
	claims := CustomClaims{
		payload,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresAt)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func (j *JwtUseCase) IsAuthorized(requestToken string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.Secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (j *JwtUseCase) ExtractIDFromToken(requestToken string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.Secret), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("Invalid Token")
	}

	return claims["id"].(string), nil
}
