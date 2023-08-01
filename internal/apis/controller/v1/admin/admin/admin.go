// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package admin

import (
	"apis/internal/apis/biz"
	"apis/internal/apis/store"
	"apis/internal/pkg/core"
	"apis/internal/pkg/errno"
	v1 "apis/internal/pkg/request/apis/v1"
	"apis/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Controller struct {
	b biz.IBiz
}

func New(ds store.IStore) *Controller {
	return &Controller{
		b: biz.NewBiz(ds),
	}
}

func (ctrl Controller) Lists(c *gin.Context) {
	log.C(c).Info("admin admin lists function called")

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

	lists, total, err := ctrl.b.Admins().Lists(c, r)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, gin.H{
		"current_page": r.Page,
		"per_page":     r.PageSize,
		"data":         lists,
		"total":        total,
	}, "get admin lists success")
}

func (ctrl Controller) Detail(c *gin.Context) {
	log.C(c).Info("admin admin detail function called")

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

	info, err := ctrl.b.Admins().Detail(c, r.Id)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, info, "get admin detail success")
}

func (ctrl Controller) Update(c *gin.Context) {
	log.C(c).Info("admin admin update function called")

	var r v1.AdminUpdateRequest
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

	err := ctrl.b.Admins().Update(c, r)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, nil, "update admin success")
}

func (ctrl Controller) Enable(c *gin.Context) {
	log.C(c).Info("admin admin enable function called")

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

	err := ctrl.b.Admins().Enable(c, r.Id)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, nil, "enable admin success")
}

func (ctrl Controller) Disable(c *gin.Context) {
	log.C(c).Info("admin admin disable function called")

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

	err := ctrl.b.Admins().Disable(c, r.Id)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, nil, "disable admin success")
}
