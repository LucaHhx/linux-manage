package resources

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ATorrentRouter struct {
}

// InitATorrentRouter 初始化 aTorrent表 路由信息
func (s *ATorrentRouter) InitATorrentRouter(Router *gin.RouterGroup) {
	aTorrentRouter := Router.Group("aTorrent").Use(middleware.OperationRecord())
	aTorrentRouterWithoutRecord := Router.Group("aTorrent")
	var aTorrentApi = v1.ApiGroupApp.ResourcesApiGroup.ATorrentApi
	{
		aTorrentRouter.POST("createATorrent", aTorrentApi.CreateATorrent)   // 新建aTorrent表
		aTorrentRouter.DELETE("deleteATorrent", aTorrentApi.DeleteATorrent) // 删除aTorrent表
		aTorrentRouter.DELETE("deleteATorrentByIds", aTorrentApi.DeleteATorrentByIds) // 批量删除aTorrent表
		aTorrentRouter.PUT("updateATorrent", aTorrentApi.UpdateATorrent)    // 更新aTorrent表
	}
	{
		aTorrentRouterWithoutRecord.GET("findATorrent", aTorrentApi.FindATorrent)        // 根据ID获取aTorrent表
		aTorrentRouterWithoutRecord.POST("getATorrentList", aTorrentApi.GetATorrentList)  // 获取aTorrent表列表
	}
}
