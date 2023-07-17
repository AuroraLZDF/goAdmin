// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package area

import (
	"apis/internal/apis/biz"
	"apis/internal/apis/store"
)

// AreaController 是 area 模块在 Controller 层的实现，用来处理区域信息的请求.
type AreaController struct {
	b biz.IBiz
}

// New 创建一个 area controller.
func New(ds store.IStore) *AreaController {
	return &AreaController{
		b: biz.NewBiz(ds),
	}
}
