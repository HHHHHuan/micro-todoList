package weblib

import (
	"api/weblib/handlers"
	"api/weblib/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter(service ...interface{}) *gin.Engine{
	r:=gin.Default()
	r.Use(middleware.Cors(),middleware.InitMiddleware(service),middleware.ErrorMiddleware())
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))

	v1:=r.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200,"success")
		})

		//用户服务
		v1.POST("/user/login",handlers.UserLogin)
		v1.POST("/user/register",handlers.UserRegister)

		// 需要登录保护
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.GET("tasks", handlers.GetTasksList)
			authed.POST("task", handlers.CreateTask)
			authed.GET("task/:id", handlers.GetTaskDetail) // task_id
			authed.PUT("task/:id", handlers.UpdateTask)    // task_id
			authed.DELETE("task/:id", handlers.DeleteTask) // task_id
		}
	}
	return r
}
