// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package place

import (
	"context"

	"apis/internal/apis/store"
	"apis/internal/pkg/model/admin"
	v1 "apis/internal/pkg/request/apis/v1"
)

// PlaceBiz 定义要实现的接口
type PlaceBiz interface {
	Count(ctx context.Context) int
	Lists(ctx context.Context, r v1.PageRequest) (*[]admin.Places, error)
	Detail(ctx context.Context, id int) (*admin.Places, error)
	CreateOrUpdate(ctx context.Context, r v1.PlaceUpdateRequest) error
}

// PlaceBiz 接口的实现.
type placeBiz struct {
	ds store.IStore
}

// 确保 placeBiz 实现了 PlaceBiz 接口.
var _ PlaceBiz = (*placeBiz)(nil)

// New 创建一个实现了 PlaceBiz 接口的实例.
func New(ds store.IStore) *placeBiz {
	return &placeBiz{ds: ds}
}

func (b *placeBiz) Count(ctx context.Context) int {
	count := b.ds.Places().Count()
	return count
}

func (b *placeBiz) Lists(ctx context.Context, r v1.PageRequest) (*[]admin.Places, error) {
	lists, err := b.ds.Places().Gets(r)
	if err != nil {
		return nil, err
	}

	return lists, nil
}

func (b *placeBiz) Detail(ctx context.Context, id int) (*admin.Places, error) {
	info, err := b.ds.Places().Get(id)
	if err != nil {
		return nil, err
	}

	return info, nil
}

func (b *placeBiz) CreateOrUpdate(ctx context.Context, r v1.PlaceUpdateRequest) error {
	if err := b.ds.Places().CreateOrUpdate(r); err != nil {
		return err
	}

	return nil
}
