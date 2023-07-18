// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package general

import (
	"context"

	"apis/internal/apis/store"
	"apis/internal/pkg/model/admin"
	v1 "apis/internal/pkg/request/apis/v1"
)

// GeneralBiz 定义要实现的接口
type GeneralBiz interface {
	Lists(ctx context.Context, r v1.PageRequest) (*[]admin.GeneralConfigs, error)
	Detail(ctx context.Context, id int) (*admin.GeneralConfigs, error)
	CreateOrUpdate(ctx context.Context, r v1.GeneralUpdateRequest) error
	Enable(ctx context.Context, id int) error
	Disable(ctx context.Context, id int) error
}

// GeneralBiz 接口的实现.
type generalBiz struct {
	ds store.IStore
}

// 确保 generalBiz 实现了 GeneralBiz 接口.
var _ GeneralBiz = (*generalBiz)(nil)

// New 创建一个实现了 GeneralBiz 接口的实例.
func New(ds store.IStore) *generalBiz {
	return &generalBiz{ds: ds}
}

func (b *generalBiz) Lists(ctx context.Context, r v1.PageRequest) (*[]admin.GeneralConfigs, error) {
	lists, err := b.ds.Generals().Gets(r)
	if err != nil {
		return nil, err
	}

	return lists, nil
}

func (b *generalBiz) Detail(ctx context.Context, id int) (*admin.GeneralConfigs, error) {
	info, err := b.ds.Generals().Get(id)
	if err != nil {
		return nil, err
	}

	return info, nil
}

func (b *generalBiz) CreateOrUpdate(ctx context.Context, r v1.GeneralUpdateRequest) error {
	if err := b.ds.Generals().CreateOrUpdate(r); err != nil {
		return err
	}

	return nil
}

func (b *generalBiz) Enable(ctx context.Context, id int) error {
	if err := b.ds.Generals().Enable(id); err != nil {
		return err
	}

	return nil
}

func (b *generalBiz) Disable(ctx context.Context, id int) error {
	if err := b.ds.Generals().Disable(id); err != nil {
		return err
	}

	return nil
}
