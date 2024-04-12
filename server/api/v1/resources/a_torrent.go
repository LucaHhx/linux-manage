package resources

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/resources"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/torrent"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type ATorrentApi struct {
}

var aTorrentService = service.ServiceGroupApp.ResourcesServiceGroup.ATorrentService

// CreateATorrent 创建aTorrent表
// @Tags ATorrent
// @Summary 创建aTorrent表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body resources.ATorrent true "创建aTorrent表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /aTorrent/createATorrent [post]
func (aTorrentApi *ATorrentApi) CreateATorrent(c *gin.Context) {
	var aTorrent resources.ATorrent
	err := c.ShouldBindJSON(&aTorrent)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := aTorrentService.CreateATorrent(&aTorrent); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteATorrent 删除aTorrent表
// @Tags ATorrent
// @Summary 删除aTorrent表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body resources.ATorrent true "删除aTorrent表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /aTorrent/deleteATorrent [delete]
func (aTorrentApi *ATorrentApi) DeleteATorrent(c *gin.Context) {
	ID := c.Query("ID")
	if err := aTorrentService.DeleteATorrent(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteATorrentByIds 批量删除aTorrent表
// @Tags ATorrent
// @Summary 批量删除aTorrent表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /aTorrent/deleteATorrentByIds [delete]
func (aTorrentApi *ATorrentApi) DeleteATorrentByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := aTorrentService.DeleteATorrentByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateATorrent 更新aTorrent表
// @Tags ATorrent
// @Summary 更新aTorrent表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body resources.ATorrent true "更新aTorrent表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /aTorrent/updateATorrent [put]
func (aTorrentApi *ATorrentApi) UpdateATorrent(c *gin.Context) {
	var upID request.UpdateById
	err := c.ShouldBindJSON(&upID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := aTorrentService.UpdateATorrent(upID); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindATorrent 用id查询aTorrent表
// @Tags ATorrent
// @Summary 用id查询aTorrent表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query resources.ATorrent true "用id查询aTorrent表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /aTorrent/findATorrent [get]
func (aTorrentApi *ATorrentApi) FindATorrent(c *gin.Context) {
	ID := c.Query("ID")
	if reaTorrent, err := aTorrentService.GetATorrent(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reaTorrent": reaTorrent}, c)
	}
}

// GetATorrentList 分页获取aTorrent表列表
// @Tags ATorrent
// @Summary 分页获取aTorrent表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query resourcesReq.ATorrentSearch true "分页获取aTorrent表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /aTorrent/getATorrentList [get]
func (aTorrentApi *ATorrentApi) GetATorrentList(c *gin.Context) {
	var pageInfo request.ListSearch
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, groupCount, err := aTorrentService.GetATorrentInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.DxDataGridPageResult{
			Data:       list,
			TotalCount: total,
			GroupCount: groupCount,
		}, "获取成功", c)
	}
}

func (aTorrentApi *ATorrentApi) Download(c *gin.Context) {
	ID := c.Query("ID")
	id, err := strconv.Atoi(ID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	task, err := torrent.NewDownloadTask(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = global.GVA_TORRENT.Download(task)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("下载成功", c)
}

func (aTorrentApi *ATorrentApi) StopDownload(c *gin.Context) {
	ID := c.Query("ID")
	id, err := strconv.Atoi(ID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = global.GVA_TORRENT.Stop(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("停止成功", c)
}
