package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/Text"
	"github.com/flipped-aurora/gin-vue-admin/server/service/control"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/resources"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup    system.ServiceGroup
	ExampleServiceGroup   example.ServiceGroup
	TextServiceGroup      Text.ServiceGroup
	ControlServiceGroup   control.ServiceGroup
	ResourcesServiceGroup resources.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
