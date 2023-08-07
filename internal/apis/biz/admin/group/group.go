// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package group

import (
	"context"

	"apis/internal/apis/store"
	"apis/internal/pkg/model/admin"
)

type AdminGroupBiz interface {
	Get(ctx context.Context, status int) (*[]admin.AdminGroup, error)
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
