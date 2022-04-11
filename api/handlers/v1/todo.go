package v1

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	_ "github.com/Shahboz4131/api-gateway/api/handlers/models"
	pb "github.com/Shahboz4131/api-gateway/genproto"
	l "github.com/Shahboz4131/api-gateway/pkg/logger"
	"github.com/Shahboz4131/api-gateway/pkg/utils"
)

// CreateTask ...
// @Summary CreateTask
// @Description This API for creating a new task
// @Tags task
// @Accept  json
// @Produce  json
// @Param Task request body models.Task true "taskCreateRequest"
// @Success 200 {object} models.Task
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/tasks/ [post]
func (h *handlerV1) CreateTask(c *gin.Context) {
	var (
		body         pb.Task
		jsonbMarshal protojson.MarshalOptions
	)
	jsonbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TaskService().Create(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create task", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetTask ...
// @Router /v1/tasks/{id} [get]
// @Summary GetTask
// @Description This API for getting task detail
// @Tags task
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.Task
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) GetTask(c *gin.Context) {
	var jsonbMarshal protojson.MarshalOptions
	jsonbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TaskService().Get(
		ctx, &pb.ByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get task", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListTasks ...
// @Router /v1/tasks [get]
// @Summary ListTasks
// @Description This API for getting list of tasks
// @Tags task
// @Accept  json
// @Produce  json
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} models.ListTasks
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) ListTasks(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to parse query params json" + errStr[0])
		return
	}

	var jsonbMarshal protojson.MarshalOptions
	jsonbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TaskService().List(
		ctx, &pb.ListReq{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list tasks", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateTask ...
// @Router /v1/tasks/{id} [put]
// @Summary UpdateTask
// @Description This API for updating task
// @Tags task
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param Task request body models.UpdateTask true "taskUpdateRequest"
// @Success 200
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) UpdateTask(c *gin.Context) {
	var (
		body         pb.Task
		jsonbMarshal protojson.MarshalOptions
	)
	jsonbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	body.Id = c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TaskService().Update(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update task", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteTask ...
// @Router /v1/tasks/{id} [delete]
// @Summary DeleteTask
// @Description This API for deleting task
// @Tags task
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) DeleteTask(c *gin.Context) {
	var jsonbMarshal protojson.MarshalOptions
	jsonbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TaskService().Delete(
		ctx, &pb.ByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete task", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// OverdueTasks ...
// @Router /v1/overduetasks [get]
// @Summary OverdueTasks
// @Description This API for getting list of overdue tasks
// @Tags task
// @Accept  json
// @Produce  json
// @Param Task request body models.Overdue true "taskOverdueRequest"
// @Success 200
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) OverdueTasks(c *gin.Context) {
	body := pb.OverdueReq{}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	fmt.Println(body)

	var jsonbMarshal protojson.MarshalOptions
	jsonbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.TaskService().Overdue(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list tasks ", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
