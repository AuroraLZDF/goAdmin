// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package model

type Admins struct {
	Model
	Name          string `gorm:"column:name" json:"name"`        //管理员姓名
	Avatar        string `gorm:"column:avatar" json:"avatar"`    //管理员头像地址
	Phone         string `gorm:"column:phone" json:"phone"`      //管理员手机号码、同时也是登录账号
	Password      string `gorm:"column:password" json:"-"`       //登录密码
	RememberToken string `gorm:"column:remember_token" json:"-"` //登录记住账户token
	Status        int    `gorm:"column:status" json:"status"`    //账号状态 0 正常 1 被禁用
}

// TableName sets the insert table name for this struct type
func (a *Admins) TableName() string {
	return "admins"
}
