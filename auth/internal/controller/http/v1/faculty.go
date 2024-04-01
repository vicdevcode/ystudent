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

func newFaculty(handler *gin.RouterGroup, u usecase.Faculty, l *slog.Logger) {
	r := &facultyRoute{u, l}
	h := handler.Group("/faculty")
	{
		h.POST("/", r.create)
		h.GET("/", r.findAll)
	}
}

type createFacultyResponse struct {
	*entity.Faculty
	dto.CUD
}

func (r *facultyRoute) create(c *gin.Context) {
	var body dto.CreateFaculty

	if err := c.ShouldBindJSON(&body); err != nil {
		badRequest(c, err.Error())
		return
	}

	faculty, err := r.u.Create(c.Request.Context(), body)
	if err != nil {
		r.l.Error(err.Error())
		internalServerError(c, err.Error())
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
		r.l.Error(err.Error())
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, findAllFacultyResponse{Faculties: faculties})
}
