package router

import (
	"fmt"

	"github.com/KavindaKAAL/school-management-system-v2/internal/core/port"
	"github.com/gin-gonic/gin"
)

type router struct {
	engine *gin.Engine
}

func NewRouter(mode string) port.Router {
	gin.SetMode(mode)
	eng := gin.Default()
	r := router{
		engine: eng,
	}
	return &r
}

func (r *router) LoadControllers(controllers []port.Controller) {
	api := r.GetEngine().Group("/api/v1")

	for _, c := range controllers {
		routeGroup := api.Group(c.Path())
		c.MountRoutes(routeGroup)
	}
}

func (r *router) GetEngine() *gin.Engine {
	return r.engine
}

func (r *router) Start(ip string, port uint16) {
	address := fmt.Sprintf("%s:%d", ip, port)
	r.engine.Run(address)
}

func (r *router) AttachMiddleware() *gin.Engine {
	return r.engine
}
