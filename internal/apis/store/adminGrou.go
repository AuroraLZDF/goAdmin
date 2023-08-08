// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package store

import (
	"strconv"
	"strings"

	"gorm.io/gorm"

	"apis/internal/pkg/model/admin"
	v1 "apis/internal/pkg/request/apis/v1"
)

type AdminGroupStore interface {
	Get(status int) (*[]admin.AdminGroup, error)
	Gets(r v1.PageRequest) (*[]admin.AdminGroup, int, error)
	InfoById(id int) (*admin.AdminGroup, error)
	UpdateOrCreate(r v1.AdminGroupUpdateRequest) error
	UpdateStatus(id int, status int) error
}

type adminGroups struct {
	db *gorm.DB
}

func NewAdminGroups(db *gorm.DB) *adminGroups {
	return &adminGroups{db: db}
}

func (s *adminGroups) Gets(r v1.PageRequest) (*[]admin.AdminGroup, int, error) {
	var groups *[]admin.AdminGroup
	var total int64

	if err := s.db.Model(&admin.AdminGroup{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := s.db.Offset((r.Page - 1) * r.PageSize).Limit(r.PageSize).Find(&groups).Error; err != nil {
		return nil, 0, err
	}

	return groups, int(total), nil
}

func (s *adminGroups) Get(status int) (*[]admin.AdminGroup, error) {
	var adminGroup *[]admin.AdminGroup
	err := s.db.Where("status = ?", status).Find(&adminGroup).Error
	return adminGroup, err
}

func (s *adminGroups) InfoById(id int) (*admin.AdminGroup, error) {
	var adminGroup *admin.AdminGroup
	err := s.db.Where("id = ?", id).Find(&adminGroup).Error
	return adminGroup, err
}

func (s *adminGroups) UpdateOrCreate(r v1.AdminGroupUpdateRequest) error {
	var group *admin.AdminGroup
	var err error

	if r.Id > 0 {
		// 更新
		group, err = s.InfoById(r.Id)
		if err != nil {
			return err
		}
	} else {
		// 创建
		group.Id = 0
	}

	strArr := make([]string, len(r.Rules))
	for i, v := range r.Rules {
		strArr[i] = strconv.Itoa(v)
	}
	result := strings.Join(strArr, ",")

	group.Title = r.Title
	group.Status = r.Status
	group.Rules = result

	return s.db.Save(&group).Error
}

func (s *adminGroups) UpdateStatus(id int, status int) error {
	var group *admin.AdminGroup
	err := s.db.Where("id = ?", id).Find(&group).Error
	if err != nil {
		return err
	}

	group.Status = status
	return s.db.Save(&group).Error
}
