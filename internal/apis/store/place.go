// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package store

import (
	"gorm.io/gorm"

	"apis/internal/pkg/model/admin"
	v1 "apis/internal/pkg/request/apis/v1"
)

// PlaceStore 定义了 place 模块在 store 层所实现的方法
type PlaceStore interface {
	Count() int
	Gets(r v1.PageRequest) (*[]admin.Places, error)
	Get(id int) (*admin.Places, error)
	CreateOrUpdate(r v1.PlaceUpdateRequest) error
}

// PlaceStore 接口的实现.
type places struct {
	db *gorm.DB
}

// 确保 places 实现了 PlaceStore 接口.
var _ PlaceStore = (*places)(nil)

func newPlace(db *gorm.DB) *places {
	return &places{db}
}

func (c *places) Count() int {
	var count int64
	if err := c.db.Model(&admin.Places{}).Count(&count).Error; err != nil {
		return 0
	}
	return int(count)
}

func (p *places) Gets(r v1.PageRequest) (*[]admin.Places, error) {
	var lists []admin.Places

	err := p.db.Limit(r.PageSize).Offset((r.Page - 1) * r.PageSize).Order("id desc").Find(&lists).Error
	if err != nil {
		return nil, err
	}

	return &lists, nil
}

func (p *places) Get(id int) (*admin.Places, error) {
	var info admin.Places

	err := p.db.Where("id=?", id).First(&info).Error
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func (p *places) CreateOrUpdate(r v1.PlaceUpdateRequest) error {
	if r.Id != 0 {
		var info admin.Places

		if err := p.db.Where("id=?", r.Id).First(&info).Error; err != nil {
			return err
		}

		info.Name = r.Name
		if err := p.db.Save(&info).Error; err != nil {
			return err
		}
	} else {
		newPlace := admin.Places{Name: r.Name}
		if err := p.db.Create(&newPlace).Error; err != nil {
			return err
		}
	}
	return nil
}
