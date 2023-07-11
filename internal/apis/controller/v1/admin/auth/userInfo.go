// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package auth

import (
	"github.com/gin-gonic/gin"

	"apis/internal/pkg/core"
	"apis/pkg/log"
)

func (ctrl *AdminController) UserInfo(c *gin.Context) {
	log.C(c).Info("user detail info function called")

	info, err := ctrl.b.Admins().Info(c, c.GetString("phone"))
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, info, "success.")
}
