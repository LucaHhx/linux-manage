package control

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/service/cache"
	"github.com/gin-gonic/gin"
)

type GridApi struct {
}

var GridService = service.ServiceGroupApp.ControlServiceGroup.GridService
var ServiceCacheGroup = cache.ServiceCacheGroup.ControlGrid

func (g *GridApi) GetGridCache(c *gin.Context) {
	key := c.Query("key")
	grid, _ := ServiceCacheGroup.GetControlGrid(key)
	response.OkWithDetailed(grid, "获取成功", c)
}

func (g *GridApi) SetGridCache(c *gin.Context) {
	gridCache := struct {
		Key   string `json:"key" form:"key"`
		Value string `json:"value" form:"value"`
	}{}
	err := c.ShouldBindJSON(&gridCache)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = ServiceCacheGroup.SetControlGrid(gridCache.Key, gridCache.Value)
	if err != nil {
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithMessage("设置成功", c)
	}
}

func (g *GridApi) DelGridCache(c *gin.Context) {
	key := c.Query("key")
	err := ServiceCacheGroup.DelControlGrid(key)
	if err != nil {
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
