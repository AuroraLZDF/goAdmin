// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package admin

import "apis/pkg/util"

type Menus struct {
	//model.Model	// 结构体嵌套会导致反射无法导出嵌套字段
	Id          int       `gorm:"column:id;primary_key" json:"id"`     //ID
	CreatedAt   util.Time `gorm:"column:created_at" json:"created_at"` //创建时间
	UpdatedAt   util.Time `gorm:"column:updated_at" json:"updated_at"` //更新时间
	Pid         int       `gorm:"column:pid" json:"pid"`               //父级菜单ID
	URL         string    `gorm:"column:url" json:"url"`               //菜单URL
	Title       string    `gorm:"column:title" json:"title"`           //菜单名称
	Component   string    `gorm:"column:component" json:"component"`   //组件路径
	Icon        string    `gorm:"column:icon" json:"icon"`             //菜单ICON图标
	IsShow      int       `gorm:"column:is_show" json:"is_show"`       //是否在操作菜单列表显示 0 显示 1 不显示
	Level       int       `gorm:"column:level" json:"level"`           //菜单层级
	Sort        int       `gorm:"column:sort" json:"sort"`             //排序 越低的越排前面
	Order       string    `gorm:"column:order" json:"order"`           //菜单节点排序
	Type        string    `gorm:"column:type" json:"type"`             //
	Tips        string    `gorm:"column:tips" json:"tips"`             //提示
	Status      int       `gorm:"column:status" json:"status"`         //菜单是否可用 0 可用 1不可用
	Leaf        bool      `gorm:"-" json:"leaf"`
	HasChildren bool      `gorm:"-" json:"hasChildren"`
}

// TableName sets the insert table name for this struct type
func (m *Menus) TableName() string {
	return "menus"
}
