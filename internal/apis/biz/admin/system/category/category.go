// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package category

import (
	v1 "apis/internal/pkg/request/apis/v1"
	"context"

	"apis/internal/apis/store"
	"apis/internal/pkg/model/common"
)

// CategoryBiz 定义要实现的接口
type CategoryBiz interface {
	Count(ctx context.Context) int
	Lists(ctx context.Context, r v1.PageRequest) (*[]common.Categories, error)
	Detail(ctx context.Context, id int) (*common.Categories, error)
	CreateOrUpdate(ctx context.Context, r v1.CategoryUpdateRequest) error
	Enable(ctx context.Context, id int) error
	Disable(ctx context.Context, id int) error
}

// CategoryBiz 接口的实现.
type categoryBiz struct {
	ds store.IStore
}

// 确保 categoryBiz 实现了 CategoryBiz 接口.
var _ CategoryBiz = (*categoryBiz)(nil)

// New 创建一个实现了 CategoryBiz 接口的实例.
func New(ds store.IStore) *categoryBiz {
	return &categoryBiz{ds: ds}
}

func (b *categoryBiz) Count(ctx context.Context) int {
	count := b.ds.Categories().Count()
	return count
}

func (b *categoryBiz) Lists(ctx context.Context, r v1.PageRequest) (*[]common.Categories, error) {

	lists, err := b.ds.Categories().Gets(r)
	if err != nil {
		return nil, err
	}

	return lists, nil
}

func (b *categoryBiz) Detail(ctx context.Context, id int) (*common.Categories, error) {
	info, err := b.ds.Categories().Get(id)
	if err != nil {
		return nil, err
	}

	return info, nil
}

func (b *categoryBiz) CreateOrUpdate(ctx context.Context, r v1.CategoryUpdateRequest) error {
	if err := b.ds.Categories().CreateOrUpdate(r); err != nil {
		return err
	}

	return nil
}

func (b *categoryBiz) Enable(ctx context.Context, id int) error {
	info, err := b.ds.Categories().Get(id)
	if err != nil {
		return nil
	}

	if err := b.ds.Categories().Enable(info); err != nil {
		return err
	}

	return nil
}

func (b *categoryBiz) Disable(ctx context.Context, id int) error {
	info, err := b.ds.Categories().Get(id)
	if err != nil {
		return nil
	}

	if err := b.ds.Categories().Disable(info); err != nil {
		return err
	}

	return nil
}
