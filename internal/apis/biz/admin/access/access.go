// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package access

import (
	"apis/internal/apis/store"
	v1 "apis/internal/pkg/request/apis/v1"
	"context"
)

type AdminAccessBiz interface {
	UserRoles(ctx context.Context, uid int) ([]int, error)
	Update(ctx context.Context, r v1.AdminAccessUpdateRequest) error
}

type adminAccessBiz struct {
	ds store.IStore
}

var _ AdminAccessBiz = (*adminAccessBiz)(nil)

func New(ds store.IStore) *adminAccessBiz {
	return &adminAccessBiz{ds: ds}
}

func (b *adminAccessBiz) UserRoles(ctx context.Context, uid int) ([]int, error) {
	return b.ds.AdminAccesses().GetUserRoles(uid)
}

func (b *adminAccessBiz) Update(ctx context.Context, r v1.AdminAccessUpdateRequest) error {
	return b.ds.AdminAccesses().Update(r)
}
