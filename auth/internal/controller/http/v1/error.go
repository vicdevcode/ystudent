package v1

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Message string `json:"message"`
}

type ErrorJSONResponse struct {
	Message interface{} `json:"message"`
}

func errorResponse(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, ErrorResponse{message})
}

func errorJSONResponse(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, ErrorJSONResponse{message})
}
