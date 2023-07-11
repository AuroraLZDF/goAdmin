// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package admin

import (
	"context"
	"errors"
	"net/http"

	"apis/internal/apis/store"
	"apis/internal/pkg/errno"
	"apis/internal/pkg/model"
	"apis/internal/pkg/request/apis/v1"
	"apis/pkg/auth"
	"apis/pkg/token"
)

type AdminBiz interface {
	Login(ctx context.Context, r *v1.LoginRequest) (v1.TokenResponse, error)
	Logout(ctx context.Context, request *http.Request) error
	Info(ctx context.Context, phone string) (*model.Admins, error)
	RefreshToken(ctx context.Context, request *http.Request, phone string) (v1.TokenResponse, error)
}

// AdminBiz 接口的实现.
type adminBiz struct {
	ds store.IStore
}

// 确保 userBiz 实现了 UserBiz 接口.
var _ AdminBiz = (*adminBiz)(nil)

// New 创建一个实现了 UserBiz 接口的实例.
func New(ds store.IStore) *adminBiz {
	return &adminBiz{ds: ds}
}

// Login 是 UserBiz 接口中 `Login` 方法的实现.
func (b *adminBiz) Login(ctx context.Context, r *v1.LoginRequest) (v1.TokenResponse, error) {
	var response = v1.TokenResponse{}

	// 获取登录用户的所有信息
	user, err := b.ds.Admins().Get(r.Phone)
	if err != nil {
		return response, errors.New("user not found")
	}

	// 对比传入的明文密码和数据库中已加密过的密码是否匹配
	if err := auth.Compare(user.Password, r.Password); err != nil {
		return response, errors.New("password is not correct")
	}

	// 如果匹配成功，说明登录成功，签发 token 并返回
	response, err = token.Sign(r.Phone)
	if err != nil {
		return response, errno.ErrSignToken
	}

	return response, nil
}

func (b *adminBiz) Info(ctx context.Context, phone string) (*model.Admins, error) {
	user, err := b.ds.Admins().Get(phone)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (b *adminBiz) Logout(ctx context.Context, r *http.Request) error {
	if err := token.AddToBlacklist(r); err != nil {
		return err
	}

	return nil
}

func (b *adminBiz) RefreshToken(ctx context.Context, r *http.Request, phone string) (v1.TokenResponse, error) {
	// 如果匹配成功，说明登录成功，签发 token 并返回
	var response, err = token.Sign(phone)
	if err != nil {
		return response, errno.ErrSignToken
	}

	// 过期之前的 token
	if err = token.AddToBlacklist(r); err != nil {
		return response, err
	}

	return response, nil
}
