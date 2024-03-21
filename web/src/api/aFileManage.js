import service from '@/utils/request'

export const getFileItems = (params) => {
  return service({
    url: '/aFileManage/getFileItems',
    method: 'get',
    params
  })
}

export const uploadFileChunk = (data) => {
  return service({
    url: '/aFileManage/uploadFileChunk',
    method: 'post',
    donNotShowLoading: true,
    headers: { 'Content-Type': 'multipart/form-data' },
    data
  })
}

export const abortFileUpload = (params) => {
  return service({
    url: '/aFileManage/abortFileUpload',
    method: 'get',
    params
  })
}

export const deleteItem = (params) => {
  return service({
    url: '/aFileManage/deleteItem',
    method: 'delete',
    params
  })
}

export const renameItem = (data) => {
  return service({
    url: '/aFileManage/renameItem',
    method: 'put',
    data
  })
}

export const createDirectory = (data) => {
  return service({
    url: '/aFileManage/createDirectory',
    method: 'post',
    data
  })
}

export const copyItem = (data) => {
  return service({
    url: '/aFileManage/copyItem',
    method: 'put',
    data
  })
}

export const addTag = (data) => {
  return service({
    url: '/aFileManage/addTag',
    method: 'put',
    data
  })
}

export const getMainList = (params) => {
  return service({
    url: '/aFileManage/getMainList',
    method: 'get',
    params
  })
}

export const getRoot = (params) => {
  return service({
    url: '/aFileManage/getRoot',
    method: 'get',
    params
  })
}

export const upAttribute = (data) => {
  return service({
    url: '/aFileManage/upAttribute',
    method: 'put',
    data
  })
}
