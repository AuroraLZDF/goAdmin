// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package store

import (
	"apis/internal/pkg/model/admin"
	v1 "apis/internal/pkg/request/apis/v1"
	"apis/pkg/util"
	"gorm.io/gorm"
)

// ConfigStore 定义了 config 模块在 store 中的接口
type ConfigStore interface {
	// Get returns the value for a key.
	Get() (*[]admin.Configs, error)
	// Update sets the value for a key.
	Update(r v1.ConfigUpdateRequest) error
}

// ConfigStore 接口的实现
type configs struct {
	db *gorm.DB
}

// 确保 configs 实现了 ConfigStore 接口
var _ ConfigStore = (*configs)(nil)

// NewConfigs 创建一个 configs 实现类
func NewConfigs(db *gorm.DB) *configs {
	return &configs{db}
}

// Get 获取 configs 中 config 的值。
func (c *configs) Get() (*[]admin.Configs, error) {
	var lists []admin.Configs
	if err := c.db.Find(&lists).Error; err != nil {
		return nil, err
	}

	if len(lists) == 0 {
		return nil, nil
	}

	return &lists, nil
}

// Update Save configs 的值。
func (c *configs) Update(r v1.ConfigUpdateRequest) error {
	var request = util.StructToMap(r)

	for key, value := range request {
		var config admin.Configs
		err := c.db.Where("k=?", key).First(&config).Error
		if err != nil {
			return err
		}
		config.V = value
		if err := c.db.Save(&config).Error; err != nil {
			return err
		}
	}

	return nil
}
