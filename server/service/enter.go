package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/Text"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/fileManage"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup     system.ServiceGroup
	ExampleServiceGroup    example.ServiceGroup
	FileManageServiceGroup fileManage.ServiceGroup
	TextServiceGroup       Text.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
