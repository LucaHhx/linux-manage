package Text

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Text"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ATextApi struct {
}

var aTextService = service.ServiceGroupApp.TextServiceGroup.ATextService

// CreateAText 创建aText表
// @Tags AText
// @Summary 创建aText表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Text.AText true "创建aText表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /aText/createAText [post]
func (aTextApi *ATextApi) CreateAText(c *gin.Context) {
	var aText Text.AText
	err := c.ShouldBindJSON(&aText)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := aTextService.CreateAText(&aText); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAText 删除aText表
// @Tags AText
// @Summary 删除aText表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Text.AText true "删除aText表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /aText/deleteAText [delete]
func (aTextApi *ATextApi) DeleteAText(c *gin.Context) {
	ID := c.Query("ID")
	if err := aTextService.DeleteAText(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteATextByIds 批量删除aText表
// @Tags AText
// @Summary 批量删除aText表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /aText/deleteATextByIds [delete]
func (aTextApi *ATextApi) DeleteATextByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := aTextService.DeleteATextByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAText 更新aText表
// @Tags AText
// @Summary 更新aText表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Text.AText true "更新aText表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /aText/updateAText [put]
func (aTextApi *ATextApi) UpdateAText(c *gin.Context) {
	var upID request.UpdateById
	err := c.ShouldBindJSON(&upID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := aTextService.UpdateAText(upID); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAText 用id查询aText表
// @Tags AText
// @Summary 用id查询aText表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Text.AText true "用id查询aText表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /aText/findAText [get]
func (aTextApi *ATextApi) FindAText(c *gin.Context) {
	ID := c.Query("ID")
	if reaText, err := aTextService.GetAText(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reaText": reaText}, c)
	}
}

// GetATextList 分页获取aText表列表
// @Tags AText
// @Summary 分页获取aText表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query TextReq.ATextSearch true "分页获取aText表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /aText/getATextList [get]
func (aTextApi *ATextApi) GetATextList(c *gin.Context) {
	var pageInfo request.ListSearch
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, groupCount, err := aTextService.GetATextInfoList(pageInfo); err != nil {
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
