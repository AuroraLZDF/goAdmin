// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package menu

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
	log.C(c).Info("system menu list function called")

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

	//count := ctrl.b.Menus().Count(c)
	lists, err := ctrl.b.Menus().Lists(c, r)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, lists, "get menu lists success")
}

func (ctrl *Controller) Detail(c *gin.Context) {
	log.C(c).Info("system menu detail function called")

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

	info, err := ctrl.b.Menus().Detail(c, r.Id)
	if err != nil {
		core.Error(c, err)
		return
	}

	_type := c.GetString("type")
	selectOptions, err := ctrl.b.Menus().GetPid(info.Id)
	if err != nil {
		core.Error(c, err)
		return
	}

	if _type == "add" {
		selectOptions = append(selectOptions, info.Id)
	}

	core.Success(c, gin.H{
		"info":          info,
		"selectOptions": selectOptions,
	}, "get menu detail success")
}

func (ctrl *Controller) Update(c *gin.Context) {
	log.C(c).Info("system menu update function called")

	var r v1.MenuUpdateRequest
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

	err := ctrl.b.Menus().CreateOrUpdate(c, r)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, nil, "update menu success")
}

func (ctrl *Controller) Enable(c *gin.Context) {
	log.C(c).Info("system menu enable function called")

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

	err := ctrl.b.Menus().Enable(c, r.Id)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, nil, "enable menu success")
}

func (ctrl *Controller) Disable(c *gin.Context) {
	log.C(c).Info("system menu disable function called")

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

	err := ctrl.b.Menus().Disable(c, r.Id)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, nil, "disable menu success")
}

func (ctrl *Controller) RoleMenu(c *gin.Context) {
	log.C(c).Info(`system menu roleMenu function called`)

	menus, err := ctrl.b.Menus().RoleMenu(c, c.GetString("phone"))
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, menus, "get system menu roleMenu success")
}
