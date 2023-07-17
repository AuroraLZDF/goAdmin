// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package biz

import (
	"apis/internal/apis/biz/admin/admin"
	"apis/internal/apis/biz/admin/profile"
	"apis/internal/apis/biz/admin/system/area"
	"apis/internal/apis/biz/admin/system/category"

	"apis/internal/apis/store"
)

// IBiz 定义了 Biz 层需要实现的方法.
type IBiz interface {
	Admins() admin.AdminBiz
	Profiles() profile.ProfileBiz
	Areas() area.AreaBiz
	Categories() category.CategoryBiz
}

// biz 是 IBiz 的一个具体实现.
type biz struct {
	ds store.IStore
}

// 确保 biz 实现了 IBiz 接口.
var _ IBiz = (*biz)(nil)

// NewBiz 创建一个 IBiz 类型的实例.
func NewBiz(ds store.IStore) *biz {
	return &biz{ds: ds}
}

// Admins 返回一个实现了 AdminBiz 接口的实例.
func (b *biz) Admins() admin.AdminBiz {
	return admin.New(b.ds)
}

// Profiles 返回一个实现了 ProfileBiz 接口的实例.
func (b *biz) Profiles() profile.ProfileBiz {
	return profile.New(b.ds)
}

// Areas 返回一个实现了 AreaBiz 接口的实例.
func (b *biz) Areas() area.AreaBiz {
	return area.New(b.ds)
}

// Categories 返回一个实现了 CategoryBiz 接口的实例.
func (b *biz) Categories() category.CategoryBiz {
	return category.New(b.ds)
}
