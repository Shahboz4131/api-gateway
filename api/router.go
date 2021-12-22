package api

import (
	"github.com/gin-gonic/gin"

	"github.com/Shahboz4131/api-gateway/api/handlers/v1"
	"github.com/Shahboz4131/api-gateway/config"
	"github.com/Shahboz4131/api-gateway/pkg/logger"
	"github.com/Shahboz4131/api-gateway/services"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

// New ...
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	api := router.Group("/v1")
	api.POST("/tasks", handlerV1.CreateTask)
	api.GET("/tasks/:id", handlerV1.GetTask)
	api.GET("/tasks", handlerV1.ListTasks)
	api.PUT("/tasks/:id", handlerV1.UpdateTask)
	api.DELETE("/tasks/:id", handlerV1.DeleteTask)
	api.GET("/overduetasks", handlerV1.OverdueTasks)

	return router
}
