// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package category

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

// Controller 是 category 模块在 Controller 层的实现，用来处理分类信息的请求.
type Controller struct {
	b biz.IBiz
}

// New 创建一个 category controller.
func New(ds store.IStore) *Controller {
	return &Controller{
		b: biz.NewBiz(ds),
	}
}

func (ctrl Controller) Lists(c *gin.Context) {
	log.C(c).Info("system attributes category lists function called")

	var r v1.PageRequest
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

	count := ctrl.b.Categories().Count(c)
	lists, err := ctrl.b.Categories().Lists(c, r)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, gin.H{
		"current_page": r.Page,
		"per_page":     r.PageSize,
		"data":         lists,
		"total":        count,
	}, "get category lists success")
}

func (ctrl Controller) Detail(c *gin.Context) {
	log.C(c).Info("system attributes category detail function called")

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

	info, err := ctrl.b.Categories().Detail(c, r.Id)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, info, "get category detail success")
}

// Update 更新区域信息
func (ctrl Controller) Update(c *gin.Context) {
	log.C(c).Info("system attributes category update function called")

	var r v1.CategoryUpdateRequest
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

	err := ctrl.b.Categories().CreateOrUpdate(c, r)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, nil, "update category info success")
}

// Enable 启用区域信息
func (ctrl Controller) Enable(c *gin.Context) {
	log.C(c).Info("system attributes category enable function called")

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

	err := ctrl.b.Categories().Enable(c, r.Id)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, nil, "enable category success")
}

// Disable 启用区域信息
func (ctrl Controller) Disable(c *gin.Context) {
	log.C(c).Info("system attributes category disable function called")

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

	err := ctrl.b.Categories().Disable(c, r.Id)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, nil, "disable category success")
}
