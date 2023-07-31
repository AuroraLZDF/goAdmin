// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package store

import (
	"apis/internal/pkg/model/admin"
	"gorm.io/gorm"
)

// AdminAccessStore 定义用于存储管理员访问权限的接口
type AdminAccessStore interface {
	GetAdminRoleMenus(uid int) ([]int, error)
}

// AdminStore 接口的实现.
type adminAccesses struct {
	db *gorm.DB
}

// 确保 adminAccesses 实现了 AdminStore 接口.
var _ AdminAccessStore = (*adminAccesses)(nil)

func NewAdminAccesses(db *gorm.DB) *adminAccesses {
	return &adminAccesses{db}
}

func (a *adminAccesses) GetAdminRoleMenus(uid int) ([]int, error) {
	var res []int
	var menus []admin.AdminAccess
	if err := a.db.Model(admin.AdminAccess{}).Find(&menus, "uid =?", uid).Error; err != nil {
		return nil, err
	}

	if len(menus) > 0 {
		for _, menu := range menus {
			res = append(res, menu.GroupID)
		}
	}

	return res, nil
}
