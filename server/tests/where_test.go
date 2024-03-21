package fileManage

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Text"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/filter"
	"github.com/goccy/go-json"
	"testing"
)

func TestWhere(t *testing.T) {
	core.Init()
	db := global.GVA_DB.Model(Text.AText{})
	var input interface{}
	str := `[[["type","=","3"],"and",["bool","=","0"]],"or",[["type","=","3"],"and",["bool","=","1"]]]`

	// 解析JSON字符串为interface{}
	err := json.Unmarshal([]byte(str), &input)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	filter.NewParsing(input, db).Apply()
	list := make([]Text.AText, 0)
	db.Find(&list)
	fmt.Println(list)
}

//type Condition struct {
//	Field    string
//	Operator string
//	Value    string
//}
//
//type Expression struct {
//	Operator string
//	Operands []interface{} // 可以是Condition也可以是*Expression
//}
//
//func parseExpression(data interface{}) *Expression {
//	switch reflect.TypeOf(data).Kind() {
//	case reflect.Slice:
//		slice := data.([]interface{})
//		if len(slice) == 3 && reflect.TypeOf(slice[1]).Kind() == reflect.String {
//			// 检测是否为Condition
//			left, okLeft := slice[0].([]interface{})
//			right, okRight := slice[2].([]interface{})
//			fmt.Println(left, right)
//			if !okLeft && !okRight {
//				return &Expression{
//					Operator: slice[1].(string),
//					Operands: []interface{}{parseCondition(slice)},
//				}
//			}
//
//			expr := &Expression{
//				Operator: slice[1].(string),
//				Operands: make([]interface{}, 0),
//			}
//			for _, item := range []interface{}{slice[0], slice[2]} {
//				if subExpr := parseExpression(item); subExpr != nil {
//					expr.Operands = append(expr.Operands, subExpr)
//				}
//			}
//			return expr
//		}
//	}
//	return nil
//}
//
//func parseCondition(data []interface{}) *Condition {
//	if len(data) != 3 {
//		return nil
//	}
//	return &Condition{
//		Field:    data[0].(string),
//		Operator: data[1].(string),
//		Value:    data[2].(string),
//	}
//}
