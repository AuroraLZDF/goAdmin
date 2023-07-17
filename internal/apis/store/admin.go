// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package store

import (
	"gorm.io/gorm"

	"apis/internal/pkg/model/admin"
)

// AdminStore 定义了 auth 模块在 store 层所实现的方法.
type AdminStore interface {
	Create(user *admin.Admins) error
	Get(phone string) (*admin.Admins, error)
	Update(user *admin.Admins) error
}

// AdminStore 接口的实现.
type admins struct {
	db *gorm.DB
}

// 确保 admins 实现了 AdminStore 接口.
var _ AdminStore = (*admins)(nil)

func newAdmins(db *gorm.DB) *admins {
	return &admins{db}
}

// Create 插入一条 auth 记录.
func (u *admins) Create(user *admin.Admins) error {
	return u.db.Create(&user).Error
}

// Get 获取一条用户数据
func (u *admins) Get(phone string) (*admin.Admins, error) {
	var user admin.Admins
	if err := u.db.Where("phone = ?", phone).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Update 更新账户信息
func (u *admins) Update(user *admin.Admins) error {
	return u.db.Updates(&user).Error
}
