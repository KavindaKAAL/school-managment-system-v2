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

	msg, err := c.service.GetAllTeacherService()

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, msg)

	// mongoId, err := network.ReqParams(ctx, coredto.EmptyMongoId())
	// if err != nil {
	// 	c.Send(ctx).BadRequestError(err.Error(), err)
	// 	return
	// }

	// data, err := c.service.GetUserPublicProfile(mongoId.ID)
	// if err != nil {
	// 	c.Send(ctx).MixedError(err)
	// 	return
	// }

	// c.Send(ctx).SuccessDataResponse("success", data)
}

func (c *teacherController) getTeacherByEmail(ctx *gin.Context) {

	d, err := network.ReqParams(ctx, teacher.EmptyGetTeacherReqDto())
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	student, err := c.service.GetTeacherByEmailService(d.Email)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, student)
	// c.Send(ctx).SuccessDataResponse("success", data)
}

func (c *teacherController) createTeacher(ctx *gin.Context) {
	dto, _ := network.ReqBody(ctx, teacher.EmptyCreateTeacherDto())
	newTeacher, err := domain.NewTeacher(dto.Name, dto.Email)

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

// Call BindJSON to bind the received JSON to
// newAlbum.
// 	var newAlbum album

// 	if err := ctx.BindJSON(&newAlbum); err != nil {
// 		ctx.IndentedJSON(http.StatusCreated, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Add the new album to the slice.
// 	albums = append(albums, newAlbum)
// 	fmt.Println(albums)

// 	ctx.IndentedJSON(http.StatusCreated, albums)
// }

func (c *teacherController) updateTeacher(ctx *gin.Context) {
	fmt.Println("lahiru test 4")
	body, err := network.ReqBody(ctx, teacher.EmptyUpdateTeacherDto())

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	updateTeacher, err2 := domain.NewTeacher(body.Name, body.Email)

	if err2 != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	msg, err3 := c.service.UpdateTeacherService(updateTeacher)

	if err3 != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, msg)

	// mongoId, err := network.ReqParams(ctx, coredto.EmptyMongoId())
	// if err != nil {
	// 	c.Send(ctx).BadRequestError(err.Error(), err)
	// 	return
	// }

	// data, err := c.service.GetUserPublicProfile(mongoId.ID)
	// if err != nil {
	// 	c.Send(ctx).MixedError(err)
	// 	return
	// }

	// c.Send(ctx).SuccessDataResponse("success", data)
}

func (c *teacherController) deleteTeacher(ctx *gin.Context) {

	d, err := network.ReqParams(ctx, teacher.EmptyGetTeacherReqDto())
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	_, err2 := c.service.DeleteTeacherByEmailService(d.Email)
	if err2 != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err2)
		return
	}
	ctx.IndentedJSON(200, "Successfully deleted")

}
