// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package role

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

type Controller struct {
	b biz.IBiz
}

func New(ds store.IStore) *Controller {
	return &Controller{
		b: biz.NewBiz(ds),
	}
}

func (ctrl *Controller) Lists(c *gin.Context) {
	log.C(c).Info("admin role lists function called")

	var r v1.PageRequest
	if err := c.ShouldBind(&r); err != nil {
		core.Error(c, err)
		return
	}

	validate := validator.New()
	if errs := validate.Struct(r); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			core.Error(c, errno.ErrInvalidParameter.SetMessage(err.Error()))
			return
		}
	}

	lists, total, err := ctrl.b.AdminGroups().Lists(c, r)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, gin.H{
		"current_page": r.Page,
		"per_page":     r.PageSize,
		"data":         lists,
		"total":        total,
	}, "get admin role list success")
}

func (ctrl *Controller) Detail(c *gin.Context) {
	log.C(c).Info("admin role detail function called")

	var r v1.IdRequest
	if err := c.ShouldBind(&r); err != nil {
		core.Error(c, err)
		return
	}

	validate := validator.New()
	if errs := validate.Struct(r); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			core.Error(c, errno.ErrInvalidParameter.SetMessage(err.Error()))
			return
		}
	}

	group, err := ctrl.b.AdminGroups().Detail(c, r.Id)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, group, "get admin role detail success")
}

func (ctrl *Controller) Update(c *gin.Context) {
	log.C(c).Info("admin role update function called")

	var r v1.AdminGroupUpdateRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.Error(c, err)
		return
	}

	validate := validator.New()
	if errs := validate.Struct(r); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			core.Error(c, errno.ErrInvalidParameter.SetMessage(err.Error()))
			return
		}
	}

	if err := ctrl.b.AdminGroups().Update(c, r); err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, nil, "admin role update success")
}

func (ctrl *Controller) Enable(c *gin.Context) {
	log.C(c).Info("admin role enable function called")

	var r v1.IdRequest
	if err := c.ShouldBind(&r); err != nil {
		core.Error(c, err)
		return
	}

	validate := validator.New()
	if errs := validate.Struct(r); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			core.Error(c, errno.ErrInvalidParameter.SetMessage(err.Error()))
			return
		}
	}

	if err := ctrl.b.AdminGroups().Enable(c, r.Id); err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, nil, "admin role enable success")
}

func (ctrl *Controller) Disable(c *gin.Context) {
	log.C(c).Info("admin role disable function called")

	var r v1.IdRequest
	if err := c.ShouldBind(&r); err != nil {
		core.Error(c, err)
		return
	}

	validate := validator.New()
	if errs := validate.Struct(r); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			core.Error(c, errno.ErrInvalidParameter.SetMessage(err.Error()))
			return
		}
	}

	if err := ctrl.b.AdminGroups().Disable(c, r.Id); err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, nil, "admin role disable success")
}

func (ctrl *Controller) Rules(c *gin.Context) {
	log.C(c).Info("admin role rules function called")

	rules, err := ctrl.b.Menus().Rules(c)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, rules, "admin role rules success")
}
