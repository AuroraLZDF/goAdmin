// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package apis

import (
	"github.com/gin-gonic/gin"

	"apis/internal/apis/controller/v1/admin/auth"
	"apis/internal/apis/store"
	"apis/internal/pkg/core"
	"apis/internal/pkg/errno"
	"apis/internal/pkg/middleware"
	"apis/pkg/log"
)

// installRouters 安装 apis 接口路由.
func installRouters(g *gin.Engine) error {
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.Error(c, errno.ErrPageNotFound)
	})

	// 注册 /healthz handler.
	g.GET("/healthz", func(c *gin.Context) {
		log.C(c).Info("Healthz function called")

		core.Success(c, nil, "ok")
	})

	ac := auth.New(store.S)

	// 创建 v1 路由分组
	v1 := g.Group("/v1")
	{
		admin := v1.Group("/admin")
		{
			_auth := admin.Group("/auth")
			{
				_auth.POST("/login", ac.Login)

				//创建 users 路由分组
				user := _auth.Use(middleware.Authn())
				{
					user.GET("userInfo", ac.UserInfo)
					user.POST("refresh", ac.RefreshToken)
					user.POST("logout", ac.Logout)
				}
			}
		}
	}

	return nil
}
