package core

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"user/model"
	"user/services"
)

/*
* @Author: hh
* @Date:   2022/4/14 15:13
 */

func BuildUser(user model.User) *services.UserModel {
	userModel := &services.UserModel{
		ID:        uint32(user.ID),
		UserName:  user.UserName,
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
	}
	return userModel
}

func (*UserService) UserLogin(ctx context.Context, req *services.UserRequest, resp *services.UserDetailResponse) error {
	var user model.User
	resp.Code = 200

	if err := model.DB.Where("user_name=?", req.UserName).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.Code = 400
			return nil
		}
		resp.Code = 500
		return nil
	}
	// 密码错误
	if !user.CheckPassword(req.Password) {
		resp.Code = 400
		return nil
	}

	resp.UserDetail = BuildUser(user)
	return nil
}

func (*UserService) UserRegister(ctx context.Context, req *services.UserRequest, resp *services.UserDetailResponse) error {
	if req.Password!=req.PasswordConfirm{
		err:=errors.New("两次密码不一致")
		return err
	}

	var count int64
	if err:=model.DB.Model(&model.User{}).Where("user_name=?",req.UserName).Count(&count).Error;err!=nil{
		return err
	}
	if count>0{
		err:=errors.New("用户名已存在")
		return err
	}

	user:=model.User{
		UserName: req.UserName,
	}
	// 加密密码
	if err:=user.SetPassword(req.Password);err!=nil{
		return err
	}
	if err:=model.DB.Create(&user).Error;err!=nil{
		return err
	}
	resp.UserDetail=BuildUser(user)
	return nil
}
