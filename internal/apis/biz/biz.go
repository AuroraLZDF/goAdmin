// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package biz

import (
	"apis/internal/apis/biz/admin/admin"
	"apis/internal/apis/biz/admin/profile"

	"apis/internal/apis/store"
)

// IBiz 定义了 Biz 层需要实现的方法.
type IBiz interface {
	Admins() admin.AdminBiz
	Profiles() profile.ProfileBiz
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
