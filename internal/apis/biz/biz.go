// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package biz

import (
	"apis/internal/apis/biz/admin/access"
	"apis/internal/apis/biz/admin/admin"
	"apis/internal/apis/biz/admin/group"
	"apis/internal/apis/biz/admin/profile"
	"apis/internal/apis/biz/admin/system/area"
	"apis/internal/apis/biz/admin/system/category"
	"apis/internal/apis/biz/admin/system/config"
	"apis/internal/apis/biz/admin/system/general"
	"apis/internal/apis/biz/admin/system/menu"
	"apis/internal/apis/biz/admin/system/place"
	"apis/internal/apis/store"
)

// IBiz 定义了 Biz 层需要实现的方法.
type IBiz interface {
	Admins() admin.AdminBiz
	Profiles() profile.ProfileBiz
	Areas() area.AreaBiz
	Categories() category.CategoryBiz
	Places() place.PlaceBiz
	Generals() general.GeneralBiz
	Configs() config.ConfigBiz
	Menus() menu.MenuBiz
	AdminGroups() group.AdminGroupBiz
	AdminAccesses() access.AdminAccessBiz
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

// Places 返回一个实现了 PlaceBiz 接口的实例.
func (b *biz) Places() place.PlaceBiz {
	return place.New(b.ds)
}

// Generals 返回一个实现了 GeneralBiz 接口的实例.
func (b *biz) Generals() general.GeneralBiz {
	return general.New(b.ds)
}

// Configs  返回一个实现了 ConfigBiz 接口的实例.
func (b *biz) Configs() config.ConfigBiz {
	return config.New(b.ds)
}

// Menus 返回一个实现了 MenuBiz 接口的实例.
func (b *biz) Menus() menu.MenuBiz {
	return menu.New(b.ds)
}

func (b *biz) AdminGroups() group.AdminGroupBiz {
	return group.New(b.ds)
}

func (b *biz) AdminAccesses() access.AdminAccessBiz {
	return access.New(b.ds)
}
