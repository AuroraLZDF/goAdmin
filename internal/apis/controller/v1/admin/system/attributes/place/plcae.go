// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package place

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

// PlaceController 是 place 模块在 Controller 层的实现，用来处理位置信息的请求.
type PlaceController struct {
	b biz.IBiz
}

// New 创建一个 place controller.
func New(ds store.IStore) *PlaceController {
	return &PlaceController{
		b: biz.NewBiz(ds),
	}
}

func (ctrl *PlaceController) Lists(c *gin.Context) {
	log.C(c).Info("system attributes place lists function called")

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

	lists, err := ctrl.b.Places().Lists(c, r)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, lists, "get place lists success")
}

func (ctrl *PlaceController) Detail(c *gin.Context) {
	log.C(c).Info("system attributes place detail function called")

	var r v1.PlaceRequest
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

	info, err := ctrl.b.Places().Detail(c, r.Id)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, info, "get place detail success")
}

func (ctrl *PlaceController) Update(c *gin.Context) {
	log.C(c).Info("system attributes place update function called")

	var r v1.PlaceUpdateRequest
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

	err := ctrl.b.Places().CreateOrUpdate(c, r)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, nil, "update place info success")
}
