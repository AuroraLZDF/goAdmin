// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package common

import "apis/internal/pkg/model"

type Areas struct {
	model.Model
	Name        string `gorm:"column:name" json:"name"`     //省份、城市、区县、街道等名称
	Pid         int    `gorm:"column:pid" json:"pid"`       //父级ID
	Level       int    `gorm:"column:level" json:"level"`   //层级
	Status      int    `gorm:"column:status" json:"status"` //是否启用 0 启用 1 禁用
	Leaf        bool   `gorm:"-" json:"leaf"`
	HasChildren bool   `gorm:"-" json:"hasChildren"`
}

// TableName sets the insert table name for this struct type
func (a *Areas) TableName() string {
	return "areas"
}
