// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package apis

import (
	"github.com/gin-gonic/gin"

	"apis/internal/apis/controller/v1/admin/auth"
	"apis/internal/apis/controller/v1/admin/profile"
	"apis/internal/apis/controller/v1/admin/system/attributes/area"
	"apis/internal/apis/controller/v1/admin/system/attributes/category"
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
	pr := profile.New(store.S)
	ar := area.New(store.S)
	ca := category.New(store.S)

	// 创建 v1 路由分组
	v1 := g.Group("/v1")
	{
		admin := v1.Group("/admin")
		{
			// login 路由
			admin.POST("/auth/login", ac.Login)

			// 添加登录认证
			admin.Use(middleware.Authn())

			// 创建 auth 路由分组
			_auth := admin.Group("/auth")
			{
				_auth.GET("userInfo", ac.UserInfo)
				_auth.POST("refresh", ac.RefreshToken)
				_auth.POST("logout", ac.Logout)
			}
			// 个人资料
			_profile := admin.Group("/profile")
			{
				_profile.POST("save", pr.Save)
			}
			// 系统设置
			system := admin.Group("/system")
			{
				// 属性设置
				attributes := system.Group("/attributes")
				{
					//区域配置
					_area := attributes.Group("/area")
					{
						_area.GET("lists", ar.Lists)
						_area.GET("detail", ar.Detail)
						_area.POST("update", ar.Update)
						_area.POST("enable", ar.Enable)
						_area.POST("disable", ar.Disable)
					}
					//分类管理
					_category := attributes.Group("/category")
					{
						_category.GET("lists", ca.Lists)
						_category.GET("detail", ca.Detail)
						_category.POST("update", ca.Update)
						_category.POST("enable", ca.Enable)
						_category.POST("disable", ca.Disable)
					}
				}

				//位置配置
				//公共配置

				//站点设置

				//菜单管理
			}

		}
	}

	return nil
}
