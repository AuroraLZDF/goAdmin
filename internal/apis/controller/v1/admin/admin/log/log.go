// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package log

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
	log *admin.AdminLogs
}

func New(db *gorm.DB) *Controller {
	return &Controller{
		log: admin.NewAdminLogs(db),
	}
}

func (ctrl Controller) Lists(c *gin.Context) {
	log.C(c).Info("admin log lists function called")

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

	lists, total, err := ctrl.log.Lists(r)
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, gin.H{
		"current_page": r.Page,
		"per_page":     r.PageSize,
		"data":         lists,
		"total":        total,
	}, "get admin log lists success")
}
