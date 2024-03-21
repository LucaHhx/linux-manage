package filter

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/helper"
	"gorm.io/gorm"
	"strings"
)

type Parsing struct {
	data   interface{}
	db     *gorm.DB
	e      *Expression
	jcMap  map[string]string
	query  string
	values []interface{}
}

func NewParsing(data interface{}, db *gorm.DB) *Parsing {
	return &Parsing{
		data:  data,
		db:    db,
		e:     ParsingExpression(data),
		jcMap: helper.GormModelJsonToColumn(db.Statement.Model),
	}
}

func (p *Parsing) Apply() {
	if p.e == nil {
		return
	}
	p.parsing()
	p.db = p.db.Where(p.query, p.values...)
}

func (p *Parsing) parsing() {
	switch c := p.e.Conditions.(type) {
	case *Condition:
		p.query, p.values = p.parsingCondition(c)
	case []interface{}:
		p.query, p.values = p.parsingExpressions(p.e)
	default:
		fmt.Println("Unknown type")
	}
}

func (p *Parsing) parsingCondition(c *Condition) (query string, value []interface{}) {
	query, value = "", make([]interface{}, 0)
	value = []interface{}{c.Value}
	switch c.Operator {
	case "=":
		query = fmt.Sprintf("`%s` = ?", p.jcMap[c.Field])
	case "<>":
		query = fmt.Sprintf("`%s` <> ?", p.jcMap[c.Field])
	case "<":
		query = fmt.Sprintf("`%s` < ?", p.jcMap[c.Field])
	case ">":
		query = fmt.Sprintf("`%s` > ?", p.jcMap[c.Field])
	case "<=":
		query = fmt.Sprintf("`%s` <= ?", p.jcMap[c.Field])
	case ">=":
		query = fmt.Sprintf("`%s` >= ?", p.jcMap[c.Field])
	case "between":
		if len(value) == 2 {
			value = []interface{}{value[0], value[1]}
			query = fmt.Sprintf("`%s` between ? and ?", p.jcMap[c.Field])
		}
	case "contains":
		query = fmt.Sprintf("`%s` like %%?%%", p.jcMap[c.Field])
	case "notcontains":
		query = fmt.Sprintf("`%s` not like  %%?%%", p.jcMap[c.Field])
	case "startswith":
		query = fmt.Sprintf("`%s` like ?%%", p.jcMap[c.Field])
	case "endswith":
		query = fmt.Sprintf("`%s` like %%?", p.jcMap[c.Field])
	default:
		if c.Value != "" {
			query = fmt.Sprintf("`%s` = ?", p.jcMap[c.Field])
		}
	}
	return
}

func (p *Parsing) parsingExpressions(c *Expression) (query string, value []interface{}) {
	switch c.Conditions.(type) {
	case *Condition:
		return p.parsingCondition(c.Conditions.(*Condition))
	}

	cs := c.Conditions.([]interface{})
	if len(cs) != 2 {
		return
	}
	query = fmt.Sprintf("(left) %s (right)", c.Operator)
	left := cs[0]
	right := cs[1]
	switch left.(type) {
	case *Condition:
		q, v := p.parsingCondition(left.(*Condition))
		query = strings.ReplaceAll(query, "left", q)
		value = append(value, v...)
	case *Expression:
		q, v := p.parsingExpressions(left.(*Expression))
		query = strings.ReplaceAll(query, "left", q)
		value = append(value, v...)
	}

	switch right.(type) {
	case *Condition:
		q, v := p.parsingCondition(right.(*Condition))
		query = strings.ReplaceAll(query, "right", q)
		value = append(value, v...)
	case *Expression:
		q, v := p.parsingExpressions(right.(*Expression))
		query = strings.ReplaceAll(query, "right", q)
		value = append(value, v...)
	}
	return
}
