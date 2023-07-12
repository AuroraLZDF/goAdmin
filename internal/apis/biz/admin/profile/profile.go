// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package profile

import (
	"errors"

	"github.com/gin-gonic/gin"

	"apis/internal/apis/store"
	"apis/internal/pkg/known"
	v1 "apis/internal/pkg/request/apis/v1"
	"apis/pkg/auth"
)

// ProfileBiz 定义要实现的接口
type ProfileBiz interface {
	Save(ctx *gin.Context, r *v1.ProfileRequest) error
}

// ProfileBiz 接口的实现.
type profileBiz struct {
	ds store.IStore
}

// 确保 profileBiz 实现了 ProfileBiz 接口.
var _ ProfileBiz = (*profileBiz)(nil)

// New 创建一个实现了 ProfileBiz 接口的实例.
func New(ds store.IStore) *profileBiz {
	return &profileBiz{ds: ds}
}

// Save 更新账户信息
func (b *profileBiz) Save(ctx *gin.Context, r *v1.ProfileRequest) error {
	if ctx.GetString(known.XUsernameKey) != r.Phone {
		return errors.New("invalid request")
	}

	admin, err := b.ds.Admins().Get(r.Phone)
	if err != nil {
		return err
	}

	if len(r.Password) != 0 {
		if r.Password != r.PasswordConfirmation {
			return errors.New("confirmation password not same as password")
		}
		admin.Password, _ = auth.Encrypt(r.Password)
	}

	if len(r.Name) != 0 {
		admin.Name = r.Name
	}

	if err := b.ds.Admins().Update(admin); err != nil {
		return err
	}

	return nil
}
