// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package area

import (
	"github.com/gin-gonic/gin"

	"apis/internal/pkg/core"
	"apis/pkg/log"
)

// Lists 区域信息列表
func (ctrl AreaController) Lists(c *gin.Context) {
	log.C(c).Info("system attributes area lists function called")

	lists, err := ctrl.b.Areas().Lists(c, c.GetInt("pid"))
	if err != nil {
		core.Error(c, err)
		return
	}

	core.Success(c, lists, "get area lists success")
}
