// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package group

import (
	"apis/internal/pkg/core"
	"apis/internal/pkg/errno"
	"apis/internal/pkg/model"
	v1 "apis/internal/pkg/request/apis/v1"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"apis/internal/apis/biz"
	"apis/internal/apis/store"
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

func (ctrl *Controller) Index(c *gin.Context) {
	log.C(c).Info("admin group index function called")

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

	user, err := ctrl.b.Admins().Detail(c, r.Id)
	if err != nil {
		core.Error(c, err)
		return
	}

	groups, err := ctrl.b.AdminGroups().Get(c, model.StatusOn)
	if err != nil {
		core.Error(c, err)
		return
	}

	accesses, err := ctrl.b.AdminAccesses().UserRoles(c, r.Id)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, gin.H{
		"info":     user,
		"groups":   groups,
		"accesses": accesses,
	}, "get group info success")
}

func (ctrl *Controller) Update(c *gin.Context) {
	log.C(c).Info("admin group update function called")

	var r v1.AdminAccessUpdateRequest
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

	err := ctrl.b.AdminAccesses().Update(c, r)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, nil, "update group success")
}
