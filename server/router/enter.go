package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/Text"
	"github.com/flipped-aurora/gin-vue-admin/server/router/control"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/resources"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System    system.RouterGroup
	Example   example.RouterGroup
	Text      Text.RouterGroup
	Control   control.RouterGroup
	Resources resources.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
