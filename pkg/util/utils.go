// Copyright 2022 Innkeeper auroralzdf auroralzdf@gmail.com. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/auroralzdf/apis.

package util

import (
	"fmt"
	"reflect"
	"unicode"
)

// StructToMap converts a struct to a map
func StructToMap(obj interface{}) map[string]string {
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() != reflect.Struct {
		return nil
	}

	m := make(map[string]string)
	objType := objValue.Type()

	for i := 0; i < objValue.NumField(); i++ {
		fieldName := objType.Field(i).Name
		// 转换为下划线形式的字段名称
		snakeFieldName := toSnakeCase(fieldName)
		fieldValue := fmt.Sprintf("%v", objValue.Field(i).Interface())
		m[snakeFieldName] = fieldValue
	}

	return m
}

// toSnakeCase Converts a string to snake case.
func toSnakeCase(name string) string {
	var result []rune

	for i, r := range name {
		if unicode.IsUpper(r) {
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}

	return string(result)
}
