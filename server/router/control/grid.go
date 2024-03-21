package control

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GridRouter struct{}

// InitGridRouter 初始化 grid表 路由信息
func (s *GridRouter) InitGridRouter(Router *gin.RouterGroup) {
	gridRouter := Router.Group("grid").Use(middleware.OperationRecord())
	gridRouterWithoutRecord := Router.Group("grid")
	var gridApi = v1.ApiGroupApp.ControlApiGroup.GridApi
	{
		gridRouter.DELETE("delGridCache", gridApi.DelGridCache) // 删除grid表
		gridRouter.PUT("setGridCache", gridApi.SetGridCache)    // 更新grid表
	}
	{
		gridRouterWithoutRecord.GET("getGridCache", gridApi.GetGridCache) // 根据ID获取grid表
	}
}
