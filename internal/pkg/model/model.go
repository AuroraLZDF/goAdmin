// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package model

import (
	"apis/pkg/util"
)

type Model struct {
	ID        int64     `gorm:"column:id;primary_key" json:"id"`      //ID
	CreatedAt util.Time `gorm:"column:createdAt" json:"created_at"`   //创建时间
	UpdatedAt util.Time `gorm:"column:updatedAt" json:"updated_at"`   //更新时间
	DeletedAt util.Time `gorm:"column:deleted_at" json:"deleted_at" ` //删除时间
}
