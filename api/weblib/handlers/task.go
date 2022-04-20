package handlers

import (
	"api/pkg/utils"
	"api/service"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetTasksList(c *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(c.Bind(&taskReq))
	taskService := c.Keys["taskService"].(service.TaskService)
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	taskResp, err := taskService.GetTasksList(context.Background(), &taskReq)
	if err != nil {
		PanicIfTaskError(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"task":  taskResp.TaskList,
			"count": taskResp.Count,
		},
	})
}

func CreateTask(c *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(c.Bind(&taskReq))
	//从gin.keys取出服务实例
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	taskService := c.Keys["taskService"].(service.TaskService)
	taskRes, err := taskService.CreateTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	c.JSON(200, gin.H{"data": taskRes.TaskDetail})
}

func GetTaskDetail(c *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(c.BindUri(&taskReq))
	//从gin.keys取出服务实例
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	id, _ := strconv.Atoi(c.Param("id")) // 获取task_id
	taskReq.Id = uint64(id)
	productService := c.Keys["taskService"].(service.TaskService)
	productRes, err := productService.GetTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	c.JSON(200, gin.H{"data": productRes.TaskDetail})
}

func UpdateTask(c *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(c.Bind(&taskReq))
	//从gin.keys取出服务实例
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	id, _ := strconv.Atoi(c.Param("id"))
	taskReq.Id = uint64(id)
	taskReq.Uid = uint64(claim.Id)
	taskService := c.Keys["taskService"].(service.TaskService)
	taskRes, err := taskService.UpdateTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	c.JSON(200, gin.H{"data": taskRes.TaskDetail})
}

func DeleteTask(c *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(c.Bind(&taskReq))
	//从gin.keys取出服务实例
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	id, _ := strconv.Atoi(c.Param("id"))
	taskReq.Id = uint64(id)
	taskService := c.Keys["taskService"].(service.TaskService)
	taskRes, err := taskService.DeleteTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	c.JSON(200, gin.H{"data": taskRes.TaskDetail})
}
