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

        <!-- <DxHeaderFilter :visible="true" /> -->
        <DxColumn data-field="ID" data-type="number" width="100" />
        <DxColumn data-field="bool" data-type="boolean" width="100" />
        <DxColumn data-field="type" cell-template="cellTemplate" is-required width="100">
          <DxLookup :data-source="typeOpts" display-expr="value" value-expr="key" />
        </DxColumn>
        <DxColumn data-field="int32" data-type="number" width="100" />
        <DxColumn data-field="int64" data-type="number" width="100" />
        <DxColumn data-field="varchar" data-type="string" width="100" />
        <DxColumn data-field="text" data-type="string" width="100" />
        <DxColumn data-field="tinyText" data-type="string" width="100" />
        <DxColumn data-field="mediumText" data-type="string" width="100" />
        <DxColumn data-field="longText" data-type="string" width="100" />
        <DxColumn data-field="float" data-type="number" width="100" />
        <DxColumn data-field="CreatedAt" data-type="datetime" width="150" />
        <DxColumn data-field="UpdatedAt" data-type="datetime" width="150" />

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
            title="Add aText"
          />
          <DxForm :show-validation-summary="true">
            <FDxItem data-field="ID" data-type="number" :editor-options="{disabled:true}" />
            <FDxItem data-field="bool" data-type="boolean" />
            <FDxItem data-field="type" data-type="number" />
            <FDxItem data-field="int32" data-type="number" is-required />
            <FDxItem data-field="int64" data-type="number" />
            <FDxItem data-field="varchar" data-type="string" />
            <FDxItem data-field="text" data-type="string" />
            <FDxItem data-field="tinyText" data-type="string" />
            <FDxItem data-field="mediumText" data-type="string" />
            <FDxItem data-field="longText" data-type="string" />
            <FDxItem data-field="float" data-type="number" />
            <FDxItem data-field="CreatedAt" data-type="datetime" />
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
          <el-tag :v-slot="data.displayValue" :type="tagStyle[data.value].type" :effect="tagStyle[data.value].effect" />
        </template>
      </DxDataGrid>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref } from 'vue'
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
  DxGroupItem
} from 'devextreme-vue/data-grid'
import {
  DxItem as FDxItem,
} from 'devextreme-vue/form'
import DxButton from 'devextreme-vue/button'
import DxDropDownButton from 'devextreme-vue/drop-down-button'
import { confirm } from 'devextreme/ui/dialog'
import CustomStore from 'devextreme/data/custom_store'
import { tagStyle, isNullOrUndefined, getGridCacheOperate, pageSizes } from '@/utils/dataGrid'
import {
  createAText,
  deleteAText,
  deleteATextByIds,
  updateAText,
  findAText,
  getATextList
} from '@/api/aText'
import { websocket } from '@/protobuf/websocket/register'
import { userMessage } from '@/protobuf/proto/room'
// 选项
const typeOpts = [
  { key: 1, value: 'Type1' },
  { key: 2, value: 'Type2' },
  { key: 3, value: 'Type3' },
  { key: 4, value: 'Type4' }
]
const aa = userMessage.encode({ name: 'huang', content: 'hello' })

websocket.notify('room.message', aa.finish())
// 缓存
const gridCacheKey = 'aTextGrid'
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
      const result = await getATextList({
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
    console.log(value)

    if (isNullOrUndefined(value.type)) {
      throw new Error('Type is required')
    }

    if (isNullOrUndefined(value.int64)) {
      throw new Error('int64 is required')
    }

    const result = await createAText(value)
    if (result.code !== 0) {
      throw new Error(result.msg)
    }
  },
  // 按key删除
  async remove(key) {
    const result = await deleteAText({ ID: key })
    if (result.code !== 0) {
      throw new Error(result.msg)
    }
  },
  // 更新数据
  async update(key, value) {
    const result = await updateAText({ id: key, fields: value })
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
  const result = await findAText({ ID: e.key })
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
      const result = await deleteATextByIds({ IDs: ids })
      if (result.code === 0) {
        grid.value.instance.refresh()
      }
    }
  })
}

</script>
<style scoped></style>
