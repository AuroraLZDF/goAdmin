// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package v1

type PlaceUpdateRequest struct {
	Id   int    `form:"id" validate:"numeric"`
	Name string `form:"name" validate:"required"`
}
