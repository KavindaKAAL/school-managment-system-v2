package controller

import (
	"encoding/json"
	"errors"
	"fmt"

	"net/http"

	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/handler/http/dto/class"
	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/handler/http/network"
	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/storage/postgres/repository"
	"github.com/KavindaKAAL/school-management-system-v2/internal/core/domain"
	"github.com/KavindaKAAL/school-management-system-v2/internal/core/port"
	"github.com/gin-gonic/gin"
)

type classController struct {
	port.BaseController
	service port.ClassService
}

func NewClassController(service port.ClassService) port.Controller {

	return &classController{
		BaseController: NewBaseController("/classes"),
		service:        service,
	}
}

func (c *classController) MountRoutes(group *gin.RouterGroup) {
	group.GET("/", c.getClasses)
	group.GET("/:name", c.getClassByName)
	group.POST("/", c.createClass)
	group.PUT("/", c.updateClass)
	group.DELETE("/:name", c.deleteClass)
	group.PUT("/assign", c.assignTeacher)
	group.PUT("/unassign", c.unAssignTeacher)
}

func (c *classController) getClasses(ctx *gin.Context) {
	classes, err := c.service.GetAllClassService()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internel error"})
		return
	}
	res := class.NewClassesListResDto(classes)

	ctx.JSON(http.StatusOK, res)

}

func (c *classController) getClassByName(ctx *gin.Context) {

	dto, err := network.ReqParams(ctx, class.EmptyGetClassReqDto())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid request param"})
		return
	}

	classRes, err := c.service.GetClassByNameService(dto.Name)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Class is not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	res := class.NewClassResDto(classRes)

	ctx.JSON(http.StatusOK, res)
}

func (c *classController) createClass(ctx *gin.Context) {

	dto, err := network.ReqBody(ctx, class.EmptyCreateClassDto())

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

	newClass, err2 := domain.NewClass(dto.Name, dto.Subject)

	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internel error"})
		return
	}

	err2 = c.service.CreateClassService(newClass)

	if err2 != nil {
		if errors.Is(err2, repository.ErrClassNameAlreadyInUse) {
			ctx.JSON(http.StatusConflict, gin.H{"error": "Class name is already used"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internel error"})
		}

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Successfully created"})

}

func (c *classController) updateClass(ctx *gin.Context) {
	dto, err := network.ReqBody(ctx, class.EmptyUpdateClassDto())

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	updateClass, err2 := domain.NewClass(dto.Name, dto.Subject)

	if err2 != nil {
		ctx.JSON(http.StatusBadRequest, err2)
		return
	}

	err3 := c.service.UpdateClassService(updateClass)

	if err3 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internel error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully updated"})
}

func (c *classController) deleteClass(ctx *gin.Context) {

	dto, err := network.ReqParams(ctx, class.EmptyGetClassReqDto())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	err2 := c.service.DeleteClassByNameService(dto.Name)
	if err2 != nil {
		if errors.Is(err2, repository.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Class is not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internel error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully deleted"})

}

func (c *classController) assignTeacher(ctx *gin.Context) {
	body, err := network.ReqBody(ctx, class.EmptyAssignTeacherToClassDto())

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err2 := c.service.AssignTeacherService(body.ClassName, body.TeacherEmail)

	if err2 != nil {

		if errors.Is(err2, repository.ErrTeacherAlreadyAssigned) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Teacher is already assigned in the class"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Assignment failed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully assigned"})
}

func (c *classController) unAssignTeacher(ctx *gin.Context) {
	body, err := network.ReqBody(ctx, class.EmptyAssignTeacherToClassDto())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err2 := c.service.UnAssignTeacherService(body.ClassName, body.TeacherEmail)

	if err2 != nil {

		if errors.Is(err2, repository.ErrTeacherNotAssigned) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Teacher is not assigned in the class"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Un-assigned failed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully unassigned"})

}
