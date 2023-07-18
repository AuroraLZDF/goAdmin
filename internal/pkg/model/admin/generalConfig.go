// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package admin

import "apis/internal/pkg/model"

type GeneralConfig struct {
	model.Model
	Key    string `gorm:"column:key" json:"key"`       //键
	Status int    `gorm:"column:status" json:"status"` //状态：0=》启用；1=》禁用
}

// TableName sets the insert table name for this struct type
func (g *GeneralConfig) TableName() string {
	return "general_configs"
}
