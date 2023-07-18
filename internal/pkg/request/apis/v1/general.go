// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package v1

type GeneralUpdateRequest struct {
	Id     int    `form:"id" validate:"numeric"`
	Key    string `form:"key" validate:"required"`
	Status int    `form:"status" validate:"numeric"`
}
