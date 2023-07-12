// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package profile

import (
	"apis/internal/apis/biz"
	"apis/internal/apis/store"
)

// ProfileController 是 profile 模块在 Controller 层的实现，用来处理个人资料模块的请求.
type ProfileController struct {
	b biz.IBiz
}

// New 创建一个 profile controller.
func New(ds store.IStore) *ProfileController {
	return &ProfileController{
		b: biz.NewBiz(ds),
	}
}
