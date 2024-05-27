package v1

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sethvargo/go-password/password"

	"github.com/vicdevcode/ystudent/main/internal/dto"
	"github.com/vicdevcode/ystudent/main/internal/entity"
	"github.com/vicdevcode/ystudent/main/internal/usecase"
)

type employeeRoute struct {
	rmq *RabbitMQ
	u   usecase.Employee
	l   *slog.Logger
}

func newEmployee(router *router, rmq *RabbitMQ, u usecase.Employee, l *slog.Logger) {
	r := &employeeRoute{rmq, u, l}
	{
		router.protected.POST("/employee/", r.create)
		router.public.GET("/employees/", r.findAll)
	}
}

type createEmployeeRequest dto.CreateEmployee

type createEmployeeResponse dto.EmployeeResponse

func (r *employeeRoute) create(c *gin.Context) {
	var body createEmployeeRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		badRequest(c, err.Error())
		return
	}

	employee, err := r.u.Create(c.Request.Context(), dto.CreateEmployee(body))
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	password, err := password.Generate(8, 8, 0, false, false)
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	response, err := json.Marshal(dto.EmployeeUserResponse{
		User:     &employee.User,
		Password: password,
	})
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, createEmployeeResponse{
		Employee: employee,
		Password: password,
	})

	r.rmq.ch.PublishWithContext(
		c.Request.Context(),
		r.rmq.exchange,
		"main.employee.created",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        response,
		},
	)
}

type findAllEmployeeResponse struct {
	Employees []entity.Employee `json:"employees"`
}

func (r *employeeRoute) findAll(c *gin.Context) {
	page, err := GetPage(c.Query("page"), c.Query("count"))
	if err != nil {
		badRequest(c, err.Error())
		return
	}
	employees, err := r.u.FindAll(c.Request.Context(), page)
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, findAllEmployeeResponse{
		Employees: employees,
	})
}

func (r *employeeRoute) delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		badRequest(c, err.Error())
		return
	}
	if err := r.u.Delete(c.Request.Context(), id); err != nil {
		internalServerError(c, err.Error())
		return
	}

	response, err := json.Marshal(dto.Deleted{ID: id})
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, deleteGroupResponse{
		Message: "moderator was deleted",
	})

	r.rmq.ch.PublishWithContext(
		c.Request.Context(),
		r.rmq.exchange,
		"main.employee.deleted",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        response,
		},
	)
}
