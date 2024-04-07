package cache

import "github.com/flipped-aurora/gin-vue-admin/server/service/cache/control"

type CacheGroup struct {
	control.ControlGroup
}

var CacheExamples = new(CacheGroup)
