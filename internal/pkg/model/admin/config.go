// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package admin

import (
	"apis/internal/pkg/model"
)

type Configs struct {
	model.Model
	K    string `gorm:"column:k;primary_key" json:"k"` //变量
	V    string `gorm:"column:v" json:"v"`             //值
	Type int    `gorm:"column:type" json:"type"`       //0系统，1自定义
	Name string `gorm:"column:name" json:"name"`       //说明
}

// TableName sets the insert table name for this struct type
func (c *Configs) TableName() string {
	return "configs"
}
