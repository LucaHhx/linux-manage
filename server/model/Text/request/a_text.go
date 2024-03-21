package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ATextSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	request.PageInfo
	Conditions interface{}     `json:"conditions" form:"conditions"`
	Sorts      []request.Sort  `json:"sorts" form:"sorts"`
	Groups     []request.Group `json:"groups" form:"groups"`
}
