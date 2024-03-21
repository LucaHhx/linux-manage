package resources

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AFileInfoRouter struct {
}

// InitAFileInfoRouter 初始化 aFileInfo表 路由信息
func (s *AFileInfoRouter) InitAFileInfoRouter(Router *gin.RouterGroup) {
	aFileInfoRouter := Router.Group("aFileInfo").Use(middleware.OperationRecord())
	aFileInfoRouterWithoutRecord := Router.Group("aFileInfo")
	var aFileInfoApi = v1.ApiGroupApp.ResourcesApiGroup.AFileInfoApi
	{
		aFileInfoRouter.POST("createAFileInfo", aFileInfoApi.CreateAFileInfo)             // 新建aFileInfo表
		aFileInfoRouter.DELETE("deleteAFileInfo", aFileInfoApi.DeleteAFileInfo)           // 删除aFileInfo表
		aFileInfoRouter.DELETE("deleteAFileInfoByIds", aFileInfoApi.DeleteAFileInfoByIds) // 批量删除aFileInfo表
		aFileInfoRouter.PUT("updateAFileInfo", aFileInfoApi.UpdateAFileInfo)              // 更新aFileInfo表
	}
	{
		aFileInfoRouterWithoutRecord.GET("findAFileInfo", aFileInfoApi.FindAFileInfo)        // 根据ID获取aFileInfo表
		aFileInfoRouterWithoutRecord.POST("getAFileInfoList", aFileInfoApi.GetAFileInfoList) // 获取aFileInfo表列表
	}
}
