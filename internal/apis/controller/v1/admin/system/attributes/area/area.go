// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package area

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

// Controller 是 area 模块在 Controller 层的实现，用来处理区域信息的请求.
type Controller struct {
	b biz.IBiz
}

// New 创建一个 area controller.
func New(ds store.IStore) *Controller {
	return &Controller{
		b: biz.NewBiz(ds),
	}
}

// Lists 区域信息列表
func (ctrl Controller) Lists(c *gin.Context) {
	log.C(c).Info("system attributes area lists function called")

	lists, err := ctrl.b.Areas().Lists(c, c.GetInt("pid"))
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, lists, "get area lists success")
}

// Detail 区域详情
func (ctrl Controller) Detail(c *gin.Context) {
	log.C(c).Info("system attributes area detail function called")

	var r v1.IdRequest
	if err := c.ShouldBind(&r); err != nil {
		core.Error(c, errno.ErrBind)
		return
	}

	validate := validator.New()
	if errs := validate.Struct(r); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			core.Error(c, errno.ErrInvalidParameter.SetMessage(err.Error()))
			return
		}
	}

	info, err := ctrl.b.Areas().Detail(c, r.Id)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, info, "get area detail success")
}

// Update 更新区域信息
func (ctrl Controller) Update(c *gin.Context) {
	log.C(c).Info("system attributes area update function called")

	var r v1.AreaUpdateRequest
	if err := c.ShouldBind(&r); err != nil {
		core.Error(c, errno.ErrBind)
		return
	}

	validate := validator.New()
	if errs := validate.Struct(r); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			core.Error(c, errno.ErrInvalidParameter.SetMessage(err.Error()))
			return
		}
	}

	id, err := ctrl.b.Areas().CreateOrUpdate(c, r)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, gin.H{"id": id}, "update area info success")
}

// Enable 启用区域信息
func (ctrl Controller) Enable(c *gin.Context) {
	log.C(c).Info("system attributes area enable function called")

	var r v1.IdRequest
	if err := c.ShouldBind(&r); err != nil {
		core.Error(c, errno.ErrBind)
		return
	}

	validate := validator.New()
	if errs := validate.Struct(r); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			core.Error(c, errno.ErrInvalidParameter.SetMessage(err.Error()))
			return
		}
	}

	err := ctrl.b.Areas().Enable(c, r.Id)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, nil, "启用区域信息成功")
}

// Disable 禁用区域信息
func (ctrl Controller) Disable(c *gin.Context) {
	log.C(c).Info("system attributes area disable function called")

	var r v1.IdRequest
	if err := c.ShouldBind(&r); err != nil {
		core.Error(c, errno.ErrBind)
		return
	}

	validate := validator.New()
	if errs := validate.Struct(r); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			core.Error(c, errno.ErrInvalidParameter.SetMessage(err.Error()))
			return
		}
	}

	err := ctrl.b.Areas().Disable(c, r.Id)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, nil, "禁用区域信息成功")
}
