// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package admin

type AdminAccess struct {
	Uid     int `gorm:"column:uid;primary_key" json:"uid"`           //
	GroupID int `gorm:"column:group_id;primary_key" json:"group_id"` //
}

// TableName sets the insert table name for this struct type
func (a *AdminAccess) TableName() string {
	return "admin_accesses"
}
