// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package admin

import (
	"apis/internal/pkg/model"
)

type Places struct {
	model.Model
	Name string `gorm:"column:name" json:"name"` //
}

// TableName sets the insert table name for this struct type
func (p *Places) TableName() string {
	return "places"
}
