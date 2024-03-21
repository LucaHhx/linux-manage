import { getGridCache, setGridCache, delGridCache } from '@/api/grid'
import notify from 'devextreme/ui/notify'
export function GetGridSearch(grid) {
  const search = {
    page: grid.pageIndex(),
    pageSize: grid.pageSize(),
    conditions: []
  }
  grid.getVisibleColumns().forEach(element => {
    if (!isNullOrUndefined(element.filterValue) || element.sortOrder) {
      // console.log(element)
      search.conditions.push({
        key: element.dataField,
        value: String(!isNullOrUndefined(element.filterValue) ? element.filterValue : ''),
        operation: element.selectedFilterOperation ? element.selectedFilterOperation : '',
      })
    }
  })
  return search
}

export function isNullOrUndefined(value) {
  return value === null || value === undefined
}

export const getGridCacheOperate = (grid, gridCacheKey) => {
  return [
    {
      text: '保存布局',
      icon: 'save',
      onClick: async() => {
        const state = grid.value.instance.state()
        const res = await setGridCache({ key: gridCacheKey, value: JSON.stringify(state) })
        if (res.code === 0) {
          notify({
            position: 'top center',
            message: '保存成功',
          }, 'success', 1000)
        }
      }
    }, {
      text: '获取布局',
      icon: 'download',
      onClick: async() => {
        const res = await getGridCache({ key: gridCacheKey })
        if (res.code === 0 && res.data) {
          console.log(res.data)
          const data = JSON.parse(res.data)
          grid?.value?.instance?.state(data)
        }
      }
    }, {
      text: '清除布局',
      icon: 'clear',
      onClick: async() => {
        grid?.value?.instance?.state(null)
      }
    }]
}

export const tagStyle = {
  0: { type: 'danger', effect: 'dark' },
  1: { type: 'primary', effect: 'dark' },
  2: { type: 'success', effect: 'dark' },
  3: { type: 'info', effect: 'dark' },
  4: { type: 'warning', effect: 'dark' },
  5: { type: 'danger', effect: 'dark' },
  6: { type: 'primary', effect: 'light' },
  7: { type: 'success', effect: 'light' },
  8: { type: 'info', effect: 'light' },
  9: { type: 'warning', effect: 'light' },
  10: { type: 'danger', effect: 'light' },
  11: { type: 'primary', effect: 'plain' },
  12: { type: 'success', effect: 'plain' },
  13: { type: 'info', effect: 'plain' },
  14: { type: 'warning', effect: 'plain' },
  15: { type: 'danger', effect: 'plain' },
}

export const pageSizes = [10, 20, 50, 100, 'all']
