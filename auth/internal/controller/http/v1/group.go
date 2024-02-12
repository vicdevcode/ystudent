package v1

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
	"github.com/vicdevcode/ystudent/auth/internal/usecase"
)

type groupRoute struct {
	u usecase.Group
	l *slog.Logger
}

func NewGroup(handler *gin.RouterGroup, u usecase.Group, l *slog.Logger) {
	r := &groupRoute{u, l}
	h := handler.Group("/group")
	{
		h.POST("/", r.create)
		h.GET("/", r.findAll)
	}
}

type createGroupRequest struct {
	dto.CreateGroup
}

type createGroupResponse struct {
	*entity.Group
	dto.CUD
}

func (r *groupRoute) create(c *gin.Context) {
	var body createGroupRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		errorJSONResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	group, err := r.u.Create(c.Request.Context(), body.CreateGroup)
	if err != nil {
		r.l.Error("sign up", slog.Any(err.Error(), err))
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, createGroupResponse{Group: group})
}

type findAllGroupResponse struct {
	Groups []entity.Group `json:"groups"`
}

func (r *groupRoute) findAll(c *gin.Context) {
	groups, err := r.u.FindAll(c.Request.Context())
	if err != nil {
		r.l.Error("sign up", slog.Any(err.Error(), err))
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, findAllGroupResponse{Groups: groups})
}
