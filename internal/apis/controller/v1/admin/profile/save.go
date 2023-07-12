// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package profile

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	"apis/internal/pkg/core"
	"apis/internal/pkg/errno"
	v1 "apis/internal/pkg/request/apis/v1"
	"apis/pkg/log"
)

// Save 更新账户信息
func (ctrl ProfileController) Save(c *gin.Context) {
	log.C(c).Info("profile save function called")

	var r v1.ProfileRequest
	if err := c.ShouldBind(&r); err != nil {
		core.Error(c, errno.ErrBind)
		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.Error(c, errno.ErrInvalidParameter.SetMessage(err.Error()))
		return
	}

	if err := ctrl.b.Profiles().Save(c, &r); err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, nil, "update profile info success")
}
