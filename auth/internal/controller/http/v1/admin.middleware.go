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
		token := strings.Split(headerToken[0], " ")[1]
		ok, err := uj.IsTokenValid(token, true)
		if err != nil {
			errorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}
		if !ok {
			unauthorized(c)
			return
		}

		id, err := uj.ExtractFromToken(token, "id", true)
		if err != nil {
			unauthorized(c)
			return
		}
		uid64, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			unauthorized(c)
			return
		}
		login, err := uj.ExtractFromToken(token, "email", true)

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
