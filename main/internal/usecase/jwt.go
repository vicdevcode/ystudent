package usecase

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type JwtConfig struct {
	AccessTokenSecret  string
	RefreshTokenSecret string
	AccessExpiresAt    time.Duration
	RefreshExpiresAt   time.Duration
}

type JwtUseCase struct {
	JwtConfig
}

func newJwt(c JwtConfig) *JwtUseCase {
	return &JwtUseCase{c}
}

func (j *JwtUseCase) ExtractFromToken(
	requestToken string,
	key string,
	isAccessToken bool,
) (string, error) {
	secret := j.RefreshTokenSecret
	if isAccessToken {
		secret = j.AccessTokenSecret
	}
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("Invalid Token")
	}

	return claims[key].(string), nil
}
