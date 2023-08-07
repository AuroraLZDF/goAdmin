// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package store

import (
	"apis/internal/pkg/model/admin"
	v1 "apis/internal/pkg/request/apis/v1"
	"gorm.io/gorm"
)

// AdminAccessStore 定义用于存储管理员访问权限的接口
type AdminAccessStore interface {
	GetAdminRoleMenus(uid int) ([]int, error)
	GetUserRoles(uid int) ([]int, error)
	Update(r v1.AdminAccessUpdateRequest) error
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

func (a *adminAccesses) GetUserRoles(uid int) ([]int, error) {
	var res []int
	var roles []admin.AdminAccess
	if err := a.db.Model(admin.AdminAccess{}).Find(&roles, "uid =?", uid).Error; err != nil {
		return nil, err
	}

	if len(roles) > 0 {
		for _, role := range roles {
			res = append(res, role.GroupID)
		}
	}

	return res, nil
}

func (a *adminAccesses) Update(r v1.AdminAccessUpdateRequest) error {
	err := a.db.Transaction(func(tx *gorm.DB) error {
		// 执行数据库操作
		if err := tx.Where("uid = ?", r.Uid).Delete(admin.AdminAccess{}).Error; err != nil {
			return err
		}

		if len(r.Rules) > 0 {
			for _, rule := range r.Rules {
				if err := tx.Create(&admin.AdminAccess{
					Uid:     r.Uid,
					GroupID: rule,
				}).Error; err != nil {
					return err
				}
			}
		}

		return nil // 返回 nil 表示提交事务，返回错误表示回滚事务
	})
	return err
}
