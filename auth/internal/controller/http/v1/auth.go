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
	uh usecase.Hash
	uj usecase.Jwt
	l  *slog.Logger
}

func newAuth(handler *gin.RouterGroup, ua usecase.Admin, uh usecase.Hash, uj usecase.Jwt, l *slog.Logger) {
	r := &authRoute{ua, uh, uj, l}
	h := handler.Group("/auth")
	{
		h.POST("/admin", r.signInAdmin)
	}
}

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

	candidate, err := r.ua.FindOne(c, body.Login)
	if err != nil {
		badRequest(c, err.Error())
		return
	}

	ok := r.uh.CheckPasswordHash(body.Password, candidate.Password)
	if !ok {
		badRequest(c, "incorrect password")
		return
	}

	accessToken, err := r.uj.CreateToken(dto.TokenPayload{
		ID:    fmt.Sprintf("%v", candidate.ID),
		Email: candidate.Login,
	}, true)
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	refreshToken, err := r.uj.CreateToken(dto.TokenPayload{
		ID:    fmt.Sprintf("%v", candidate.ID),
		Email: candidate.Login,
	}, false)
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	admin, err := r.ua.UpdateRefreshToken(c.Request.Context(), dto.UpdateRefreshToken{
		ID:           candidate.ID,
		RefreshToken: refreshToken,
	})
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, &signInAdminResponse{
		Admin:        admin,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
