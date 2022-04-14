package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/*
* @Author: hh
* @Date:   2022/4/14 19:21
 */

// InitMiddleware 讲实例存在gin.Keys中
func InitMiddleware(service []interface{}) gin.HandlerFunc{
	return func(c *gin.Context) {
		c.Keys=make(map[string]interface{})
		c.Keys["userService"]=service[0]
		c.Next()
	}
}

// 错误处理中间件

func ErrorMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		defer func(){
			if err:=recover();err!=nil{
				c.JSON(200,gin.H{
					"code":404,
					"msg":fmt.Sprintf("%s",err),
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
