<template>
  <div>
    <div class="content-block dx-card responsive-paddings">
      <DxDataGrid
        id="gridContainer"
        ref="grid"
        :data-source="store"
        key-expr="ID"
        :show-borders="true"
        show-column-lines="true"
        :remote-operations="true"
        :focused-row-enabled="true"
        allow-column-resizing="true"
        show-row-lines="true"
        row-alternation-enabled="true"
        column-resizing-mode="widget"
        width="100%"
        allow-column-reordering="true"
        @editing-start="editingStart"
      >
        <DxLoadPanel :enabled="false" />
        <DxScrolling row-rendering-mode="virtual" /> <!-- 虚拟滚动 -->
        <DxExport :enabled="true" :allow-export-selected-data="true" />
        <DxGroupPanel :visible="true" /> <!-- 分组 -->
        <DxGrouping :auto-expand-all="false" />
        <DxRemoteOperations :group-paging="true" />
        <DxColumnChooser :enabled="true" />
        <DxColumnFixing :enabled="true" />

        <DxFilterPanel :visible="true" />
        <DxSorting mode="multiple" show-sort-indexes="false" />
        <DxSelection select-all-mode="page" show-check-boxes-mode="onClick" mode="multiple" />
        <DxFilterRow :visible="true" apply-filter-text="Filter" apply-filter="onClick" />
        <DxToolbar>
          <DxItem location="before" name="addRowButton" />
          <DxItem location="before">
            <DxButton
              icon="trash"
              text="Delete Selected Records"
              @click="deleteRecords"
            />
          </DxItem>
          <DxItem location="before" name="groupPanel" show-text="always" />
          <DxItem location="center" name="applyFilterButton" show-text="always" />
          <DxItem location="after">
            <DxButton icon="refresh" @click="refreshDataGrid" />
          </DxItem>
          <DxItem location="after" name="columnChooserButton" />
          <DxItem location="after" name="exportButton" />
          <DxItem location="after">
            <DxDropDownButton
              :items="gridCacheOperate"
              :drop-down-options="{ width: 150 }"
              icon="detailslayout"
            />
          </DxItem>
        </DxToolbar>
        <DxColumn data-field="ID" data-type="number" width="100" />
        <DxColumn data-field="state" data-type="number" caption="状态">
          <DxLookup :data-source="stateOpts" display-expr="value" value-expr="key" />
        </DxColumn>
        <DxColumn data-field="name" data-type="string" caption="名称" />
        <DxColumn data-field="progress" data-type="number" caption="进度" cell-template="progress-bar" />
        <DxColumn data-field="magnet" data-type="string" caption="链接" />
        <DxColumn data-field="CreatedAt" data-type="datetime" width="150" caption="创建时间" />
        <DxColumn data-field="UpdatedAt" data-type="datetime" width="150" caption="最近修改" />
        <DxColumn type="buttons">
          <GDxButton name="save" css-class="my-class" />
          <GDxButton name="edit" />
          <GDxButton name="delete" />
          <GDxButton icon="download" @click="download" />
          <GDxButton icon="close" @click="stopdownload" />
        </DxColumn>
        <DxEditing
          :allow-updating="true"
          :allow-adding="true"
          :allow-deleting="true"
          mode="popup"
        >
          <DxPopup
            :show-title="true"
            :width="700"
            :height="525"
            title="Add ATorrent"
          />
          <DxForm :show-validation-summary="true">
            <FDxItem data-field="ID" data-type="number" :editor-options="{disabled:true}" />
            <FDxItem data-field="state" data-type="number" :editor-options="{disabled:true}" />
            <FDxItem data-field="name" data-type="string" />
            <FDxItem data-field="progress" data-type="number" />
            <FDxItem
              data-field="magnet"
              data-type="string"
              editor-type="dxTextArea"
              :col-span="2"
              :editor-options="{autoResizeEnabled:'true'}"
            />

          </DxForm>
        </DxEditing>

        <DxPaging :page-size="10" />
        <DxPager
          :visible="true"
          :allowed-page-sizes="pageSizes"
          display-mode="full"
          :show-page-size-selector="true"
          :show-info="true"
          :show-navigation-buttons="true"
        />

        <DxSummary>
          <DxGroupItem
            column="Id"
            summary-type="count"
          />
        </DxSummary>

        <template #cellTemplate="{ data }">
          <el-tag :type="tagStyle[data.value].type" :effect="tagStyle[data.value].effect">  {{ data.displayValue }} </el-tag>
        </template>
        <template #progress-bar="{ data }">
          <div>
            <DxProgressBar
              id="progress-bar-status"
              :min="0"
              :max="100"
              :value="downloadInfo[data.data.ID]?.progress"
              :status-format="downloadInfo[data.data.ID]?.statusFormat"
              width="90%"
            />
            <!-- <button @click=" console.log(data);">{{ data.value }}</button> -->
          </div>

        </template>
      </DxDataGrid>
    </div>
  </div>
</template>
<script setup lang="ts">
import { onUnmounted, ref } from 'vue'
import {
  DxDataGrid,
  DxScrolling,
  DxPager,
  DxPaging,
  DxColumnFixing,
  DxSorting,
  DxFilterRow,
  DxColumn,
  DxEditing,
  DxExport,
  DxSelection,
  DxToolbar,
  DxItem,
  DxGroupPanel,
  DxColumnChooser,
  DxFilterPanel,
  DxGrouping,
  DxLookup,
  DxRemoteOperations,
  DxForm,
  DxPopup,
  DxSummary,
  DxGroupItem,
  DxLoadPanel,
  DxButton as GDxButton
} from 'devextreme-vue/data-grid'
import {
  DxItem as FDxItem,
} from 'devextreme-vue/form'
import DxButton from 'devextreme-vue/button'
import { DxProgressBar } from 'devextreme-vue/progress-bar'
import DxDropDownButton from 'devextreme-vue/drop-down-button'
import { confirm } from 'devextreme/ui/dialog'
import CustomStore from 'devextreme/data/custom_store'
import { tagStyle, isNullOrUndefined, getGridCacheOperate, pageSizes } from '@/utils/dataGrid'

import {
  createATorrent,
  deleteATorrent,
  deleteATorrentByIds,
  updateATorrent,
  findATorrent,
  getATorrentList,
  startDownload,
  downloadInfo,
  stopDownload
} from '@/api/aTorrent'

// 选项
const stateOpts = [
  { key: 1, value: 'Type1' },
  { key: 2, value: 'Type2' },
  { key: 3, value: 'Type3' },
  { key: 4, value: 'Type4' }
]
// 缓存
const gridCacheKey = 'ATorrentGrid'
const grid = ref(null)

const gridCacheOperate = getGridCacheOperate(grid, gridCacheKey)
gridCacheOperate[1].onClick()

// 刷新
const refreshDataGrid = () => {
  grid.value.instance.refresh()
}

// 数据源
const store = new CustomStore({
  key: 'ID',
  async load(loadOptions) {
    try {
      const result = await getATorrentList({
        page: grid?.value?.instance?.pageIndex(),
        pageSize: grid?.value?.instance?.pageSize(),
        conditions: loadOptions.filter,
        sorts: loadOptions.sort,
        groups: loadOptions.group })

      if (result.code !== 0) {
        throw new Error('Data Loading Error')
      }
      return {
        data: result.data.data,
        totalCount: result.data.totalCount,
        groupCount: result.data.groupCount
      }
    } catch (err) {
      throw new Error('Data Loading Error')
    }
  },
  // 插入数据
  async insert(value) {
    if (isNullOrUndefined(value.name)) {
      throw new Error(' name is required')
    }
    if (isNullOrUndefined(value.magnet)) {
      throw new Error(' magnet is required')
    }

    const result = await createATorrent(value)
    if (result.code !== 0) {
      throw new Error(result.msg)
    }
  },
  // 按key删除
  async remove(key) {
    const result = await deleteATorrent({ ID: key })
    if (result.code !== 0) {
      throw new Error(result.msg)
    }
  },
  // 更新数据
  async update(key, value) {
    const result = await updateATorrent({ id: key, fields: value })
    if (result.code !== 0) {
      throw new Error(result.msg)
    }
  },
  async byKey(key, a) {
    console.log('ByKEY', key, a)
  }
})

// 编辑开始
const editingStart = async(e) => {
  // e.cancel = true
  const result = await findATorrent({ ID: e.key })
  if (result.code !== 0) {
    e.cancel = true
    throw new Error(result.msg)
  }
  e.data = result.data
}

// 批量删除
const deleteRecords = async() => {
  const result = confirm('<i>确定删除所选数据?</i>', '批量删除确认')
  result.then(async(dialogResult) => {
    if (dialogResult) {
      const ids = grid.value.instance.getSelectedRowKeys()
      const result = await deleteATorrentByIds({ IDs: ids })
      if (result.code === 0) {
        grid.value.instance.refresh()
      }
    }
  })
}

const download = (aa) => {
  aa.row.data.ID
  startDownload({ ID: aa.row.data.ID }).then((r) => {
    console.log(r)
  })
}

const stopdownload = (aa) => {
  aa.row.data.ID
  stopDownload({ ID: aa.row.data.ID }).then((r) => {
    console.log(r)
  })
}
const progressBarstatus = (r, v) => {
  console.log(r, v)
}
</script>
<style scoped></style>
