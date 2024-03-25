import service from '@/utils/request'

// @Tags AText
// @Summary 创建aText表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AText true "创建aText表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /aText/createAText [post]
export const createAText = (data) => {
  return service({
    url: '/aText/createAText',
    method: 'post',
    data
  })
}

// @Tags AText
// @Summary 删除aText表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AText true "删除aText表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /aText/deleteAText [delete]
export const deleteAText = (params) => {
  return service({
    url: '/aText/deleteAText',
    method: 'delete',
    params
  })
}

// @Tags AText
// @Summary 批量删除aText表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除aText表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /aText/deleteAText [delete]
export const deleteATextByIds = (params) => {
  return service({
    url: '/aText/deleteATextByIds',
    method: 'delete',
    params
  })
}

// @Tags AText
// @Summary 更新aText表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AText true "更新aText表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /aText/updateAText [put]
export const updateAText = (data) => {
  return service({
    url: '/aText/updateAText',
    method: 'put',
    data
  })
}

// @Tags AText
// @Summary 用id查询aText表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AText true "用id查询aText表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /aText/findAText [get]
export const findAText = (params) => {
  return service({
    url: '/aText/findAText',
    method: 'get',
    params
  })
}

// @Tags AText
// @Summary 分页获取aText表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取aText表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /aText/getATextList [get]
export const getATextList = (data) => {
  return service({
    url: '/aText/getATextList',
    method: 'post',
    data
  })
}
