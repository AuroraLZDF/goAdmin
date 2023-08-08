// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package menu

import (
	"apis/internal/apis/store"
	"apis/internal/pkg/model/admin"
	v1 "apis/internal/pkg/request/apis/v1"
	"context"
	"errors"
)

// MenuBiz 定义要实现的接口
type MenuBiz interface {
	Count(ctx context.Context) int
	Lists(ctx context.Context, pid int) (*[]admin.Menus, error)
	Detail(ctx context.Context, id int) (*admin.Menus, error)
	CreateOrUpdate(ctx context.Context, r v1.MenuUpdateRequest) error
	Enable(ctx context.Context, id int) error
	Disable(ctx context.Context, id int) error
	RoleMenu(ctx context.Context, phone string) ([]map[string]interface{}, error)
	GetPid(menuId int) ([]int, error)
	Rules(ctx context.Context) ([]map[string]interface{}, error)
}

// MenuBiz 接口的实现
type menuBiz struct {
	ds store.IStore
}

// 确保 menuBiz 实现了 MenuBiz 接口
var _ MenuBiz = (*menuBiz)(nil)

// New 创建一个新的 menuBiz 实现类
func New(ds store.IStore) *menuBiz {
	return &menuBiz{ds: ds}
}

func (b *menuBiz) Count(ctx context.Context) int {
	count := b.ds.Menus().Count()
	return count
}

func (b *menuBiz) Lists(ctx context.Context, pid int) (*[]admin.Menus, error) {
	lists, err := b.ds.Menus().Gets(pid)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(*lists); i++ {
		menu := &(*lists)[i]
		menu.Leaf = true
		if menu.Type != "button" {
			menu.HasChildren = true
			menu.Leaf = false
		}
	}

	return lists, nil
}

func (b *menuBiz) Detail(ctx context.Context, id int) (*admin.Menus, error) {
	info, err := b.ds.Menus().Get(id)
	if err != nil {
		return nil, err
	}

	return info, nil
}

func (b *menuBiz) CreateOrUpdate(ctx context.Context, r v1.MenuUpdateRequest) error {
	var parent = &admin.Menus{}
	if r.Pid != 0 {
		if r.Id == 0 {
			// 判断是否存在相同名称的菜单
			has, _ := b.ds.Menus().MenuExists(r.Pid, r.Title)
			if has {
				return errors.New("已存在相同名称的菜单，请联系管理员")
			}
		}
		parent, _ = b.ds.Menus().Get(r.Pid)
	}

	r.Level = 1
	if parent.Id != 0 {
		r.Level = parent.Level + 1
	}
	if r.Level > 4 {
		return errors.New("最多允许创建 4 级菜单")
	}

	if err := b.ds.Menus().CreateOrUpdate(r); err != nil {
		return err
	}

	return nil
}

func (b *menuBiz) Enable(ctx context.Context, id int) error {
	info, err := b.ds.Menus().Get(id)
	if err != nil {
		return nil
	}

	if err := b.ds.Menus().Enable(info); err != nil {
		return err
	}

	return nil
}

func (b *menuBiz) Disable(ctx context.Context, id int) error {
	info, err := b.ds.Menus().Get(id)
	if err != nil {
		return nil
	}

	if err := b.ds.Menus().Disable(info); err != nil {
		return err
	}

	return nil
}

func (b *menuBiz) RoleMenu(ctx context.Context, phone string) ([]map[string]interface{}, error) {
	// 获取当前用户的所有菜单
	userInfo, err := b.ds.Admins().Get(phone)
	if err != nil {
		return nil, err
	}

	// 获取当前用户的所有菜单ID
	roleMenuIds, err := b.ds.Menus().GetAdminRoleMenus(userInfo.Id)
	if err != nil {
		return nil, err
	}

	// 获取当前用户左侧权限显示的菜单列表
	menuList, err := b.ds.Menus().GetRoleMenus(roleMenuIds)
	if err != nil {
		return nil, err
	}

	return menuList, nil
}

func (b *menuBiz) GetPid(menuId int) ([]int, error) {
	var pid = make([]int, 0)
	if menuId == 0 {
		return pid, nil
	}

	menu, err := b.ds.Menus().Get(menuId)
	if err != nil {
		return nil, err
	}

	if menu.Id != 0 && menu.Pid != 0 {
		pid = append(pid, menu.Pid)
		b.GetPid(menu.Pid)
	}

	return pid, nil
}

func (b *menuBiz) Rules(ctx context.Context) ([]map[string]interface{}, error) {
	return b.ds.Menus().GetRules()
}
