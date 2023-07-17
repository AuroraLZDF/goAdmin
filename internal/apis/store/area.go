// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package store

import (
	"gorm.io/gorm"

	"apis/internal/pkg/model"
	"apis/internal/pkg/model/common"
	v1 "apis/internal/pkg/request/apis/v1"
)

// AreaStore 定义了 area 模块在 store 层所实现的方法.
type AreaStore interface {
	Gets(pid int) (*[]common.Areas, error)
	Get(id int) (*common.Areas, error)
	HasSameArea(pid int, name string, id int) bool
	CreateOrUpdate(request v1.AreaUpdateRequest) (int, error)
	Enable(area *common.Areas) error
	Disable(area *common.Areas) error
}

// AreaStore 接口的实现.
type areas struct {
	db *gorm.DB
}

// 确保 areas 实现了 AreaStore 接口.
var _ AreaStore = (*areas)(nil)

func newAreas(db *gorm.DB) *areas {
	return &areas{db}
}

// Gets 获取 指定 pid area 列表
func (a *areas) Gets(pid int) (*[]common.Areas, error) {
	var lists []common.Areas

	err := a.db.Where("pid=?", pid).Order("id asc").Find(&lists).Error
	if err != nil {
		return nil, err
	}

	return &lists, nil
}

// Get 获取 指定 id area 详情
func (a *areas) Get(id int) (*common.Areas, error) {
	var info common.Areas

	err := a.db.Where("id=?", id).First(&info).Error
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// HasSameArea 判断是否存在相同的区域信息
func (a *areas) HasSameArea(pid int, name string, id int) bool {
	var info common.Areas

	res := a.db.Where("pid=?", pid).Where("name=?", name)
	if id > 0 {
		res = res.Where("id!=?", id)
	}

	if err := res.First(&info).Error; err != nil {
		return false
	}

	if id <= 0 || (id > 0 && info.ID != id) {
		return true
	}

	return false
}

// CreateOrUpdate 创建或更新区域信息
func (a *areas) CreateOrUpdate(r v1.AreaUpdateRequest) (int, error) {
	var info common.Areas

	if r.Id != 0 {
		if err := a.db.Where("id=?", r.Id).First(&info).Error; err != nil {
			return 0, err
		}

		info.Pid = r.Pid
		info.Name = r.Name
		info.Level = r.Level
		if err := a.db.Updates(&info).Error; err != nil {
			return 0, err
		}
	} else {
		if err := a.db.Create(r).Error; err != nil {
			return 0, err
		}
		a.db.Last(&info)
	}

	return info.ID, nil
}

// Enable 启用区域信息
func (a *areas) Enable(area *common.Areas) error {
	area.Status = model.StatusOn
	if err := a.db.Save(&area).Error; err != nil {
		return err
	}
	return nil
}

// Disable 禁用区域信息
func (a *areas) Disable(area *common.Areas) error {
	area.Status = model.StatusOff
	if err := a.db.Save(&area).Error; err != nil {
		return err
	}
	return nil
}
