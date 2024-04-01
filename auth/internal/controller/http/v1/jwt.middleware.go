package v1

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/vicdevcode/ystudent/auth/internal/usecase"
)

func jwtCheckMiddleware(uj usecase.Jwt) gin.HandlerFunc {
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
		c.Next()
	}
}
