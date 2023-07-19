// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package store

import (
	"apis/internal/pkg/model"
	"gorm.io/gorm"

	"apis/internal/pkg/model/admin"
	v1 "apis/internal/pkg/request/apis/v1"
)

// GeneralStore 定义了 general 模块在 store 层所实现的方法
type GeneralStore interface {
	Count() int
	Gets(r v1.PageRequest) (*[]admin.GeneralConfigs, error)
	Get(id int) (*admin.GeneralConfigs, error)
	CreateOrUpdate(r v1.GeneralUpdateRequest) error
	Enable(id int) error
	Disable(id int) error
}

// GeneralStore 接口的实现.
type generals struct {
	db *gorm.DB
}

// 确保 places 实现了 PlaceStore 接口.
var _ PlaceStore = (*places)(nil)

func newGeneral(db *gorm.DB) *generals {
	return &generals{db}
}

func (c *generals) Count() int {
	var count int64
	if err := c.db.Model(&admin.GeneralConfigs{}).Count(&count).Error; err != nil {
		return 0
	}
	return int(count)
}

func (g *generals) Gets(r v1.PageRequest) (*[]admin.GeneralConfigs, error) {
	var lists []admin.GeneralConfigs

	err := g.db.Limit(r.PageSize).Offset((r.Page - 1) * r.PageSize).Order("id desc").Find(&lists).Error
	if err != nil {
		return nil, err
	}

	return &lists, nil
}

func (g *generals) Get(id int) (*admin.GeneralConfigs, error) {
	var info admin.GeneralConfigs

	err := g.db.Where("id=?", id).First(&info).Error
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func (g *generals) CreateOrUpdate(r v1.GeneralUpdateRequest) error {
	if r.Id != 0 {
		var info admin.GeneralConfigs

		if err := g.db.Where("id=?", r.Id).First(&info).Error; err != nil {
			return err
		}

		info.Key = r.Key
		info.Status = r.Status
		if err := g.db.Save(&info).Error; err != nil {
			return err
		}
	} else {
		newGeneral := admin.GeneralConfigs{
			Key:    r.Key,
			Status: r.Status,
		}
		if err := g.db.Create(&newGeneral).Error; err != nil {
			return err
		}
	}
	return nil
}

func (g *generals) Enable(id int) error {
	var info admin.GeneralConfigs

	err := g.db.Where("id=?", id).First(&info).Error
	if err != nil {
		return err
	}

	info.Status = model.StatusOn
	if err := g.db.Save(&info).Error; err != nil {
		return err
	}
	return nil
}

func (g *generals) Disable(id int) error {
	var info admin.GeneralConfigs

	err := g.db.Where("id=?", id).First(&info).Error
	if err != nil {
		return err
	}

	info.Status = model.StatusOff
	if err := g.db.Save(&info).Error; err != nil {
		return err
	}
	return nil
}
