package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Text"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/control"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/resources"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup    system.ApiGroup
	ExampleApiGroup   example.ApiGroup
	TextApiGroup      Text.ApiGroup
	ControlApiGroup   control.ApiGroup
	ResourcesApiGroup resources.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
