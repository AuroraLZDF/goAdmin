// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package area

import (
	"context"
	"errors"

	"apis/internal/apis/store"
	"apis/internal/pkg/model/common"
	v1 "apis/internal/pkg/request/apis/v1"
)

// AreaBiz 定义要实现的接口
type AreaBiz interface {
	Lists(ctx context.Context, pid int) (*[]common.Areas, error)
	Detail(ctx context.Context, id int) (*common.Areas, error)
	CreateOrUpdate(ctx context.Context, r v1.AreaUpdateRequest) (int, error)
	Enable(ctx context.Context, id int) error
	Disable(ctx context.Context, id int) error
}

// AreaBiz 接口的实现.
type areaBiz struct {
	ds store.IStore
}

// 确保 areaBiz 实现了 AreaBiz 接口.
var _ AreaBiz = (*areaBiz)(nil)

// New 创建一个实现了 AreaBiz 接口的实例.
func New(ds store.IStore) *areaBiz {
	return &areaBiz{ds: ds}
}

// Lists 区域列表
func (b *areaBiz) Lists(ctx context.Context, pid int) (*[]common.Areas, error) {
	lists, err := b.ds.Areas().Gets(pid)
	if err != nil {
		return nil, err
	}

	return lists, nil
}

// Detail 区域详情
func (b *areaBiz) Detail(ctx context.Context, id int) (*common.Areas, error) {
	info, err := b.ds.Areas().Get(id)
	if err != nil {
		return nil, err
	}

	return info, nil
}

// CreateOrUpdate 更新区域信息
func (b *areaBiz) CreateOrUpdate(ctx context.Context, r v1.AreaUpdateRequest) (int, error) {
	if r.Pid > 0 {
		parent, err := b.ds.Areas().Get(r.Pid)
		if err != nil {
			return 0, err
		}

		r.Level = 1
		if parent.Level != 0 {
			r.Level = parent.Level + 1
		}
		if r.Level > 3 {
			return 0, errors.New("最多允许创建 3 级区域")
		}
	}

	has := b.ds.Areas().HasSameArea(r.Pid, r.Name, r.Id)
	if has {
		return 0, errors.New("存在相同的省市区项目")
	}

	id, err := b.ds.Areas().CreateOrUpdate(r)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Enable 启用区域信息
func (b *areaBiz) Enable(ctx context.Context, id int) error {
	info, err := b.ds.Areas().Get(id)
	if err != nil {
		return err
	}

	if err := b.ds.Areas().Enable(info); err != nil {
		return err
	}

	return nil
}

// Disable 禁用区域信息
func (b *areaBiz) Disable(ctx context.Context, id int) error {
	info, err := b.ds.Areas().Get(id)
	if err != nil {
		return err
	}

	if err := b.ds.Areas().Disable(info); err != nil {
		return err
	}

	return nil
}
