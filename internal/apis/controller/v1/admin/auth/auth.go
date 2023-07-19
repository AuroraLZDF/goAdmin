// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"apis/internal/apis/biz"
	"apis/internal/apis/store"
	"apis/internal/pkg/core"
	"apis/internal/pkg/errno"
	v1 "apis/internal/pkg/request/apis/v1"
	"apis/pkg/log"
)

// Controller 是 admin 模块在 Controller 层的实现，用来处理用户模块的请求.
type Controller struct {
	b biz.IBiz
}

// New 创建一个 admin controller.
func New(ds store.IStore) *Controller {
	return &Controller{
		b: biz.NewBiz(ds),
	}
}

// Login 登录 apis 并返回一个 JWT Token.
func (ctrl *Controller) Login(c *gin.Context) {
	log.C(c).Info("Login function called")

	var r v1.LoginRequest
	if err := c.ShouldBind(&r); err != nil {
		core.Error(c, errno.ErrBind)
		return
	}

	//实例化验证器
	validate := validator.New()

	if errs := validate.Struct(r); errs != nil {
		//core.Error(c, errno.ErrInvalidParameter.SetMessage(errs.Error()))
		//return
		for _, err := range errs.(validator.ValidationErrors) {
			core.Error(c, errno.ErrInvalidParameter.SetMessage(err.Error()))
			return
		}
	}

	token, err := ctrl.b.Admins().Login(c, &r)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, gin.H{
		"access_token": token.AccessToken,
		"token_type":   token.TokenType,
		"expires_in":   token.ExpiresIn,
	}, "login success.")
}

// UserInfo 查看账户信息
func (ctrl *Controller) UserInfo(c *gin.Context) {
	log.C(c).Info("user detail info function called")

	info, err := ctrl.b.Admins().Info(c, c.GetString("phone"))
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, info, "success.")
}

// RefreshToken 刷新 Token
func (ctrl *Controller) RefreshToken(c *gin.Context) {
	log.C(c).Info("RefreshToken function called")

	token, err := ctrl.b.Admins().RefreshToken(c, c.Request, c.GetString("phone"))
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, gin.H{
		"access_token": token.AccessToken,
		"token_type":   token.TokenType,
		"expires_in":   token.ExpiresIn,
	}, "login success.")
}

// Logout 退出登录
func (ctrl *Controller) Logout(c *gin.Context) {
	log.C(c).Info("Logout function called")

	err := ctrl.b.Admins().Logout(c, c.Request)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, nil, "logout success.")
}
