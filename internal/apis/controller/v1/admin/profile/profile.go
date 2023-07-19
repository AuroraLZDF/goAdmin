// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package profile

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

// Controller 是 profile 模块在 Controller 层的实现，用来处理个人资料模块的请求.
type Controller struct {
	b biz.IBiz
}

// New 创建一个 profile controller.
func New(ds store.IStore) *Controller {
	return &Controller{
		b: biz.NewBiz(ds),
	}
}

// Save 更新账户信息
func (ctrl Controller) Save(c *gin.Context) {
	log.C(c).Info("profile save function called")

	var r v1.ProfileRequest
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

	if err := ctrl.b.Profiles().Save(c, &r); err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, nil, "update profile info success")
}
