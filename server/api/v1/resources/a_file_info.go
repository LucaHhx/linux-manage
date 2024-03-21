package resources

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/resources"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AFileInfoApi struct {
}

var aFileInfoService = service.ServiceGroupApp.ResourcesServiceGroup.AFileInfoService

// CreateAFileInfo 创建aFileInfo表
// @Tags AFileInfo
// @Summary 创建aFileInfo表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body resources.AFileInfo true "创建aFileInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /aFileInfo/createAFileInfo [post]
func (aFileInfoApi *AFileInfoApi) CreateAFileInfo(c *gin.Context) {
	var aFileInfo resources.AFileInfo
	err := c.ShouldBindJSON(&aFileInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := aFileInfoService.CreateAFileInfo(&aFileInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAFileInfo 删除aFileInfo表
// @Tags AFileInfo
// @Summary 删除aFileInfo表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body resources.AFileInfo true "删除aFileInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /aFileInfo/deleteAFileInfo [delete]
func (aFileInfoApi *AFileInfoApi) DeleteAFileInfo(c *gin.Context) {
	ID := c.Query("ID")
	if err := aFileInfoService.DeleteAFileInfo(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAFileInfoByIds 批量删除aFileInfo表
// @Tags AFileInfo
// @Summary 批量删除aFileInfo表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /aFileInfo/deleteAFileInfoByIds [delete]
func (aFileInfoApi *AFileInfoApi) DeleteAFileInfoByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := aFileInfoService.DeleteAFileInfoByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAFileInfo 更新aFileInfo表
// @Tags AFileInfo
// @Summary 更新aFileInfo表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body resources.AFileInfo true "更新aFileInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /aFileInfo/updateAFileInfo [put]
func (aFileInfoApi *AFileInfoApi) UpdateAFileInfo(c *gin.Context) {
	var upID request.UpdateById
	err := c.ShouldBindJSON(&upID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := aFileInfoService.UpdateAFileInfo(upID); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAFileInfo 用id查询aFileInfo表
// @Tags AFileInfo
// @Summary 用id查询aFileInfo表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query resources.AFileInfo true "用id查询aFileInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /aFileInfo/findAFileInfo [get]
func (aFileInfoApi *AFileInfoApi) FindAFileInfo(c *gin.Context) {
	ID := c.Query("ID")
	if reaFileInfo, err := aFileInfoService.GetAFileInfo(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reaFileInfo": reaFileInfo}, c)
	}
}

// GetAFileInfoList 分页获取aFileInfo表列表
// @Tags AFileInfo
// @Summary 分页获取aFileInfo表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query resourcesReq.AFileInfoSearch true "分页获取aFileInfo表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /aFileInfo/getAFileInfoList [get]
func (aFileInfoApi *AFileInfoApi) GetAFileInfoList(c *gin.Context) {
	var pageInfo request.ListSearch
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, groupCount, err := aFileInfoService.GetAFileInfoInfoList(pageInfo); err != nil {
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
