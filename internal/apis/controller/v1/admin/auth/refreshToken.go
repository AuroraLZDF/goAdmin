// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package auth

import (
	"github.com/gin-gonic/gin"

	"apis/internal/pkg/core"
	"apis/pkg/log"
)

// RefreshToken 刷新 Token
func (ctrl *AdminController) RefreshToken(c *gin.Context) {
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
