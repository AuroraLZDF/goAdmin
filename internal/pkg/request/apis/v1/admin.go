// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package v1

// LoginRequest 指定了 `POST /login` 接口的请求参数.
type LoginRequest struct {
	Phone    string `form:"phone" valid:"required,stringlength(11)"`
	Password string `form:"password" valid:"required,stringlength(6|18)"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

type ProfileRequest struct {
	Uid                  int    `form:"uid" valid:"required,numeric"`
	Name                 string `form:"name" valid:"required,maxstringlength(10)"`
	Phone                string `form:"phone" valid:"required,numeric"`
	Password             string `form:"password" valid:"minstringlength(6)"`
	PasswordConfirmation string `form:"password_confirmation" valid:"minstringlength(6)"`
}
