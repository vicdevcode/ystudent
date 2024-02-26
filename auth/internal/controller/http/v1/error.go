package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type messageResponse struct {
	Message string `json:"message"`
}

func errorResponse(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, messageResponse{message})
}

func badRequest(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, messageResponse{message})
}

func internalServerError(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, messageResponse{message})
}
