package filter

import (
	"fmt"
	"reflect"
)

type Condition struct {
	Operator string //操作符
	Field    string //字段
	Value    string //值
}

type Expression struct {
	Operator   string
	Conditions interface{}
}

func ParsingExpression(data interface{}) *Expression {
	if data == nil {
		return nil
	}
	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		slice := data.([]interface{})
		if ok, con := isCondition(slice); ok {
			return &Expression{
				Operator:   slice[1].(string),
				Conditions: con,
			}
		}
		if len(slice) != 3 {
			return nil
		}

		exp := &Expression{
			Operator:   slice[1].(string),
			Conditions: make([]interface{}, 0),
		}

		for _, item := range []interface{}{slice[0], slice[2]} {
			if subExpr := ParsingExpression(item); subExpr != nil {
				exp.Conditions = append(exp.Conditions.([]interface{}), subExpr)
			}
		}

		return exp
	}
	return nil
}

func isCondition(data []interface{}) (bool, *Condition) {
	if len(data) == 3 && reflect.TypeOf(data[0]).Kind() == reflect.String {
		return true, &Condition{
			Operator: data[1].(string),
			Field:    data[0].(string),
			Value:    fmt.Sprintf("%v", data[2]),
		}
	}
	return false, nil
}

func valueToString(i interface{}) {

}
