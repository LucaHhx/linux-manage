package helper

import (
	"reflect"
	"strings"
)

func GormModelJsonToColumn(i interface{}) map[string]string {
	t := reflect.TypeOf(i)

	// 如果是指针，获取指针指向的结构体类型
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		i = reflect.ValueOf(i).Elem().Interface()
	}

	result := make(map[string]string)
	result["ID"] = "id"
	result["CreatedAt"] = "created_at"
	result["UpdatedAt"] = "updated_at"
	result["DeletedAt"] = "deleted_at"
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		gormTag := field.Tag.Get("gorm")
		columnName := extractColumnName(gormTag)
		if jsonTag != "" && jsonTag != "-" && columnName != "" {
			result[jsonTag] = columnName
		}
	}
	return result
}

// extractColumnName 从gorm标签字符串中提取column名称
func extractColumnName(gormTag string) string {
	parts := strings.Split(gormTag, ";")
	for _, part := range parts {
		if strings.HasPrefix(part, "column:") {
			return strings.TrimPrefix(part, "column:")
		}
	}
	return ""
}
