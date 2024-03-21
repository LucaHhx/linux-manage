import service from '@/utils/request'

export const getGridCache = (params) => {
  return service({
    url: '/grid/getGridCache',
    method: 'get',
    params
  })
}

export const setGridCache = (data) => {
  return service({
    url: '/grid/setGridCache',
    method: 'put',
    data
  })
}

export const delGridCache = (params) => {
  return service({
    url: '/grid/delGridCache',
    method: 'delete',
    params
  })
}
