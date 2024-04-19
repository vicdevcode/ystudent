package v1

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
	"github.com/vicdevcode/ystudent/auth/internal/usecase"
)

type authRoute struct {
	ua usecase.Admin
	uu usecase.User
	uh usecase.Hash
	uj usecase.Jwt
	l  *slog.Logger
}

func newAuth(
	public *gin.RouterGroup,
	private *gin.RouterGroup,
	ua usecase.Admin,
	uu usecase.User,
	uh usecase.Hash,
	uj usecase.Jwt,
	l *slog.Logger,
) {
	r := &authRoute{ua, uu, uh, uj, l}
	{
		public.GET("/auth/logout", r.logout)
		public.POST("/auth/", r.signIn)
		public.POST("/auth/admin", r.signInAdmin)
		public.GET("/auth/refresh-token", r.refreshTokens)
		private.GET("/auth/check", r.check)
	}
}

func (r *authRoute) check(c *gin.Context) {
	c.JSON(200, gin.H{"message": "successfully logged in"})
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

	var role string
	if candidate.Student.ID != 0 {
		role = "student"
	} else if candidate.Teacher.ID != 0 {
		role = "teacher"
	} else {
		internalServerError(c, "internal server error")
		return
	}

	tokens, err := r.uj.CreateTokens(dto.AccessTokenPayload{
		ID:    fmt.Sprintf("%v", candidate.ID),
		Email: candidate.Email,
		Role:  role,
	}, dto.RefreshTokenPayload{
		ID:    fmt.Sprintf("%v", candidate.ID),
		Email: candidate.Email,
		Role:  role,
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

	c.SetCookie("refresh_token", tokens.RefreshToken, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, &signInResponse{
		User:         dto.UserResponse{User: user},
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})
}

// signInAdmin

type signInAdminRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type signInAdminResponse struct {
	Admin        *entity.Admin `json:"admin"`
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
}

func (r *authRoute) signInAdmin(c *gin.Context) {
	var body signInAdminRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		badRequest(c, err.Error())
		return
	}

	candidate, err := r.ua.FindOne(c, entity.Admin{
		Login: body.Login,
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
		Email: candidate.Login,
		Role:  "admin",
	}, dto.RefreshTokenPayload{
		ID:    fmt.Sprintf("%v", candidate.ID),
		Email: candidate.Login,
		Role:  "admin",
	})
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	admin, err := r.ua.UpdateRefreshToken(c.Request.Context(), dto.UpdateRefreshToken{
		ID:           candidate.ID,
		RefreshToken: tokens.RefreshToken,
	})
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.SetCookie("refresh_token", tokens.RefreshToken, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, &signInAdminResponse{
		Admin:        admin,
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

	role, err := r.uj.ExtractFromToken(refreshToken, "role", false)
	if err != nil {
		unauthorized(c)
		return
	}

	if role == "admin" {
		admin, err := r.ua.FindOne(c, entity.Admin{
			RefreshToken: refreshToken,
		})
		if err != nil {
			unauthorized(c)
			return
		}

		tokens, err := r.uj.CreateTokens(dto.AccessTokenPayload{
			ID:    fmt.Sprintf("%v", admin.ID),
			Email: admin.Login,
			Role:  role,
		}, dto.RefreshTokenPayload{
			ID:    fmt.Sprintf("%v", admin.ID),
			Email: admin.Login,
			Role:  role,
		})

		_, err = r.ua.UpdateRefreshToken(
			c,
			dto.UpdateRefreshToken{ID: admin.ID, RefreshToken: tokens.RefreshToken},
		)

		c.SetCookie("refresh_token", tokens.RefreshToken, 3600, "/", "localhost", false, true)

		if err != nil {
			return
		}

		c.JSON(http.StatusOK, refreshTokensResponse(tokens))
	}
}

type logoutResponse struct {
	Message string `json:"message"`
}

func (r *authRoute) logout(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		badRequest(c, "cookie not set")
		return
	}
	if refreshToken == "" {
		badRequest(c, "refresh token is empty")
		return
	}

	role, err := r.uj.ExtractFromToken(refreshToken, "role", false)
	if err != nil {
		unauthorized(c)
		return
	}

	if role == "admin" {
		admin, err := r.ua.FindOne(c, entity.Admin{
			RefreshToken: refreshToken,
		})
		if err != nil {
			unauthorized(c)
			return
		}

		_, err = r.ua.UpdateRefreshToken(c, dto.UpdateRefreshToken{ID: admin.ID, RefreshToken: ""})
		if err != nil {
			internalServerError(c, "can not delete refresh token in db")
			return
		}
	} else if role == "student" || role == "teacher" {
		user, err := r.uu.FindOne(c, entity.User{
			RefreshToken: refreshToken,
		})
		if err != nil {
			unauthorized(c)
			return
		}

		_, err = r.uu.UpdateRefreshToken(c, dto.UpdateRefreshToken{ID: user.ID, RefreshToken: ""})
		if err != nil {
			internalServerError(c, "can not delete refresh token in db")
			return
		}
	} else {
		internalServerError(c, fmt.Sprintf("this role %s is not exists", role))
		return
	}
	c.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, logoutResponse{Message: "successfully logged out"})
}
