package v1

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/vicdevcode/ystudent/main/internal/entity"
	"github.com/vicdevcode/ystudent/main/internal/usecase"
)

func userAuthCheckMiddleware(uj usecase.Jwt) gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToken := c.Request.Header["Authorization"]
		role, err := getRole(headerToken, uj)
		if err != nil {
			unauthorized(c)
			return
		}

		if isValid(entity.UserRole(role)) {
			unauthorized(c)
			return
		}

		c.Next()
	}
}

func moderatorAuthCheckMiddleware(uj usecase.Jwt) gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToken := c.Request.Header["Authorization"]
		role, err := getRole(headerToken, uj)
		if err != nil {
			unauthorized(c)
			return
		}

		if role != "MODERATOR" && role != "ADMIN" {
			unauthorized(c)
			return
		}

		c.Next()
	}
}

func adminAuthCheckMiddleware(uj usecase.Jwt) gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToken := c.Request.Header["Authorization"]
		role, err := getRole(headerToken, uj)
		if err != nil {
			unauthorized(c)
			return
		}

		if role != "ADMIN" {
			unauthorized(c)
			return
		}

		c.Next()
	}
}

func isValid(role entity.UserRole) bool {
	switch role {
	case entity.TEACHER, entity.ADMIN, entity.STUDENT, entity.MODERATOR, entity.EMPLOYEE:
		return true
	default:
		return false
	}
}

func getRole(headerToken []string, uj usecase.Jwt) (string, error) {
	var role string
	if len(headerToken) == 0 {
		return "", errors.New("no authorization token")
	}
	authorization := strings.Split(headerToken[0], " ")
	if len(authorization) != 2 {
		return "", errors.New("bad authorization token")
	}

	role, err := uj.ExtractFromToken(authorization[1], "role", true)
	if err != nil {
		return "", errors.New("can not extract from token")
	}

	return role, nil
}
