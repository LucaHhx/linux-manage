export function GetColumnsFilter(colums) {
  const columnsFilter = {}
  colums.forEach(element => {
    if (element.filterValue || element.sortOrder) {
      columnsFilter[element.dataField] = {
        value: String(element.filterValue ? element.filterValue : ''),
        operation: element.selectedFilterOperation ? element.selectedFilterOperation : '',
        sort: element.sortOrder ? element.sortOrder : ''

      }
    }
  })
  return columnsFilter
}

export function GetGridSearch(grid) {
  const search = {
    page: grid.pageIndex(),
    pageSize: grid.pageSize(),
    conditions: []
  }
  grid.getVisibleColumns().forEach(element => {
    if (!isNullOrUndefined(element.filterValue) || element.sortOrder) {
      search.conditions.push({
        key: element.dataField,
        value: String(!isNullOrUndefined(element.filterValue) ? element.filterValue : ''),
        operation: element.selectedFilterOperation ? element.selectedFilterOperation : '',
      })
    }
  })
  return search
}

function isNullOrUndefined(value) {
  return value === null || value === undefined
}
