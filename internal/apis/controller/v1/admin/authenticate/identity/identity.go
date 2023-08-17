// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package identity

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	"apis/internal/pkg/core"
	"apis/internal/pkg/errno"
	"apis/internal/pkg/model/admin"
	v1 "apis/internal/pkg/request/apis/v1"
	"apis/pkg/log"
)

type Controller struct {
	identity *admin.Identities
}

func New(db *gorm.DB) *Controller {
	return &Controller{
		identity: admin.NewIdentities(db),
	}
}

func (ctrl *Controller) Lists(c *gin.Context) {
	log.C(c).Info("identity lists function called")

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

	lists, total, err := ctrl.identity.Lists(r)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, gin.H{
		"current_page": r.Page,
		"per_page":     r.PageSize,
		"data":         lists,
		"total":        total,
	}, "get identity lists success")
}

func (ctrl *Controller) Detail(c *gin.Context) {
	log.C(c).Info("identity detail function called")

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

	info, err := ctrl.identity.Detail(r)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, info, "get identity detail info success")
}

func (ctrl *Controller) Pass(c *gin.Context) {
	log.C(c).Info("identity pass function called")

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

	err := ctrl.identity.Pass(r)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, nil, "Passed identity info success")
}

func (ctrl *Controller) Reject(c *gin.Context) {
	log.C(c).Info("identity reject function called")

	var r v1.IdentityRequest
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

	err := ctrl.identity.Reject(r)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, nil, "rejected identity info success")
}
