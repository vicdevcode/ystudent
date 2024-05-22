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
	u   usecase.Department
	l   *slog.Logger
}

func newDepartment(
	router *router,
	rmq *RabbitMQ,
	u usecase.Department,
	l *slog.Logger,
) {
	r := &departmentRoute{rmq, u, l}
	{
		router.protected.POST("/department/", r.create)
		router.protected.PUT("/department/:id", r.update)
		router.protected.DELETE("/department/:id", r.delete)
		router.protected.POST("/department/add-employee/", r.addEmployee)
		router.protected.POST("/department/delete-employee/", r.deleteEmployee)
		router.public.GET("/departments/", r.findAll)
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

	department, err := r.u.Create(c.Request.Context(), dto.CreateDepartment(body))
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
	Departments []entity.Department `json:"departments"`
}

func (r *departmentRoute) findAll(c *gin.Context) {
	departments, err := r.u.FindAll(c.Request.Context())
	if err != nil {
		r.l.Error(err.Error())
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, findAllDepartmentResponse{Departments: departments})
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

	department, err := r.u.Update(c.Request.Context(), dto.UpdateDepartment{
		ID:        id,
		Name:      body.Name,
		FacultyID: body.FacultyID,
	})
	if err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, updateDepartmentResponse(department))
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
	if err := r.u.Delete(c.Request.Context(), id); err != nil {
		internalServerError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, deleteDepartmentResponse{
		Message: "department was deleted",
	})
}

type addEmployeeToDepartmentRequest dto.AddEmployeeToDepartment

type addEmployeeToDepartmentResponse *entity.Department

func (r *departmentRoute) addEmployee(c *gin.Context) {
	var body addEmployeeToDepartmentRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		badRequest(c, err.Error())
		return
	}

	department, err := r.u.AddEmployee(c.Request.Context(), dto.AddEmployeeToDepartment(body))
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

	department, err := r.u.DeleteEmployee(
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
