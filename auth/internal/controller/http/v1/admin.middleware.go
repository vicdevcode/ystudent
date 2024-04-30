package v1

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/vicdevcode/ystudent/auth/internal/entity"
	"github.com/vicdevcode/ystudent/auth/internal/usecase"
)

func adminCheckMiddleware(uj usecase.Jwt, ua usecase.Admin) gin.HandlerFunc {
	return func(c *gin.Context) {
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
		ok, err := uj.IsTokenValid(authorization[1], true)
		if err != nil {
			errorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}
		if !ok {
			unauthorized(c)
			return
		}

		id, err := uj.ExtractFromToken(authorization[1], "id", true)
		if err != nil {
			unauthorized(c)
			return
		}
		uid64, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			unauthorized(c)
			return
		}
		login, err := uj.ExtractFromToken(authorization[1], "email", true)

		admin, err := ua.FindOne(c, entity.Admin{
			ID: uint(uid64),
		})
		if err != nil {
			unauthorized(c)
			return
		}

		if admin.Login != login {
			unauthorized(c)
			return
		}

		c.Next()
	}
}
