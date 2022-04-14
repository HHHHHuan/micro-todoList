package handlers

import (
	"api/pkg/utils"
	"api/service"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
* @Author: hh
* @Date:   2022/4/14 19:38
 */

func UserRegister(c *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserError(c.Bind(&userReq))
	// 从gin.Key取出服务实例
	userService:=c.Keys["userService"].(service.UserService)
	userResp, err := userService.UserRegister(context.Background(), &userReq)
	PanicIfUserError(err)
	c.JSON(http.StatusOK,gin.H{"data":userResp})
}

func UserLogin(c *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserError(c.Bind(&userReq))
	// 从gin.Key取出服务实例
	userService:=c.Keys["userService"].(service.UserService)
	userResp, err := userService.UserLogin(context.Background(), &userReq)
	PanicIfUserError(err)
	token,err:=utils.GenerateToken(uint(userResp.UserDetail.ID))
	c.JSON(http.StatusOK,gin.H{
		"code":userResp.Code,
		"msg":"success",
		"data":gin.H{
			"user":userResp.UserDetail,
			"token":token,
	},
	})
}