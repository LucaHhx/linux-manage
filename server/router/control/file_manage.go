package control

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FileManageRouter struct {
}

// InitFileManageRouter 初始化 FileManage表 路由信息
func (s *FileManageRouter) InitFileManageRouter(Router *gin.RouterGroup) {
	fileManageRouter := Router.Group("fileManage").Use(middleware.OperationRecord())
	fileManageRouterWithoutRecord := Router.Group("fileManage")
	var fileManageApi = v1.ApiGroupApp.ControlApiGroup.FileManageApi
	{
		fileManageRouter.POST("uploadFileChunk", fileManageApi.UploadFileChunk) //切片上传文件
		fileManageRouter.GET("abortFileUpload", fileManageApi.AbortFileUpload)  //取消切片上传
		fileManageRouter.PUT("option", fileManageApi.Option)                    // 配置管理
		fileManageRouter.DELETE("deleteItem", fileManageApi.DeleteItem)         // 删除文件
		fileManageRouter.PUT("renameItem", fileManageApi.RenameItem)            // 重命名文件
		fileManageRouter.PUT("copyItem", fileManageApi.CopyItem)                // 复制文件
		fileManageRouter.POST("createDirectory", fileManageApi.CreateDirectory) // 新建文件夹

	}
	{
		fileManageRouterWithoutRecord.POST("getFileItems", fileManageApi.GetFileItems) // 获取目录下文件
		fileManageRouterWithoutRecord.POST("downloadItem", fileManageApi.DownloadItem) // 新建文件夹
	}
}
