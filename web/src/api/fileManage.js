import service from '@/utils/request'

// 获取文件列表
export const getFileItems = (data) => {
  return service({
    url: '/fileManage/getFileItems',
    method: 'post',
    data
  })
}

// 配置修改
export const upOption = (data) => {
  return service({
    url: '/fileManage/option',
    method: 'put',
    data
  })
}

// 上传文件
export const uploadFileChunk = (data) => {
  return service({
    url: '/fileManage/uploadFileChunk',
    method: 'post',
    donNotShowLoading: true,
    headers: { 'Content-Type': 'multipart/form-data' },
    data
  })
}

// 取消上传
export const abortFileUpload = (params) => {
  return service({
    url: '/fileManage/abortFileUpload',
    method: 'get',
    params
  })
}

// 删除文件
export const deleteItem = (params) => {
  return service({
    url: '/fileManage/deleteItem',
    method: 'delete',
    params
  })
}

// 重命名
export const renameItem = (data) => {
  return service({
    url: '/fileManage/renameItem',
    method: 'put',
    data
  })
}
// 创建文件夹
export const createDirectory = (data) => {
  return service({
    url: '/fileManage/createDirectory',
    method: 'post',
    data
  })
}
// 复制文件
export const copyItem = (data) => {
  return service({
    url: '/fileManage/copyItem',
    method: 'put',
    data
  })
}

// 复制文件
export const downloadItem = (data) => {
  return service({
    url: '/fileManage/downloadItem',
    responseType: 'blob', // 重要：设置响应类型为 'blob'
    method: 'post',
    data
  })
}
