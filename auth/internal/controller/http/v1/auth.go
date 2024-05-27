package v1

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
	"github.com/vicdevcode/ystudent/auth/internal/usecase"
)

type authRoute struct {
	uu usecase.User
	uh usecase.Hash
	uj usecase.Jwt
	l  *slog.Logger
}

func newAuth(
	handler *gin.RouterGroup,
	uu usecase.User,
	uh usecase.Hash,
	uj usecase.Jwt,
	l *slog.Logger,
) {
	r := &authRoute{uu, uh, uj, l}
	h := handler.Group("/auth")
	{
		h.GET("/check", r.check)
		h.POST("/", r.signIn)
		h.GET("/logout", r.logout)
		h.GET("/refresh-token", r.refreshTokens)
	}
}

func (r *authRoute) check(c *gin.Context) {
	headerToken := c.Request.Header["Authorization"]
	if len(headerToken) == 0 {
		unauthorized(c)
		return
	}
	authorization := strings.Split(headerToken[0], " ")
	if len(authorization) != 2 {
		unauthorized(c)
		return
	}
	ok, err := r.uj.IsTokenValid(authorization[1], true)
	if err != nil {
		errorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	if !ok {
		unauthorized(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"active": true})
}

// sign in user

type signInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signInResponse struct {
	User         dto.UserResponse `json:"user"`
	AccessToken  string           `json:"access_token"`
	RefreshToken string           `json:"refresh_token"`
}

func (r *authRoute) signIn(c *gin.Context) {
	var body signInRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		badRequest(c, err.Error())
		return
	}

	candidate, err := r.uu.FindOne(c, entity.User{
		Email: body.Email,
	})
	if err != nil {
		badRequest(c, err.Error())
		return
	}

	ok := r.uh.CheckPasswordHash(body.Password, candidate.Password)
	if !ok {
		badRequest(c, "incorrect password")
		return
	}

	tokens, err := r.uj.CreateTokens(dto.AccessTokenPayload{
		ID:    fmt.Sprintf("%v", candidate.ID),
		Email: candidate.Email,
		Role:  string(candidate.RoleType),
	}, dto.RefreshTokenPayload{
		ID:    fmt.Sprintf("%v", candidate.ID),
		Email: candidate.Email,
		Role:  string(candidate.RoleType),
	})
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	user, err := r.uu.UpdateRefreshToken(c.Request.Context(), dto.UpdateRefreshToken{
		ID:           candidate.ID,
		RefreshToken: tokens.RefreshToken,
	})
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.SetCookie("refresh_token", tokens.RefreshToken, 3600, "/", "y-student.ru", false, true)

	c.JSON(http.StatusOK, &signInResponse{
		User:         dto.UserResponse{User: user},
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})
}

// refresh tokens

type refreshTokensResponse *dto.Tokens

func (r *authRoute) refreshTokens(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		badRequest(c, err.Error())
		return
	}
	if refreshToken == "" {
		badRequest(c, "refresh token is empty")
		return
	}

	user, err := r.uu.FindOne(c, entity.User{
		RefreshToken: refreshToken,
	})
	if err != nil {
		unauthorized(c)
		return
	}

	tokens, err := r.uj.CreateTokens(dto.AccessTokenPayload{
		ID:    fmt.Sprintf("%v", user.ID),
		Email: user.Email,
		Role:  string(user.RoleType),
	}, dto.RefreshTokenPayload{
		ID:    fmt.Sprintf("%v", user.ID),
		Email: user.Email,
		Role:  string(user.RoleType),
	})

	_, err = r.uu.UpdateRefreshToken(
		c,
		dto.UpdateRefreshToken{ID: user.ID, RefreshToken: tokens.RefreshToken},
	)

	c.SetCookie("refresh_token", tokens.RefreshToken, 3600, "/", "y-student.ru", false, true)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, refreshTokensResponse(tokens))
}

type logoutResponse struct {
	Message string `json:"message"`
}

func (r *authRoute) logout(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil || refreshToken == "" {
		c.JSON(http.StatusOK, logoutResponse{Message: "user already logged out"})
		return
	}
	user, err := r.uu.FindOne(c, entity.User{
		RefreshToken: refreshToken,
	})
	if err != nil {
		unauthorized(c)
		return
	}

	_, err = r.uu.UpdateRefreshToken(c, dto.UpdateRefreshToken{ID: user.ID, RefreshToken: ""})
	if err != nil {
		c.SetCookie("refresh_token", "", -1, "/", "y-student.ru", false, true)
		internalServerError(c, "can not delete refresh token in db")
		return
	}
	c.SetCookie("refresh_token", "", -1, "/", "y-student.ru", false, true)
	c.JSON(http.StatusOK, logoutResponse{Message: "successfully logged out"})
}
