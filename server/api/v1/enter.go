package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Text"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/fileManage"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup     system.ApiGroup
	ExampleApiGroup    example.ApiGroup
	FileManageApiGroup fileManage.ApiGroup
	TextApiGroup       Text.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
