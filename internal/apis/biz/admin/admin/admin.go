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
	"apis/internal/pkg/model/admin"
	"apis/internal/pkg/request/apis/v1"
	"apis/pkg/auth"
	"apis/pkg/token"
)

// AdminBiz 定义要实现的接口
type AdminBiz interface {
	Login(ctx context.Context, r *v1.LoginRequest) (v1.TokenResponse, error)
	Logout(ctx context.Context, request *http.Request) error
	Info(ctx context.Context, phone string) (*admin.Admins, error)
	RefreshToken(ctx context.Context, request *http.Request, phone string) (v1.TokenResponse, error)
	Lists(ctx context.Context, r v1.PageRequest) (*[]admin.Admins, int, error)
	Update(ctx context.Context, r v1.AdminUpdateRequest) error
	Detail(ctx context.Context, id int) (*admin.Admins, error)
	Enable(ctx context.Context, id int) error
	Disable(ctx context.Context, id int) error
}

// AdminBiz 接口的实现.
type adminBiz struct {
	ds store.IStore
}

// 确保 adminBiz 实现了 AdminBiz 接口.
var _ AdminBiz = (*adminBiz)(nil)

// New 创建一个实现了 AdminBiz 接口的实例.
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

func (b *adminBiz) Info(ctx context.Context, phone string) (*admin.Admins, error) {
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

func (b *adminBiz) Lists(ctx context.Context, r v1.PageRequest) (*[]admin.Admins, int, error) {
	lists, total, err := b.ds.Admins().Gets(r)
	if err != nil {
		return nil, 0, err
	}
	return lists, total, nil
}

func (b *adminBiz) Detail(ctx context.Context, id int) (*admin.Admins, error) {
	detail, err := b.ds.Admins().InfoById(id)
	if err != nil {
		return nil, err
	}
	return detail, nil
}

func (b *adminBiz) Update(ctx context.Context, r v1.AdminUpdateRequest) error {
	if r.Id != 0 {
		user, err := b.ds.Admins().InfoById(r.Id)
		if err != nil {
			return err
		}

		if user.Password == r.Password {
			r.Password = ""
		}

		if r.Password != "" {
			r.Password, _ = auth.Encrypt(r.Password)
		}

		user.Name = r.Name
		user.Phone = r.Phone
		user.Avatar = r.Avatar
		user.Status = r.Status
		user.Password = r.Password
		if err = b.ds.Admins().Update(user); err != nil {
			return err
		}
	} else {
		if r.Password != "" {
			return errors.New("密码不能为空")
		}

		r.Password, _ = auth.Encrypt(r.Password)
		if err := b.ds.Admins().Create(r); err != nil {
			return err
		}
	}

	return nil
}

func (b *adminBiz) Enable(ctx context.Context, id int) error {
	err := b.ds.Admins().Enable(id)
	if err != nil {
		return err
	}

	return nil
}

func (b *adminBiz) Disable(ctx context.Context, id int) error {
	err := b.ds.Admins().Disable(id)
	if err != nil {
		return err
	}

	return nil
}
