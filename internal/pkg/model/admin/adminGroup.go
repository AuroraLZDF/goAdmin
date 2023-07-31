// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package admin

import "apis/internal/pkg/model"

type AdminGroup struct {
	model.Model
	Title  string `gorm:"column:title" json:"title"`   //名称
	Status int    `gorm:"column:status" json:"status"` //状态 0 启用 1 禁用
	Rules  string `gorm:"column:rules" json:"rules"`   //授权相关的规则组合
}

// TableName sets the insert table name for this struct type
func (a *AdminGroup) TableName() string {
	return "admin_groups"
}
