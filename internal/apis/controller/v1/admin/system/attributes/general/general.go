// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package general

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
	log.C(c).Info("system attributes general config lists function called")

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

	lists, err := ctrl.b.Generals().Lists(c, r)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, lists, "get general config lists success")
}

func (ctrl *Controller) Detail(c *gin.Context) {
	log.C(c).Info("system attributes general config detail function called")

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

	info, err := ctrl.b.Generals().Detail(c, r.Id)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, info, "get general config detail success")
}

func (ctrl *Controller) Save(c *gin.Context) {
	log.C(c).Info("system attributes general config update function called")

	var r v1.GeneralUpdateRequest
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

	err := ctrl.b.Generals().CreateOrUpdate(c, r)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, nil, "update general config info success")
}

func (ctrl *Controller) Enable(c *gin.Context) {
	log.C(c).Info("system attributes general config enable function called")

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

	if err := ctrl.b.Generals().Enable(c, r.Id); err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, nil, "enable general config info success")
}

func (ctrl *Controller) Disable(c *gin.Context) {
	log.C(c).Info("system attributes general config disable function called")

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

	if err := ctrl.b.Generals().Disable(c, r.Id); err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, nil, "disable general config info success")
}
