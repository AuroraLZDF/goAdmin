// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package category

import (
	"apis/internal/apis/biz"
	"apis/internal/apis/store"
)

// CategoryController 是 category 模块在 Controller 层的实现，用来处理分类信息的请求.
type CategoryController struct {
	b biz.IBiz
}

// New 创建一个 category controller.
func New(ds store.IStore) *CategoryController {
	return &CategoryController{
		b: biz.NewBiz(ds),
	}
}
