// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package admin

import (
	"gorm.io/gorm"

	"apis/internal/pkg/model"
	v1 "apis/internal/pkg/request/apis/v1"
)

type AdminLog struct {
	model.Model
	UID     int    `gorm:"column:uid" json:"uid"`         //操作者UID
	IP      string `gorm:"column:ip" json:"ip"`           //操作者IP
	Content string `gorm:"column:content" json:"content"` //日志内容
	URL     string `gorm:"column:url" json:"url"`         //操作的URL
	Request string `gorm:"column:request" json:"request"` //请求信息内容 post、get等数据
}

// TableName sets the insert table name for this struct type
func (a *AdminLog) TableName() string {
	return "admin_logs"
}

type AdminLogStore interface {
	Lists(r v1.PageRequest) (*[]AdminLog, int, error)
}

// AdminLogs 接口的实现
type AdminLogs struct {
	db *gorm.DB
}

// 确保 menus 实现了 MenuStore 接口
var _ AdminLogStore = (*AdminLogs)(nil)

// NewAdminLogs 创建一个 adminLogs 实现类
func NewAdminLogs(db *gorm.DB) *AdminLogs {
	return &AdminLogs{db}
}

func (log *AdminLogs) Lists(r v1.PageRequest) (*[]AdminLog, int, error) {
	var logs *[]AdminLog
	var total int64

	if err := log.db.Model(&AdminLog{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := log.db.Offset((r.Page - 1) * r.PageSize).Limit(r.PageSize).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	return logs, int(total), nil
}
