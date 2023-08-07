// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package store

import (
	"apis/internal/pkg/model/admin"
	"gorm.io/gorm"
)

type AdminGroupStore interface {
	Get(status int) (*[]admin.AdminGroup, error)
}

type adminGroups struct {
	db *gorm.DB
}

func NewAdminGroups(db *gorm.DB) *adminGroups {
	return &adminGroups{db: db}
}

func (s *adminGroups) Get(status int) (*[]admin.AdminGroup, error) {
	var adminGroup *[]admin.AdminGroup
	err := s.db.Where("status = ?", status).Find(&adminGroup).Error
	return adminGroup, err
}
