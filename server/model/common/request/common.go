package request

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/helper"
	"gorm.io/gorm"
	"strings"
)

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

// GetById Find by id structure
type GetById struct {
	ID int `json:"id" form:"id"` // 主键ID
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId uint `json:"authorityId" form:"authorityId"` // 角色ID
}

type Condition struct {
	Key       string `json:"key" form:"key"`
	Value     string `json:"value" form:"value"`
	Operation string `json:"operation" form:"operation"`
	Sort      int    `json:"sort" form:"sort"`
}

func (c Condition) DbCondition(db *gorm.DB) {
	jcMap := helper.GormModelJsonToColumn(db.Statement.Model)
	switch c.Operation {
	case "=":
		db.Where(fmt.Sprintf("%s = ?", jcMap[c.Key]), c.Value)
	case "<>":
		db.Where(fmt.Sprintf("%s <> ?", jcMap[c.Key]), c.Value)
	case "<":
		db.Where(fmt.Sprintf("%s < ?", jcMap[c.Key]), c.Value)
	case ">":
		db.Where(fmt.Sprintf("%s > ?", jcMap[c.Key]), c.Value)
	case "<=":
		db.Where(fmt.Sprintf("%s <= ?", jcMap[c.Key]), c.Value)
	case ">=":
		db.Where(fmt.Sprintf("%s >= ?", jcMap[c.Key]), c.Value)
	case "between":
		values := strings.Split(c.Value, ",")
		if len(values) == 2 {
			db.Where(fmt.Sprintf("%s between ? and ?", jcMap[c.Key]), values[0], values[1])
		}
		global.GVA_LOG.Error(fmt.Sprintf("%s between operation value error: %s", c.Key, c.Value))
	case "contains":
		db.Where(fmt.Sprintf("%s like ?", jcMap[c.Key]), fmt.Sprintf("%%%s%%", c.Value))
	case "notcontains":
		db.Where(fmt.Sprintf("%s not like ?", jcMap[c.Key]), fmt.Sprintf("%%%s%%", c.Value))
	case "startswith":
		db.Where(fmt.Sprintf("%s like ?", jcMap[c.Key]), fmt.Sprintf("%s%%", c.Value))
	case "endswith":
		db.Where(fmt.Sprintf("%s like ?", jcMap[c.Key]), fmt.Sprintf("%%%s", c.Value))
	default:
		if c.Value != "" {
			db.Where(fmt.Sprintf("%s = ?", jcMap[c.Key]), c.Value)
		}
	}
}

type Sort struct {
	Selector string `json:"selector" form:"selector"`
	Desc     bool   `json:"desc" form:"desc"`
}

func (s Sort) DbSort(db *gorm.DB) {
	jcMap := helper.GormModelJsonToColumn(db.Statement.Model)
	name := jcMap[s.Selector]
	if name == "" {
		name = s.Selector
	}
	if s.Desc {
		db.Order(fmt.Sprintf("%s desc", name))
	} else {
		db.Order(fmt.Sprintf("%s asc", name))
	}
}
