// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package group

import (
	"context"

	"apis/internal/apis/store"
	"apis/internal/pkg/model/admin"
	v1 "apis/internal/pkg/request/apis/v1"
)

type AdminGroupBiz interface {
	Get(ctx context.Context, status int) (*[]admin.AdminGroup, error)
	Lists(ctx context.Context, r v1.PageRequest) (*[]admin.AdminGroup, int, error)
	Detail(ctx context.Context, id int) (*admin.AdminGroup, error)
	Update(ctx context.Context, r v1.AdminGroupUpdateRequest) error
	Enable(ctx context.Context, id int) error
	Disable(ctx context.Context, id int) error
}

type adminGroupBiz struct {
	ds store.IStore
}

var _ AdminGroupBiz = (*adminGroupBiz)(nil)

func New(ds store.IStore) *adminGroupBiz {
	return &adminGroupBiz{ds: ds}
}

// Get admin group
func (b *adminGroupBiz) Get(ctx context.Context, status int) (*[]admin.AdminGroup, error) {
	return b.ds.AdminGroups().Get(status)
}

func (b *adminGroupBiz) Lists(ctx context.Context, r v1.PageRequest) (*[]admin.AdminGroup, int, error) {
	return b.ds.AdminGroups().Gets(r)
}
func (b *adminGroupBiz) Detail(ctx context.Context, id int) (*admin.AdminGroup, error) {
	return b.ds.AdminGroups().InfoById(id)
}
func (b *adminGroupBiz) Update(ctx context.Context, r v1.AdminGroupUpdateRequest) error {
	return b.ds.AdminGroups().UpdateOrCreate(r)
}
func (b *adminGroupBiz) Enable(ctx context.Context, id int) error {
	return b.ds.AdminGroups().UpdateStatus(id, 0)
}
func (b *adminGroupBiz) Disable(ctx context.Context, id int) error {
	return b.ds.AdminGroups().UpdateStatus(id, 1)
}
