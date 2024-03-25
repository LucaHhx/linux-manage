package Text

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ATextRouter struct {
}

// InitATextRouter 初始化 aText表 路由信息
func (s *ATextRouter) InitATextRouter(Router *gin.RouterGroup) {
	aTextRouter := Router.Group("aText").Use(middleware.OperationRecord())
	aTextRouterWithoutRecord := Router.Group("aText")
	var aTextApi = v1.ApiGroupApp.TextApiGroup.ATextApi
	{
		aTextRouter.POST("createAText", aTextApi.CreateAText)             // 新建aText表
		aTextRouter.DELETE("deleteAText", aTextApi.DeleteAText)           // 删除aText表
		aTextRouter.DELETE("deleteATextByIds", aTextApi.DeleteATextByIds) // 批量删除aText表
		aTextRouter.PUT("updateAText", aTextApi.UpdateAText)              // 更新aText表
	}
	{
		aTextRouterWithoutRecord.GET("findAText", aTextApi.FindAText)        // 根据ID获取aText表
		aTextRouterWithoutRecord.POST("getATextList", aTextApi.GetATextList) // 获取aText表列表
	}
}
