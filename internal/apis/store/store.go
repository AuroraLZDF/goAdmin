// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package store

import (
	"sync"

	"gorm.io/gorm"
)

var (
	once sync.Once
	// S 全局变量，方便其它包直接调用已初始化好的 S 实例.
	S *DataStore
)

// IStore 定义了 Store 层需要实现的方法.
type IStore interface {
	Admins() AdminStore
	Areas() AreaStore
	Categories() CategoryStore
	Places() PlaceStore
}

// DataStore 是 IStore 的一个具体实现.
type DataStore struct {
	db *gorm.DB
}

// 确保 datastore 实现了 IStore 接口.
var _ IStore = (*DataStore)(nil)

// NewStore 创建一个 IStore 类型的实例.
func NewStore(db *gorm.DB) *DataStore {
	// 确保 S 只被初始化一次
	once.Do(func() {
		S = &DataStore{db}
	})

	return S
}

// Admins 返回一个实现了 AdminStore 接口的实例
func (ds *DataStore) Admins() AdminStore {
	return newAdmins(ds.db)
}

// Areas 返回一个实现了 AreaStore 接口的实例
func (ds *DataStore) Areas() AreaStore {
	return newAreas(ds.db)
}

// Categories 返回一个实现了 CategoryStore 接口的实例
func (ds *DataStore) Categories() CategoryStore {
	return newCategories(ds.db)
}

// Places 返回一个实现了 PlaceStore 接口的实例
func (ds *DataStore) Places() PlaceStore {
	return newPlace(ds.db)
}
