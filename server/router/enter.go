package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/Text"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/fileManage"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System     system.RouterGroup
	Example    example.RouterGroup
	FileManage fileManage.RouterGroup
	Text       Text.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
