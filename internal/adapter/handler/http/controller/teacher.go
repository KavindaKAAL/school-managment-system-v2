package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/handler/http/dto/teacher"
	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/handler/http/network"
	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/storage/postgres/repository"
	"github.com/KavindaKAAL/school-management-system-v2/internal/core/domain"
	"github.com/KavindaKAAL/school-management-system-v2/internal/core/port"
	"github.com/gin-gonic/gin"
)

type teacherController struct {
	port.BaseController
	service port.TeacherService
}

func NewTeacherController(service port.TeacherService) port.Controller {

	return &teacherController{
		BaseController: NewBaseController("/teachers"),
		service:        service,
	}
}

func (c *teacherController) MountRoutes(group *gin.RouterGroup) {
	group.GET("/", c.getTeachers)
	group.GET("/:email", c.getTeacherByEmail)
	group.POST("/", c.createTeacher)
	group.PUT("/", c.updateTeacher)
	group.DELETE("/:email", c.deleteTeacher)
}

func (c *teacherController) getTeachers(ctx *gin.Context) {

	teachers, err := c.service.GetAllTeacherService()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Internel error"})
		return
	}
	res := teacher.NewTeachersListResDto(teachers)

	ctx.JSON(http.StatusOK, res)
}

func (c *teacherController) getTeacherByEmail(ctx *gin.Context) {

	dto, err := network.ReqParams(ctx, teacher.EmptyGetTeacherReqDto())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request param"})
		return
	}

	teacherRes, err := c.service.GetTeacherByEmailService(dto.Email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Teacher is not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	res := teacher.NewTeacherResDto(teacherRes)
	ctx.JSON(http.StatusOK, res)
}

func (c *teacherController) createTeacher(ctx *gin.Context) {
	dto, err := network.ReqBody(ctx, teacher.EmptyCreateTeacherDto())

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

	newTeacher, err2 := domain.NewTeacher(dto.Name, dto.Email)

	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internel error"})
		return
	}
	err = c.service.CreateTeacherService(newTeacher)

	if err != nil {
		if errors.Is(err, repository.ErrEmailAlreadyInUse) {
			ctx.JSON(http.StatusConflict, gin.H{"error": "Email is already used"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internel error"})
		}

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Successfully created"})

}

func (c *teacherController) updateTeacher(ctx *gin.Context) {
	dto, err := network.ReqBody(ctx, teacher.EmptyUpdateTeacherDto())

	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid request body")
		return
	}

	updateTeacher, err2 := domain.NewTeacher(dto.Name, dto.Email)

	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internel error"})
		return
	}

	err3 := c.service.UpdateTeacherService(updateTeacher)

	if err3 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internel error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully updated"})

}

func (c *teacherController) deleteTeacher(ctx *gin.Context) {

	dto, err := network.ReqParams(ctx, teacher.EmptyGetTeacherReqDto())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid request param")
		return
	}

	err2 := c.service.DeleteTeacherByEmailService(dto.Email)
	if err2 != nil {
		if errors.Is(err2, repository.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Teacher is not found"})
			return
		} else if errors.Is(err2, repository.ErrUserAssignedToSomeClasses) {
			ctx.JSON(http.StatusConflict, gin.H{"error": "Cannot delete the teacher, is still assigned to classes"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internel error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully deleted"})

}
