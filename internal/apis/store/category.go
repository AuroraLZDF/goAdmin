// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package store

import (
	"apis/internal/pkg/model"
	"gorm.io/gorm"

	"apis/internal/pkg/model/common"
	v1 "apis/internal/pkg/request/apis/v1"
)

// CategoryStore 定义了 category 模块在 store 层所实现的方法.
type CategoryStore interface {
	Gets(r v1.PageRequest) (*[]common.Categories, error)
	Get(id int) (*common.Categories, error)
	CreateOrUpdate(r v1.CategoryUpdateRequest) error
	Enable(*common.Categories) error
	Disable(*common.Categories) error
}

// CategoryStore 接口的实现.
type categories struct {
	db *gorm.DB
}

// 确保 categories 实现了 CategoryStore 接口.
var _ CategoryStore = (*categories)(nil)

func newCategories(db *gorm.DB) *categories {
	return &categories{db}
}

func (c *categories) Gets(r v1.PageRequest) (*[]common.Categories, error) {
	var lists []common.Categories

	err := c.db.Limit(r.PageSize).Offset((r.Page - 1) * r.PageSize).Order("id desc").Find(&lists).Error
	if err != nil {
		return nil, err
	}

	return &lists, nil
}

func (c *categories) Get(id int) (*common.Categories, error) {
	var info common.Categories

	err := c.db.Where("id=?", id).First(&info).Error
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func (c *categories) CreateOrUpdate(r v1.CategoryUpdateRequest) error {
	if r.Id != 0 {
		var info common.Categories

		if err := c.db.Where("id=?", r.Id).First(&info).Error; err != nil {
			return err
		}

		info.Icon = r.Icon
		info.Name = r.Name
		info.Cover = r.Cover
		info.CoverMobile = r.CoverMobile
		if err := c.db.Save(&info).Error; err != nil {
			return err
		}
	} else {
		newCategory := common.Categories{
			Name:        r.Name,
			Icon:        r.Icon,
			Cover:       r.Cover,
			CoverMobile: r.CoverMobile,
		}
		if err := c.db.Create(&newCategory).Error; err != nil {
			return err
		}
	}
	return nil
}

func (c *categories) Enable(category *common.Categories) error {
	category.IsHidden = model.StatusOn
	if err := c.db.Save(&category).Error; err != nil {
		return err
	}
	return nil
}

func (c *categories) Disable(category *common.Categories) error {
	category.IsHidden = model.StatusOff
	if err := c.db.Save(&category).Error; err != nil {
		return err
	}
	return nil
}
