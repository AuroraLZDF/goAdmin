// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package apis

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	user "apis/internal/apis/controller/v1/admin/admin"
	"apis/internal/apis/controller/v1/admin/admin/group"
	uLog "apis/internal/apis/controller/v1/admin/admin/log"
	"apis/internal/apis/controller/v1/admin/admin/role"
	"apis/internal/apis/controller/v1/admin/auth"
	"apis/internal/apis/controller/v1/admin/authenticate/identity"
	"apis/internal/apis/controller/v1/admin/profile"
	"apis/internal/apis/controller/v1/admin/system/attributes/area"
	"apis/internal/apis/controller/v1/admin/system/attributes/category"
	"apis/internal/apis/controller/v1/admin/system/attributes/general"
	"apis/internal/apis/controller/v1/admin/system/attributes/place"
	"apis/internal/apis/controller/v1/admin/system/config"
	"apis/internal/apis/controller/v1/admin/system/menu"
	"apis/internal/apis/store"
	"apis/internal/pkg/core"
	"apis/internal/pkg/errno"
	"apis/internal/pkg/middleware"
	"apis/pkg/log"
)

// installRouters 安装 apis 接口路由.
func installRouters(g *gin.Engine, db *gorm.DB) error {
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.Error(c, errno.ErrPageNotFound)
	})

	// 注册 /healthz handler.
	g.GET("/healthz", func(c *gin.Context) {
		log.C(c).Info("Healthz function called")

		core.Success(c, nil, "ok")
	})

	// 创建 v1 路由分组
	v1 := g.Group("/v1")
	{
		admin := v1.Group("/admin")
		{
			ac := auth.New(store.S)
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
			pr := profile.New(store.S)
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
					_area := attributes.Group("area")
					{
						ar := area.New(store.S)
						_area.GET("lists", ar.Lists)
						_area.GET("detail", ar.Detail)
						_area.POST("update", ar.Update)
						_area.POST("enable", ar.Enable)
						_area.POST("disable", ar.Disable)
					}
					//分类管理
					_category := attributes.Group("category")
					{
						ca := category.New(store.S)
						_category.GET("lists", ca.Lists)
						_category.GET("detail", ca.Detail)
						_category.POST("update", ca.Update)
						_category.POST("enable", ca.Enable)
						_category.POST("disable", ca.Disable)
					}
					//位置配置
					_place := attributes.Group("place")
					{
						pl := place.New(store.S)
						_place.GET("lists", pl.Lists)
						_place.GET("detail", pl.Detail)
						_place.POST("update", pl.Update)
					}
					//公共配置
					_general := attributes.Group("general")
					{
						ge := general.New(store.S)
						_general.GET("lists", ge.Lists)
						_general.GET("detail", ge.Detail)
						_general.POST("save", ge.Save)
						_general.POST("enable", ge.Enable)
						_general.POST("disable", ge.Disable)
					}
				}
				//站点设置
				_config := system.Group("config")
				{
					cf := config.New(store.S)
					_config.GET("detail", cf.Detail)
					_config.POST("save", cf.Save)
				}
				//菜单管理
				_menu := system.Group("menu")
				{
					mu := menu.New(store.S)
					_menu.GET("lists", mu.Lists)
					_menu.GET("detail", mu.Detail)
					_menu.GET("roleMenu", mu.RoleMenu)
					_menu.POST("update", mu.Update)
					_menu.POST("enable", mu.Enable)
					_menu.POST("disable", mu.Disable)
				}
			}
			// 管理员及组
			_admin := admin.Group("admin")
			{
				//管理员管理
				_user := _admin.Group("user")
				{
					us := user.New(store.S)
					_user.GET("lists", us.Lists)
					_user.GET("detail", us.Detail)
					_user.POST("update", us.Update)
					_user.POST("enable", us.Enable)
					_user.POST("disable", us.Disable)
				}
				// 权限组分配
				_group := _admin.Group("group")
				{
					gr := group.New(store.S)
					_group.GET("index", gr.Index)
					_group.POST("update", gr.Update)
				}
				// 权限管理
				_role := _admin.Group("role")
				{
					ro := role.New(store.S)
					_role.GET("lists", ro.Lists)
					_role.GET("detail", ro.Detail)
					_role.GET("rules", ro.Rules)
					_role.POST("update", ro.Update)
					_role.POST("enable", ro.Enable)
					_role.POST("disable", ro.Disable)
				}
				// 操作日志
				_log := _admin.Group("log")
				{
					lo := uLog.New(db)
					_log.GET("lists", lo.Lists)
				}
			}
			// 身份认证
			_authenticate := admin.Group("auth")
			{
				// 实名认证
				identity := identity.New(db)
				_authenticate.GET("lists", identity.Lists)
				_authenticate.GET("detail", identity.Detail)
				_authenticate.POST("pass", identity.Pass)
				_authenticate.POST("reject", identity.Reject)
			}
		}
	}

	return nil
}
