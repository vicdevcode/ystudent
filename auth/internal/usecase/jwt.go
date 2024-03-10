package usecase

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
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

type CustomAccessTokenClaims struct {
	dto.AccessTokenPayload
	jwt.RegisteredClaims
}

func (j *JwtUseCase) CreateAccessToken(payload dto.AccessTokenPayload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomAccessTokenClaims{
		payload,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.AccessExpiresAt)),
		},
	})
	t, err := token.SignedString([]byte(j.AccessTokenSecret))
	if err != nil {
		return "", err
	}
	return t, err
}

type CustomRefreshTokenClaims struct {
	dto.RefreshTokenPayload
	jwt.RegisteredClaims
}

func (j *JwtUseCase) CreateRefreshToken(payload dto.RefreshTokenPayload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomRefreshTokenClaims{
		payload,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.RefreshExpiresAt)),
		},
	})
	t, err := token.SignedString([]byte(j.RefreshTokenSecret))
	if err != nil {
		return "", err
	}
	return t, err
}

func (j *JwtUseCase) CreateTokens(atPayload dto.AccessTokenPayload, rtPayload dto.RefreshTokenPayload) (*dto.Tokens, error) {
	accessToken, err := j.CreateAccessToken(atPayload)
	if err != nil {
		return nil, err
	}
	refreshToken, err := j.CreateRefreshToken(rtPayload)
	if err != nil {
		return nil, err
	}
	return &dto.Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (j *JwtUseCase) IsTokenValid(requestToken string, isAccessToken bool) (bool, error) {
	secret := j.RefreshTokenSecret
	if isAccessToken {
		secret = j.AccessTokenSecret
	}
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (j *JwtUseCase) ExtractFromToken(requestToken string, key string, isAccessToken bool) (string, error) {
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
