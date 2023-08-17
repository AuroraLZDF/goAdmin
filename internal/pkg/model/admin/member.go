// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package admin

import (
	"apis/internal/pkg/model"
	"apis/pkg/util"
)

type Member struct {
	model.Model
	Mid                    int       `gorm:"column:mid;primary_key" json:"mid"`                             //
	Phone                  string    `gorm:"column:phone" json:"phone"`                                     //手机号
	Openid                 string    `gorm:"column:openid" json:"openid"`                                   //微信登录openid
	Email                  string    `gorm:"column:email" json:"email"`                                     //邮箱
	Name                   string    `gorm:"column:name" json:"name"`                                       //真实姓名
	Nickname               string    `gorm:"column:nickname" json:"nickname"`                               //昵称
	Avatar                 string    `gorm:"column:avatar" json:"avatar"`                                   //用户头像
	Password               string    `gorm:"column:password" json:"password"`                               //密码
	RememberToken          string    `gorm:"column:remember_token" json:"remember_token"`                   //登录记住账户token
	Sex                    int       `gorm:"column:sex" json:"sex"`                                         //1为男性，2为女性
	ProvinceID             int       `gorm:"column:province_id" json:"province_id"`                         //省份ID
	CityID                 int       `gorm:"column:city_id" json:"city_id"`                                 //城市ID
	CountyID               int       `gorm:"column:county_id" json:"county_id"`                             //区县ID
	Qq                     string    `gorm:"column:qq" json:"qq"`                                           //qq号码
	Intro                  string    `gorm:"column:intro" json:"intro"`                                     //简介
	Authentication         int       `gorm:"column:authentication" json:"authentication"`                   //是否实名认证 0 没有实名 1 已实名认证
	BankAuthentication     int       `gorm:"column:bank_authentication" json:"bank_authentication"`         //是否已绑定银行卡 0 未绑定 1 已绑定
	PanoramaAuthentication int       `gorm:"column:panorama_authentication" json:"panorama_authentication"` //是否完成全景摄影师认证  0  未认证  1 已认证
	PlaneAuthentication    int       `gorm:"column:plane_authentication" json:"plane_authentication"`       //是否完成平面摄影师认证 0 未认证 1 已认证
	Oid                    int       `gorm:"column:oid" json:"oid"`                                         //机构ID
	Level                  int       `gorm:"column:level" json:"level"`                                     //用户组
	LimitNum               int       `gorm:"column:limit_num" json:"limit_num"`                             //可发布项目数量限制，默认为0，无限制
	LastLoginTime          util.Time `gorm:"column:last_login_time" json:"last_login_time"`                 //最近登录
	LastLoginIP            string    `gorm:"column:last_login_ip" json:"last_login_ip"`                     //最近登录IP
	Platform               string    `gorm:"column:platform" json:"platform"`                               //注册来源
	Status                 int       `gorm:"column:status" json:"status"`                                   //用户是否被禁用  0 正常 1禁用
	Amount                 float64   `gorm:"column:amount" json:"amount"`                                   //账户余额
	PanoramaProjectLimit   int       `gorm:"column:panorama_project_limit" json:"panorama_project_limit"`   //全景漫游项目数
	PanoramaPictureLimit   int       `gorm:"column:panorama_picture_limit" json:"panorama_picture_limit"`   //全景图片数
}

// TableName sets the insert table name for this struct type
func (m *Member) TableName() string {
	return "members"
}
