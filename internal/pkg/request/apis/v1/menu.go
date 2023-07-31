// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package v1

type MenuUpdateRequest struct {
	Id     int    `form:"id" json:"id"`
	Title  string `form:"title" json:"title" validate:"required"`
	Url    string `form:"url" json:"url" validate:"required"`
	Sort   int    `form:"sort" json:"sort" validate:"required"`
	Icon   string `form:"icon" json:"icon"`
	IsShow int    `form:"is_show" json:"is_show"`
	Level  int    `form:"level" json:"level"`
	Order  int    `form:"order" json:"order"`
	Pid    int    `form:"pid" json:"pid"`
	Status int    `form:"status" json:"status"`
	Tips   string `form:"tips" json:"tips"`
}
