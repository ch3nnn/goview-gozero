/**
 * @Author: chentong
 * @Date: 2023/11/03 14:45
 */

package gorms

import (
	"strings"

	"gorm.io/gen/field"
)

func FieldIsDesc(field string) (bool, string) {
	if strings.HasPrefix(field, "-") {
		return true, field[1:]
	}
	return false, field
}

func LikeInner(s string) string {
	return "%" + s + "%"
}

func LikeLeft(s string) string {
	return "%" + s
}

func LikeRight(s string) string {
	return s + "%"
}

// 获取 model 列名
func ColumnName(field field.Field) string {
	return field.ColumnName().String()
}
