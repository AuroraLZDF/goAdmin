// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package apis

import (
	"apis/internal/pkg/core"
	"apis/internal/pkg/errno"
	"apis/pkg/log"
	"github.com/gin-gonic/gin"
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

		core.Success(c, gin.H{"status": "ok"}, "ok")
	})

	//uc := user.New(store.S)
	//pc := post.New(store.S)

	//g.POST("/login", uc.Login)

	// 创建 v1 路由分组
	//v1 := g.Group("/v1")
	//{
	//	//创建 users 路由分组
	//	userv1 := v1.Group("/users").Use(middleware.Authn())
	//	{
	//		userv1.GET("info", uc.Detail)
	//		userv1.POST("create", uc.Create)
	//		userv1.PUT("change-password", uc.ChangePassword)
	//		userv1.POST("logout", uc.Logout)
	//		userv1.PUT(":name", uc.Update)    // 更新用户
	//		userv1.GET("", uc.List)           // 列出用户列表，只有 root 用户才能访问
	//		userv1.DELETE(":name", uc.Delete) // 删除用户
	//	}
	//
	//}

	return nil
}
