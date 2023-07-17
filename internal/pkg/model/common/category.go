// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package common

import (
	"apis/internal/pkg/model"
)

type Categories struct {
	model.Model
	Name        string `gorm:"column:name" json:"name"`                 //所属类目名称
	IsHidden    int    `gorm:"column:is_hidden" json:"is_hidden"`       //是否删除
	IconActive  string `gorm:"column:icon_active" json:"icon_active"`   //图标选中的样式
	Icon        string `gorm:"column:icon" json:"icon"`                 //图标
	Cover       string `gorm:"column:cover" json:"cover"`               //封面
	CoverMobile string `gorm:"column:cover_mobile" json:"cover_mobile"` //手机端分类图片地址
}

// TableName sets the insert table name for this struct type
func (c *Categories) TableName() string {
	return "categories"
}
