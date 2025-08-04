package controller

import (
	"encoding/json"
	"errors"
	"fmt"

	"net/http"

	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/handler/http/dto/student"
	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/handler/http/network"
	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/storage/postgres/repository"
	"github.com/KavindaKAAL/school-management-system-v2/internal/core/domain"
	"github.com/KavindaKAAL/school-management-system-v2/internal/core/port"
	"github.com/gin-gonic/gin"
)

type studentController struct {
	port.BaseController
	service port.StudentService
}

func NewStudentController(service port.StudentService) port.Controller {

	return &studentController{
		BaseController: NewBaseController("/students"),
		service:        service,
	}
}

func (c *studentController) MountRoutes(group *gin.RouterGroup) {
	group.GET("/", c.getStudents)
	group.GET("/:email", c.getStudentByEmail)
	group.POST("/", c.createStudent)
	group.PUT("/", c.updateStudent)
	group.DELETE("/:email", c.deleteStudent)
	group.PUT("/enroll", c.enrollStudentToClass)
	group.PUT("/unenroll", c.unEnrollStudentFromClass)
}

func (c *studentController) getStudents(ctx *gin.Context) {

	students, err := c.service.GetAllStudentService()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "yet to handle"})
		return
	}
	res := student.NewStudentsListResDto(students)

	ctx.JSON(http.StatusOK, res)
}

func (c *studentController) getStudentByEmail(ctx *gin.Context) {

	dto, err := network.ReqParams(ctx, student.EmptyGetStudentReqDto())
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	studentRes, err := c.service.GetStudentByEmailService(dto.Email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Student is not found"})
			return
		}
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	res := student.NewStudentResDto(studentRes)
	ctx.IndentedJSON(http.StatusOK, res)

}

func (c *studentController) createStudent(ctx *gin.Context) {
	dto, err := network.ReqBody(ctx, student.EmptyCreateStudentDto())

	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Malformed JSON: Syntax error"})
		case errors.As(err, &unmarshalTypeError):
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid type for field %s", unmarshalTypeError.Field)})
		default:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		}
		return
	}

	newStudent, err2 := domain.NewStudent(dto.Name, dto.Email)

	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internel error"})
		return
	}

	err2 = c.service.CreateStudentService(newStudent)

	if err2 != nil {
		if errors.Is(err2, repository.ErrEmailAlreadyInUse) {
			ctx.JSON(http.StatusConflict, gin.H{"error": "Email is already used"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internel error"})
		}

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Successfully created"})

}

func (c *studentController) updateStudent(ctx *gin.Context) {

	dto, err := network.ReqBody(ctx, student.EmptyUpdateStudentDto())

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	updateStudent, err2 := domain.NewStudent(dto.Name, dto.Email)

	if err2 != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	_, err3 := c.service.UpdateStudentService(updateStudent)

	if err3 != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully updated"})

}

func (c *studentController) deleteStudent(ctx *gin.Context) {

	dto, err := network.ReqParams(ctx, student.EmptyGetStudentReqDto())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	err2 := c.service.DeleteStudentByEmailService(dto.Email)
	if err2 != nil {
		if errors.Is(err2, repository.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Student is not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, err2)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully deleted"})

}

func (c *studentController) enrollStudentToClass(ctx *gin.Context) {

	dto, err := network.ReqBody(ctx, student.EmptyEnrollStudentToClassDto())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err2 := c.service.EnrollStudentService(dto.StudentEmail, dto.ClassName)

	if err2 != nil {

		if errors.Is(err2, repository.ErrStudentAlreadyEnrolled) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Student is already enrolled in the class"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Enrollment failed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully enrolled"})

}

func (c *studentController) unEnrollStudentFromClass(ctx *gin.Context) {
	dto, err := network.ReqBody(ctx, student.EmptyEnrollStudentToClassDto())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	err2 := c.service.UnEnrollStudentService(dto.StudentEmail, dto.ClassName)

	if err2 != nil {

		if errors.Is(err2, repository.ErrStudentNotEnrolled) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Student is not enrolled in the class"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Un-enrollment failed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully unenrolled"})

}
