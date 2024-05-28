package v1

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"

	"github.com/vicdevcode/ystudent/main/internal/dto"
	"github.com/vicdevcode/ystudent/main/internal/entity"
	"github.com/vicdevcode/ystudent/main/internal/usecase"
)

type departmentRoute struct {
	rmq *RabbitMQ
	ud  usecase.Department
	uf  usecase.Faculty
	l   *slog.Logger
}

func newDepartment(
	router *router,
	rmq *RabbitMQ,
	ud usecase.Department,
	uf usecase.Faculty,
	l *slog.Logger,
) {
	r := &departmentRoute{rmq, ud, uf, l}
	{
		router.protected.POST("/department/", r.create)
		router.protected.PUT("/department/:id", r.update)
		router.protected.DELETE("/department/:id", r.delete)
		router.protected.POST("/department/add-employee", r.addEmployee)
		router.protected.POST("/department/delete-employee", r.deleteEmployee)
		router.public.GET("/departments", r.findAll)
	}
}

type createDepartmentRequest dto.CreateDepartment

type createDepartmentResponse *entity.Department

func (r *departmentRoute) create(c *gin.Context) {
	var body createDepartmentRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		badRequest(c, err.Error())
		return
	}

	department, err := r.ud.Create(c.Request.Context(), dto.CreateDepartment(body))
	if err != nil {
		r.l.Error(err.Error())
		internalServerError(c, err.Error())
		return
	}

	response, err := json.Marshal(department)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, createDepartmentResponse(department))

	r.rmq.ch.PublishWithContext(
		c.Request.Context(),
		r.rmq.exchange,
		"main.department.created",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        response,
		},
	)
}

type findAllDepartmentResponse struct {
	Departments []dto.FindAllDepartmentResponse `json:"departments"`
	dto.Page
}

func (r *departmentRoute) findAll(c *gin.Context) {
	page, err := GetPage(c.Query("page"), c.Query("count"))
	if err != nil {
		badRequest(c, err.Error())
		return
	}

	departments, err := r.ud.FindAll(c.Request.Context(), page)
	if err != nil {
		r.l.Error(err.Error())
		internalServerError(c, err.Error())
		return
	}

	var response []dto.FindAllDepartmentResponse
	for _, department := range departments {
		faculty, err := r.uf.FindOne(c.Request.Context(), entity.Faculty{
			ID: *department.FacultyID,
		})
		if err != nil {
			r.l.Error(err.Error())
			internalServerError(c, err.Error())
			return
		}
		response = append(response, dto.FindAllDepartmentResponse{
			Department:  &department,
			FacultyName: faculty.Name,
		})
	}

	c.JSON(http.StatusOK, findAllDepartmentResponse{Departments: response, Page: page})
}

type updateDepartmentRequest dto.UpdateDepartmentBody

type updateDepartmentResponse *entity.Department

func (r *departmentRoute) update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		badRequest(c, err.Error())
		return
	}

	var body updateDepartmentRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		badRequest(c, err.Error())
		return
	}

	r.l.Info("", slog.Any("", body))

	department, err := r.ud.Update(c.Request.Context(), dto.UpdateDepartment{
		ID:        id,
		Name:      body.Name,
		FacultyID: body.FacultyID,
	})
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	department, err = r.ud.FindOne(c.Request.Context(), entity.Department{ID: id})

	response, err := json.Marshal(department)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, updateDepartmentResponse(department))

	r.rmq.ch.PublishWithContext(
		c.Request.Context(),
		r.rmq.exchange,
		"main.department.updated",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        response,
		},
	)
}

type deleteDepartmentResponse struct {
	Message string `json:"message"`
}

func (r *departmentRoute) delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		badRequest(c, err.Error())
		return
	}
	if err := r.ud.Delete(c.Request.Context(), id); err != nil {
		internalServerError(c, err.Error())
		return
	}

	response, err := json.Marshal(dto.Deleted{ID: id})
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, deleteDepartmentResponse{
		Message: "department was deleted",
	})

	r.rmq.ch.PublishWithContext(
		c.Request.Context(),
		r.rmq.exchange,
		"main.department.deleted",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        response,
		},
	)
}

type addEmployeeToDepartmentRequest dto.AddEmployeeToDepartment

type addEmployeeToDepartmentResponse *entity.Department

func (r *departmentRoute) addEmployee(c *gin.Context) {
	var body addEmployeeToDepartmentRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		badRequest(c, err.Error())
		return
	}

	department, err := r.ud.AddEmployee(c.Request.Context(), dto.AddEmployeeToDepartment(body))
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	response, err := json.Marshal(body)
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, addEmployeeToDepartmentResponse(department))

	r.rmq.ch.PublishWithContext(
		c.Request.Context(),
		r.rmq.exchange,
		"main.department.add_employee",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        response,
		},
	)
}

type deleteEmployeeFromDepartmentRequest dto.DeleteEmployeeFromDepartment

type deleteEmployeeDepartmentResponse *entity.Department

func (r *departmentRoute) deleteEmployee(c *gin.Context) {
	var body deleteEmployeeFromDepartmentRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		badRequest(c, err.Error())
		return
	}

	department, err := r.ud.DeleteEmployee(
		c.Request.Context(),
		dto.DeleteEmployeeFromDepartment(body),
	)
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	response, err := json.Marshal(body)
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, deleteEmployeeDepartmentResponse(department))

	r.rmq.ch.PublishWithContext(
		c.Request.Context(),
		r.rmq.exchange,
		"main.department.delete_employee",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        response,
		},
	)
}
