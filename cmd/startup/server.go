package startup

import (
	"context"
	"net/http"

	"github.com/KavindaKAAL/school-management-system-v2/config"
	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/handler/http/middleware"
	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/handler/http/router"
	"github.com/KavindaKAAL/school-management-system-v2/internal/adapter/storage/postgres"
	"github.com/KavindaKAAL/school-management-system-v2/internal/core/port"
	"github.com/gin-gonic/gin"
)

type Shutdown = func()

func Server() {
	env := config.NewEnv(".env", true)
	router, _, shutdown := create(env)
	defer shutdown()
	router.Start(env.ServerHost, env.ServerPort)
}

func create(env *config.Env) (port.Router, Module, Shutdown) {

	context := context.Background()

	dbConfig := postgres.DbConfig{
		User: env.DBUser,
		Pwd:  env.DBUserPwd,
		Host: env.DBHost,
		Port: env.DBPort,
		Name: env.DBName,
	}

	db := postgres.NewDatabase(context, dbConfig)
	db.Connect()

	module := NewModule(context, env, db)

	router := router.NewRouter(env.GoMode)
	router.GetEngine().Use(middleware.RequireJSONContentType())

	router.GetEngine().NoRoute(func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodPost:
			c.JSON(http.StatusNotFound, gin.H{"error": "POST route not found"})
		case http.MethodGet:
			c.JSON(http.StatusNotFound, gin.H{"error": "GET route not found"})
		case http.MethodPut:
			c.JSON(http.StatusNotFound, gin.H{"error": "PUT route not found"})
		case http.MethodDelete:
			c.JSON(http.StatusNotFound, gin.H{"error": "DELETE route not found"})
		default:
			c.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
		}
	})

	router.LoadControllers(module.Controllers())

	shutdown := func() {
		db.Disconnect()
	}

	return router, module, shutdown
}
