package fileManage

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Text"
	"reflect"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Printf("%+v", GetJsonToColumn(Text.AText{}))

}

func parseGormTagToGetColumnName(gormTag string) string {
	parts := strings.Split(gormTag, ";")
	for _, part := range parts {
		if strings.HasPrefix(part, "column:") {
			return strings.TrimPrefix(part, "column:")
		}
	}
	return ""
}

func GetJsonToColumn(i interface{}) map[string]string {
	result := make(map[string]string)
	aTextType := reflect.TypeOf(i)
	for i := 0; i < aTextType.NumField(); i++ {
		field := aTextType.Field(i)
		jsonTag := field.Tag.Get("json")
		gormTag := field.Tag.Get("gorm")
		columnName := func(s string) string {
			parts := strings.Split(s, ";")
			for _, part := range parts {
				if strings.HasPrefix(part, "column:") {
					return strings.TrimPrefix(part, "column:")
				}
			}
			return ""
		}(gormTag)
		if jsonTag != "" && columnName != "" {
			result[jsonTag] = columnName
		}
	}
	return result
}
