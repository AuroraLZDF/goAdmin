// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package auth

import (
	"github.com/gin-gonic/gin"

	"apis/internal/pkg/core"
	"apis/internal/pkg/errno"
	"apis/internal/pkg/request/apis/v1"
	"apis/pkg/log"
)

// Login 登录 apis 并返回一个 JWT Token.
func (ctrl *AdminController) Login(c *gin.Context) {
	log.C(c).Info("Login function called")

	var r v1.LoginRequest
	if err := c.ShouldBind(&r); err != nil {
		core.Error(c, errno.ErrBind)
		return
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
