// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package config

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

// Detail 站点设置详情
func (ctrl Controller) Detail(c *gin.Context) {
	log.C(c).Info("system config detail function called")

	info, err := ctrl.b.Configs().Detail(c)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, info, "get config detail success")
}

func (ctrl Controller) Save(c *gin.Context) {
	log.C(c).Info("system config save function called")

	var r v1.ConfigUpdateRequest
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

	err := ctrl.b.Configs().Update(c, r)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, nil, "save system config success")
}
