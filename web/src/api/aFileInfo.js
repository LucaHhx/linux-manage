import service from '@/utils/request'

// @Tags AFileInfo
// @Summary 创建aFileInfo表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AFileInfo true "创建aFileInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /aFileInfo/createAFileInfo [post]
export const createAFileInfo = (data) => {
  return service({
    url: '/aFileInfo/createAFileInfo',
    method: 'post',
    data
  })
}

// @Tags AFileInfo
// @Summary 删除aFileInfo表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AFileInfo true "删除aFileInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /aFileInfo/deleteAFileInfo [delete]
export const deleteAFileInfo = (params) => {
  return service({
    url: '/aFileInfo/deleteAFileInfo',
    method: 'delete',
    params
  })
}

// @Tags AFileInfo
// @Summary 批量删除aFileInfo表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除aFileInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /aFileInfo/deleteAFileInfo [delete]
export const deleteAFileInfoByIds = (params) => {
  return service({
    url: '/aFileInfo/deleteAFileInfoByIds',
    method: 'delete',
    params
  })
}

// @Tags AFileInfo
// @Summary 更新aFileInfo表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AFileInfo true "更新aFileInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /aFileInfo/updateAFileInfo [put]
export const updateAFileInfo = (data) => {
  return service({
    url: '/aFileInfo/updateAFileInfo',
    method: 'put',
    data
  })
}

export const saveAFileInfo = (data) => {
  return service({
    url: '/aFileInfo/saveAFileInfo',
    method: 'put',
    data
  })
}

// @Tags AFileInfo
// @Summary 用id查询aFileInfo表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AFileInfo true "用id查询aFileInfo表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /aFileInfo/findAFileInfo [get]
export const findAFileInfo = (params) => {
  return service({
    url: '/aFileInfo/findAFileInfo',
    method: 'get',
    params
  })
}

// @Tags AFileInfo
// @Summary 分页获取aFileInfo表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取aFileInfo表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /aFileInfo/getAFileInfoList [get]
export const getAFileInfoList = (data) => {
  return service({
    url: '/aFileInfo/getAFileInfoList',
    method: 'post',
    data
  })
}

export const getTagAFileInfoList = (params) => {
  return service({
    url: '/aFileInfo/getTagAFileInfoList',
    method: 'get',
    params
  })
}
