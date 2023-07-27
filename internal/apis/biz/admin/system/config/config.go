// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package config

import (
	"apis/internal/apis/store"
	v1 "apis/internal/pkg/request/apis/v1"
	"context"
)

// ConfigBiz 定义要实现的接口
type ConfigBiz interface {
	Detail(ctx context.Context) (map[string]string, error)
	Update(ctx context.Context, r v1.ConfigUpdateRequest) error
}

// ConfigBiz 接口的实现
type configBiz struct {
	ds store.IStore
}

// 确保 configBiz 实现了 ConfigBiz 接口
var _ ConfigBiz = (*configBiz)(nil)

// New 创建一个 ConfigBiz 接口的实例
func New(ds store.IStore) *configBiz {
	return &configBiz{ds}
}

// Detail 获取配置详情
func (c *configBiz) Detail(ctx context.Context) (map[string]string, error) {
	configs, err := c.ds.Configs().Get()
	if err != nil {
		return nil, err
	}

	data := make(map[string]string)
	for _, config := range *configs {
		// 获取每个 config 的 K 字段
		data[config.K] = config.V
	}

	return data, nil
}

// Update 保存配置
func (c *configBiz) Update(ctx context.Context, r v1.ConfigUpdateRequest) error {
	return c.ds.Configs().Update(r)
}
