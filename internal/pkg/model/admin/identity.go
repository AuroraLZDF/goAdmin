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

type Identity struct {
	model.Model
	Mid         int    `gorm:"column:mid" json:"mid"`                   //会员id
	RealName    string `gorm:"column:real_name" json:"real_name"`       //真实姓名
	CardNumber  string `gorm:"column:card_number" json:"card_number"`   //身份证号码
	Phone       string `gorm:"column:phone" json:"phone"`               //手机号码
	CardFront   string `gorm:"column:card_front" json:"card_front"`     //身份证照片正面
	CardBackend string `gorm:"column:card_backend" json:"card_backend"` //身份证照片反面
	Status      int    `gorm:"column:status" json:"status"`             //身份验证 0:未验证 1:审核通过 2:审核驳回
	Reason      string `gorm:"column:reason" json:"reason"`             //驳回原因
}

// TableName sets the insert table name for this struct type
func (i *Identity) TableName() string {
	return "identities"
}

type IdentityStore interface {
	Lists(r v1.PageRequest) (*[]Identity, int, error)
	Detail(r v1.IdRequest) (*Identity, error)
	Pass(r v1.IdRequest) error
	Reject(r v1.IdentityRequest) error
}

// Identities 接口的实现
type Identities struct {
	db *gorm.DB
}

// 确保 Identities 实现了 IdentityStore 接口
var _ IdentityStore = (*Identities)(nil)

// NewIdentities 创建一个 Identities 实现类
func NewIdentities(db *gorm.DB) *Identities {
	return &Identities{db}
}

func (id *Identities) Lists(r v1.PageRequest) (*[]Identity, int, error) {
	var identity *[]Identity
	var total int64

	if err := id.db.Model(&Identity{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := id.db.Offset((r.Page - 1) * r.PageSize).Limit(r.PageSize).Find(&identity).Error; err != nil {
		return nil, 0, err
	}

	// todo: 数字脱敏
	/*if len(*identity) > 0 {
		for _, list := range *identity {
			if list.CardNumber != "" {
				list.CardNumber = ""
			}
			if list.Phone != "" {
				list.Phone = ""
			}
		}
	}*/

	return identity, int(total), nil
}

func (id *Identities) Detail(r v1.IdRequest) (*Identity, error) {
	var identity *Identity
	err := id.db.Model(Identity{}).Where("id=?", r.Id).First(&identity).Error
	if err != nil {
		return nil, err
	}

	return identity, nil
}

func (id *Identities) Pass(r v1.IdRequest) error {
	var identity *Identity
	if err := id.db.Model(Identity{}).Where("id=?", r.Id).First(&identity).Error; err != nil {
		return err
	}

	identity.Status = model.StatusOff
	if err := id.db.Model(Identity{}).Save(identity).Error; err != nil {
		return err
	}

	var member *Member
	if err := id.db.Model(Member{}).Where("mid=?", identity.Mid).First(&member).Error; err != nil {
		return err
	}

	member.Authentication = model.StatusOff
	if err := id.db.Model(Member{}).Save(member).Error; err != nil {
		return err
	}

	return nil
}

func (id *Identities) Reject(r v1.IdentityRequest) error {
	var identity *Identity
	if err := id.db.Model(Identity{}).Where("id=?", r.Id).First(&identity).Error; err != nil {
		return err
	}

	identity.Status = model.StatusSecond
	identity.Reason = r.Reason
	if err := id.db.Model(Identity{}).Save(identity).Error; err != nil {
		return err
	}

	return nil
}
