package v1

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"

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

type createEmployeeResponse *entity.Employee

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

	response, err := json.Marshal(employee)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, createEmployeeResponse(employee))

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
	employees, err := r.u.FindAll(c.Request.Context())
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, findAllEmployeeResponse{
		Employees: employees,
	})
}
