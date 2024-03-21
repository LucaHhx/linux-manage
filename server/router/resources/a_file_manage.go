package resources

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AFileManageRouter struct {
}

// InitAFileManageRouter 初始化 aFileManage表 路由信息
func (s *AFileManageRouter) InitAFileManageRouter(Router *gin.RouterGroup) {
	aFileManageRouter := Router.Group("aFileManage").Use(middleware.OperationRecord())
	aFileManageRouterWithoutRecord := Router.Group("aFileManage")
	var aFileManageApi = v1.ApiGroupApp.ResourcesApiGroup.AFileManageApi
	{
		aFileManageRouter.POST("createDirectory", aFileManageApi.CreateDirectory) // 新建aFileManage表
		//aFileManageRouter.DELETE("deleteAFileManage", aFileManageApi.DeleteAFileManage)           // 删除aFileManage表
		//aFileManageRouter.DELETE("deleteAFileManageByIds", aFileManageApi.DeleteAFileManageByIds) // 批量删除aFileManage表
		aFileManageRouter.PUT("renameItem", aFileManageApi.RenameItem)   // 重命名文件
		aFileManageRouter.PUT("copyItem", aFileManageApi.CopyItem)       // 重命名文件
		aFileManageRouter.PUT("addTag", aFileManageApi.AddTag)           // 重命名文件
		aFileManageRouter.PUT("upAttribute", aFileManageApi.UpAttribute) // 重命名文件
		aFileManageRouter.POST("uploadFileChunk", aFileManageApi.UploadFileChunk)
		aFileManageRouter.DELETE("deleteItem", aFileManageApi.DeleteItem)
	}
	{
		aFileManageRouterWithoutRecord.GET("getFileItems", aFileManageApi.GetFileItems)       // 根据ID获取aFileManage表
		aFileManageRouterWithoutRecord.GET("abortFileUpload", aFileManageApi.AbortFileUpload) // 根据ID获取aFileManage表
		aFileManageRouterWithoutRecord.GET("getMainList", aFileManageApi.GetMainList)         // 根据ID获取aFileManage表
		aFileManageRouterWithoutRecord.GET("getRoot", aFileManageApi.GetRoot)                 // 根据ID获取aFileManage表
	}
}
