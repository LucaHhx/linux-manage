import { protos } from '@/utils/proto'
import service from '@/utils/request'

// @Tags ATorrent
// @Summary 创建aTorrent表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ATorrent true "创建aTorrent表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /aTorrent/createATorrent [post]
export const createATorrent = (data) => {
  return service({
    url: '/aTorrent/createATorrent',
    method: 'post',
    data
  })
}

// @Tags ATorrent
// @Summary 删除aTorrent表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ATorrent true "删除aTorrent表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /aTorrent/deleteATorrent [delete]
export const deleteATorrent = (params) => {
  return service({
    url: '/aTorrent/deleteATorrent',
    method: 'delete',
    params
  })
}

// @Tags ATorrent
// @Summary 批量删除aTorrent表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除aTorrent表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /aTorrent/deleteATorrent [delete]
export const deleteATorrentByIds = (params) => {
  return service({
    url: '/aTorrent/deleteATorrentByIds',
    method: 'delete',
    params
  })
}

// @Tags ATorrent
// @Summary 更新aTorrent表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ATorrent true "更新aTorrent表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /aTorrent/updateATorrent [put]
export const updateATorrent = (data) => {
  return service({
    url: '/aTorrent/updateATorrent',
    method: 'put',
    data
  })
}

// @Tags ATorrent
// @Summary 用id查询aTorrent表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ATorrent true "用id查询aTorrent表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /aTorrent/findATorrent [get]
export const findATorrent = (params) => {
  return service({
    url: '/aTorrent/findATorrent',
    method: 'get',
    params
  })
}

// @Tags ATorrent
// @Summary 分页获取aTorrent表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取aTorrent表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /aTorrent/getATorrentList [get]
export const getATorrentList = (data) => {
  return service({
    url: '/aTorrent/getATorrentList',
    method: 'post',
    data
  })
}

export const startDownload = (params) => {
  return service({
    url: '/aTorrent/download',
    method: 'get',
    params
  })
}

export const stopDownload = (params) => {
  return service({
    url: '/aTorrent/stopDownload',
    method: 'get',
    params
  })
}

import { ref } from 'vue'
export const downloadInfo = ref({})

export const getDownloadInfo = (params) => {
  var info = protos.resources.downloadInfo.fromJSON(params)
  if (info.isFinished) {
    downloadInfo.value[info.id] = null
  }
  // const statusFormat = (ratio) => `Loading: ${ratio * 100}%`
  info.statusFormat = (ratio) => `${ratio * 100}%                      ` + info.finishBytes + '⇆' + info.totalBytes + '↧' + info.speed
  downloadInfo.value[info.id] = info
  // console.log('getDownloadInfo', info)
}

getDownloadInfo({
  id: 1,
  name: 'test',
  isFinished: false,
  finishBytes: '100KB',
  totalBytes: '1000KB',
  speed: '100KB/s'
})
