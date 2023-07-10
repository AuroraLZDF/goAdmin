// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

/**
 * 问题：
 * 使用 json 格式化 struct 时，time.Time 被格式化成
 * ”2006-01-02T15:04:05.999999999Z07:00“ 格式
 *
 * golang 的 time.Time 的默认 json 格式化格式叫做 RFC3339。
 * 好像是一种国际标准，被推荐用作 json 时间的标准格式。
 * 但是使用中不需要这种，而且不容易解析。
 * 示例：
	type Model struct {
		ID        int64      `gorm:"column:id;primary_key" json:"id"`    //
		CreatedAt util.Time `gorm:"column:createdAt" json:"created_at"` //
		UpdatedAt util.Time `gorm:"column:updatedAt" json:"updated_at"` //
	}
*/

package util

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const (
	TimeFormat = time.DateTime
)

type Time time.Time

// UnmarshalJSON 格式化自定义时间转换为 time.Time
func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

// MarshalJSON 自定义时间格式化输出
func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

// String 格式化输出
func (t Time) String() string {
	return time.Time(t).Format(TimeFormat)
}

// Scan 实现 GORM 的 Scanner 接口
// 将数据库返回的值转换为 util.Time 类型
func (t *Time) Scan(value interface{}) error {
	if value == nil {
		*t = Time(time.Time{})
		return nil
	}

	if v, ok := value.(time.Time); ok {
		*t = Time(v)
		return nil
	}

	return fmt.Errorf("failed to scan Time value")
}

// Value 实现 GORM 的 Valuer 接口
// 将 util.Time 类型转换为数据库可接受的值
func (t Time) Value() (driver.Value, error) {
	return time.Time(t), nil
}
