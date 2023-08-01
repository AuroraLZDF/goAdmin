// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package store

import (
	"gorm.io/gorm"

	"apis/internal/pkg/model"
	"apis/internal/pkg/model/admin"
	v1 "apis/internal/pkg/request/apis/v1"
)

// AdminStore 定义了 auth 模块在 store 层所实现的方法.
type AdminStore interface {
	Create(r v1.AdminUpdateRequest) error
	Get(phone string) (*admin.Admins, error)
	Update(user *admin.Admins) error
	Gets(r v1.PageRequest) (*[]admin.Admins, int, error)
	InfoById(id int) (*admin.Admins, error)
	Enable(id int) error
	Disable(id int) error
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
func (u *admins) Create(r v1.AdminUpdateRequest) error {
	return u.db.Model(&admin.Admins{}).Create(&r).Error
}

func (u *admins) Gets(r v1.PageRequest) (*[]admin.Admins, int, error) {
	var users []admin.Admins
	var total int64
	if err := u.db.Model(&admin.Admins{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := u.db.Limit(r.PageSize).Offset((r.Page - 1) * r.PageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return &users, int(total), nil
}

// Get 获取一条用户数据
func (u *admins) Get(phone string) (*admin.Admins, error) {
	var user admin.Admins
	if err := u.db.Where("phone = ?", phone).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *admins) InfoById(id int) (*admin.Admins, error) {
	var user admin.Admins
	if err := u.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Update 更新账户信息
func (u *admins) Update(user *admin.Admins) error {
	return u.db.Updates(&user).Error
}

func (u *admins) Enable(id int) error {
	var user admin.Admins
	if err := u.db.Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}

	user.Status = model.StatusOn
	return u.db.Save(&user).Error
}

func (u *admins) Disable(id int) error {
	var user admin.Admins
	if err := u.db.Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}

	user.Status = model.StatusOff
	return u.db.Save(&user).Error
}
