package port

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type BaseController interface {
	//ResponseSender
	Path() string
	// Authentication() gin.HandlerFunc
	// Authorization(role string) gin.HandlerFunc
}

type Controller interface {
	BaseController
	MountRoutes(group *gin.RouterGroup)
}

type BaseService interface {
	Context() context.Context
}

type Dto[T any] interface {
	GetValue() *T
	ValidateErrors(errs validator.ValidationErrors) ([]string, error)
}

type BaseRouter interface {
	GetEngine() *gin.Engine
	// RegisterValidationParsers(tagNameFunc validator.TagNameFunc)
	// LoadRootMiddlewares(middlewares []RootMiddleware)
	Start(ip string, port uint16)
}

type Router interface {
	BaseRouter
	LoadControllers(controllers []Controller)
}

type BaseModule[T any] interface {
	GetInstance() *T
}

type Module[T any] interface {
	BaseModule[T]
	Controllers() []Controller
}
