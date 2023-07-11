// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package auth

import (
	"apis/internal/apis/biz"
	"apis/internal/apis/store"
)

// AdminController 是 admin 模块在 Controller 层的实现，用来处理用户模块的请求.
type AdminController struct {
	b biz.IBiz
}

// New 创建一个 admin controller.
func New(ds store.IStore) *AdminController {
	return &AdminController{
		b: biz.NewBiz(ds),
	}
}
