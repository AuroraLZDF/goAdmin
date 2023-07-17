// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package auth

import (
	"apis/internal/pkg/core"
	"apis/internal/pkg/errno"
	"apis/internal/pkg/request/apis/v1"
	"apis/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Login 登录 apis 并返回一个 JWT Token.
func (ctrl *AdminController) Login(c *gin.Context) {
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
