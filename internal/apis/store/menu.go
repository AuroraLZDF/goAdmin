// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package store

import (
	"apis/internal/pkg/model"
	"apis/internal/pkg/model/admin"
	v1 "apis/internal/pkg/request/apis/v1"
	"apis/pkg/util"
	"gorm.io/gorm"
	"reflect"
	"strconv"
	"strings"
)

// MenuStore 定义了 menu 模块在 store 中的接口
type MenuStore interface {
	Count() int
	Gets(pid int) (*[]admin.Menus, error)
	Get(id int) (*admin.Menus, error)
	CreateOrUpdate(r v1.MenuUpdateRequest) error
	Enable(*admin.Menus) error
	Disable(*admin.Menus) error
	RoleMenu(roleId int) (*[]admin.Menus, error)
	MenuExists(pid int, title string) (bool, error)
	GetAdminRoleMenus(uid int) ([]int, error)
	GetRoleMenus(menuIds []int) ([]map[string]interface{}, error)
}

// MenuStore 接口的实现
type menus struct {
	db *gorm.DB
}

// 确保 menus 实现了 MenuStore 接口
var _ MenuStore = (*menus)(nil)

// NewMenus 创建一个 menus 实现类
func NewMenus(db *gorm.DB) *menus {
	return &menus{db}
}

func (m *menus) Count() int {
	var count int64
	if err := m.db.Model(&admin.Menus{}).Count(&count).Error; err != nil {
		return 0
	}
	return int(count)
}

func (m *menus) Gets(pid int) (*[]admin.Menus, error) {
	var lists []admin.Menus

	err := m.db.Where("pid=?", pid).Order("sort asc").Find(&lists).Error
	if err != nil {
		return nil, err
	}

	return &lists, nil
}

func (m *menus) Get(id int) (*admin.Menus, error) {
	var info admin.Menus

	err := m.db.Where("id=?", id).First(&info).Error
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func (m *menus) CreateOrUpdate(r v1.MenuUpdateRequest) error {
	if r.Id != 0 {
		if err := m.db.Model(admin.Menus{}).Where("id=?", r.Id).Updates(&r).Error; err != nil {
			return err
		}
	} else {
		if err := m.db.Model(admin.Menus{}).Create(&r).Error; err != nil {
			return err
		}
	}
	return nil
}

func (m *menus) Enable(menu *admin.Menus) error {
	menu.Status = model.StatusOn
	if err := m.db.Save(&menu).Error; err != nil {
		return err
	}
	return nil
}

func (m *menus) Disable(menu *admin.Menus) error {
	menu.Status = model.StatusOff
	if err := m.db.Save(&menu).Error; err != nil {
		return err
	}
	return nil
}

func (m *menus) RoleMenu(roleId int) (*[]admin.Menus, error) {
	return nil, nil
}

func (m *menus) MenuExists(pid int, title string) (bool, error) {
	var count int64
	err := m.db.Model(&admin.Menus{}).Where("pid = ? AND title = ?", pid, title).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (m *menus) GetAdminRoleMenus(uid int) ([]int, error) {
	var access = NewAdminAccesses(m.db)

	data, err := access.GetAdminRoleMenus(uid)
	if err != nil {
		return nil, err
	}

	var groups []admin.AdminGroup
	err = m.db.Model(admin.AdminGroup{}).Where("id in (?)", data).Find(&groups).Error
	if err != nil {
		return nil, err
	}

	var menuIds []string
	if len(groups) > 0 {
		for _, group := range groups {
			menuIds = append(menuIds, group.Rules)
		}
	}

	menuIDString := strings.Join(menuIds, ",")
	// 使用 strings.Split 函数将字符串拆分为切片
	menuIDs := strings.Split(menuIDString, ",")

	// 使用 map 来进行去重
	uniqueMenuIDs := make(map[string]bool)
	for _, id := range menuIDs {
		uniqueMenuIDs[id] = true
	}

	// 将 map 中的键转换为数组
	var result []int
	for id := range uniqueMenuIDs {
		_id, _ := strconv.Atoi(id)
		result = append(result, _id)
	}

	return result, nil
}

func (m *menus) GetRoleMenus(menuIds []int) ([]map[string]interface{}, error) {
	var menus []admin.Menus
	err := m.db.Model(admin.Menus{}).Where("status = ?", model.StatusOn).
		Where("id in (?)", menuIds).Order("id asc").Find(&menus).Error
	if err != nil {
		return nil, err
	}

	if len(menus) <= 0 {
		return nil, nil // 没有权限菜单，退出
	}

	// 将 menus 转换为 []map[string]interface{}
	var result []map[string]interface{}
	for _, menu := range menus {
		menuMap := make(map[string]interface{})
		menuValue := reflect.ValueOf(menu)
		menuType := menuValue.Type()

		for i := 0; i < menuValue.NumField(); i++ {
			field := menuValue.Field(i)
			fieldName := menuType.Field(i).Name
			menuMap[fieldName] = field.Interface()
		}

		result = append(result, menuMap)
	}

	var config = util.TreeService{
		PrimaryKey: "Id",
		ParentKey:  "Pid",
	}
	t := util.NewTreeService(config)
	tree := t.MakeTree(result)

	return tree, nil
}
