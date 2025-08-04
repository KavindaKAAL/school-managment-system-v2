package controller

import (
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
	msg, err := c.service.GetAllClassService()

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

func (c *classController) getClassByName(ctx *gin.Context) {

	d, err := network.ReqParams(ctx, class.EmptyGetClassReqDto())
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	student, err := c.service.GetClassByNameService(d.Name)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, student)
	// c.Send(ctx).SuccessDataResponse("success", data)
}

func (c *classController) createClass(ctx *gin.Context) {
	body, _ := network.ReqBody(ctx, class.EmptyCreateClassDto())
	student, err := domain.NewClass(body.Name, body.Name)
	fmt.Println("test100", student)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	msg, _ := c.service.CreateClassService(student)
	ctx.IndentedJSON(http.StatusCreated, msg)
	// if err != nil {
	// 	c.Send(ctx).InternalServerError("something went wrong", err)
	// 	return
	// }

	// data, err := utils.MapTo[dto.InfoMessage](msg)
	// if err != nil {
	// 	c.Send(ctx).InternalServerError("something went wrong", err)
	// 	return
	// }

	// c.Send(ctx).SuccessDataResponse("message received successfully!", data)
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

func (c *classController) updateClass(ctx *gin.Context) {
	fmt.Println("lahiru test 4")
	body, err := network.ReqBody(ctx, class.EmptyUpdateClassDto())

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	updateClass, err2 := domain.NewClass(body.Name, body.Subject)

	if err2 != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	msg, err2 := c.service.UpdateClassService(updateClass)

	if err2 != nil {
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

func (c *classController) deleteClass(ctx *gin.Context) {
	name := ctx.Param("name")
	fmt.Println("lahiru test 9", name)
	d, err := network.ReqParams(ctx, class.EmptyGetClassReqDto())
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	fmt.Println("lahiru test 10", d)
	_, err2 := c.service.DeleteClassByNameService(d.Name)
	if err2 != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err2)
		return
	}
	ctx.IndentedJSON(200, "Successfully deleted")

}

func (c *classController) assignTeacher(ctx *gin.Context) {
	body, err := network.ReqBody(ctx, class.EmptyAssignTeacherToClassDto())

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err2 := c.service.AssignTeacherService(body.ClassName, body.TeacherEmail)

	if err2 != nil {

		if errors.Is(err2, repository.ErrTeacherAlreadyAssigned) {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Teacher is already assigned in the class"})
			return
		}
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Assignment failed"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Assigned successfully"})
}

func (c *classController) unAssignTeacher(ctx *gin.Context) {
	body, err := network.ReqBody(ctx, class.EmptyAssignTeacherToClassDto())
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err2 := c.service.UnAssignTeacherService(body.ClassName, body.TeacherEmail)

	if err2 != nil {

		if errors.Is(err2, repository.ErrTeacherNotAssigned) {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Teacher is not assigned in the class"})
			return
		}
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Un-assigned failed"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Un-assigned successfully"})

}
