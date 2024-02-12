package v1

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/entity"
	"github.com/vicdevcode/ystudent/auth/internal/usecase"
)

type facultyRoute struct {
	u usecase.Faculty
	l *slog.Logger
}

func NewFaculty(handler *gin.RouterGroup, u usecase.Faculty, l *slog.Logger) {
	r := &facultyRoute{u, l}
	h := handler.Group("/faculty")
	{
		h.POST("/", r.create)
		h.GET("/", r.findAll)
	}
}

type createFacultyRequest struct {
	dto.CreateFaculty
}

type createFacultyResponse struct {
	*entity.Faculty
	dto.CUD
}

func (r *facultyRoute) create(c *gin.Context) {
	var body createFacultyRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		errorJSONResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	faculty, err := r.u.Create(c.Request.Context(), body.CreateFaculty)
	if err != nil {
		r.l.Error("sign up", slog.Any(err.Error(), err))
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, createFacultyResponse{Faculty: faculty})
}

type findAllFacultyResponse struct {
	Faculties []entity.Faculty `json:"faculties"`
}

func (r *facultyRoute) findAll(c *gin.Context) {
	faculties, err := r.u.FindAll(c.Request.Context())
	if err != nil {
		r.l.Error("sign up", slog.Any(err.Error(), err))
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, findAllFacultyResponse{Faculties: faculties})
}
