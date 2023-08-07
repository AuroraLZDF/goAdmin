// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package v1

// LoginRequest 指定了 `POST /admin/auth/login` 接口的请求参数.
type LoginRequest struct {
	Phone    string `form:"phone" validate:"required,len=11"`
	Password string `form:"password" validate:"required,max=18,min=6"`
}

// TokenResponse 指定了 `POST /admin/auth/login` 接口的返回值
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

// ProfileRequest 指定了 `POST /admin/profile/save` 接口的请求参数
type ProfileRequest struct {
	Uid                  int    `form:"uid" validate:"required,numeric"`
	Name                 string `form:"name" validate:"required,max=10"`
	Phone                string `form:"phone" validate:"required,numeric"`
	Password             string `form:"password" validate:"min=6"`
	PasswordConfirmation string `form:"password_confirmation" validate:"min=6"`
	Avatar               string `form:"avatar"`
}

type AdminUpdateRequest struct {
	Id       int    `form:"id" validate:"numeric"`
	Name     string `form:"name" validate:"required,max=10"`
	Phone    string `form:"phone" validate:"required,numeric"`
	Password string `form:"password" validate:"min=6"`
	Avatar   string `form:"avatar"`
	Status   int    `form:"status" validate:"numeric"`
}

type AdminAccessUpdateRequest struct {
	Uid   int   `form:"uid" validate:"required,numeric"`
	Rules []int `form:"rules" validate:"required"`
}
