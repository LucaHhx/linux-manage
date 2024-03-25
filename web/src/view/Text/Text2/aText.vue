<template>
  <div>
    <div class="content-block dx-card responsive-paddings">
      <DxDataGrid
        id="gridContainer"
        ref="grid"
        :data-source="store"
        key-expr="id"
        :show-borders="true"
        show-column-lines="true"
        :remote-operations="true"
        :focused-row-enabled="true"
        allow-column-resizing="true"
        show-row-lines="true"
        row-alternation-enabled="true"
        :selected-row-keys="selectedItemKeys"
        column-resizing-mode="widget"
        @selection-changed="onSelectionChanged"
      >

        <DxToolbar>
          <DxItem
            location="before"
            name="addRowButton"
          />
          <DxItem
            location="before"
          >
            <DxButton
              :disabled="!selectedItemKeys.length"
              icon="trash"
              text="Delete Selected Records"
              @click="deleteRecords"
            />
          </DxItem>
          <DxItem
            location="center"
            name="applyFilterButton"
            show-text="always"
          />
          <DxItem
            location="before"
            name="columnChooserButton"
          />
          <DxItem
            location="after"
            name="columnChooserButton"
          />
          <DxItem
            location="after"
            name="exportButton"
          />

        </DxToolbar>
        <!-- 存储状态 -->
        <DxStateStoring
          :enabled="true"
          type="sessionStorage"
          storage-key="storage"
        />
        <!-- 内部滚动条 -->
        <DxScrolling row-rendering-mode="virtual" />

        <!-- 导出excel -->
        <DxExport
          :enabled="true"
          :allow-export-selected-data="true"
        />
        <!-- 分组 -->
        <DxGroupPanel :visible="true" />
        <!-- 列选择 -->
        <DxColumnChooser :enabled="true" />
        <!-- 列固定 -->
        <DxColumnFixing :enabled="true" />
        <!-- 编辑模式弹框编辑 -->
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
            title="Employee Info"
          />
          <DxForm>
            <DxItem
              :col-count="2"
              :col-span="2"
              item-type="group"
            />
          </DxForm>
        </DxEditing>

        <!-- 列筛选 -->
        <DxFilterRow
          :visible="true"
          apply-filter="onClick"
          apply-filter-text="Query"
        />
        <DxFilterPanel :visible="true" :customize-text="customizetext" />
        <!-- 排序 -->
        <DxSorting mode="multiple" />

        <!-- 行选择 -->
        <DxSelection
          select-all-mode="page"
          show-check-boxes-mode="onClick"
          mode="multiple"
        />

        <DxPaging :page-size="10" />
        <DxPager
          :visible="true"
          :allowed-page-sizes="pageSizes"
          :display-mode="displayMode"
          :show-page-size-selector="showPageSizeSelector"
          :show-info="showInfo"
          :show-navigation-buttons="showNavButtons"
        />
      </DxDataGrid>
    </div>
    <button @click="onclick">aaaaaaa</button>
    <button @click="onclick2">BBBBBBBBB</button>
  </div>
</template>
<script setup lang="ts">
import { computed, ref } from 'vue'
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
  DxStateStoring,
  DxSelection,
  DxToolbar,
  DxItem,
  DxGroupPanel,
  DxColumnChooser,
  DxFilterPanel
} from 'devextreme-vue/data-grid'
import DxButton from 'devextreme-vue/button'
import DxSelectBox from 'devextreme-vue/select-box'
import DxCheckBox from 'devextreme-vue/check-box'
import { generateData } from './data.js'
import CustomStore from 'devextreme/data/custom_store'
import { GetColumnsFilter, GetGridSearch } from '@/utils/dataGrid'
import {
  createAText,
  deleteAText,
  deleteATextByIds,
  updateAText,
  findAText,
  getATextList
} from '@/api/aText'
const displayModes = [
  { text: 'Display Mode \'full\'', value: 'full' },
  { text: 'Display Mode \'compact\'', value: 'compact' },
]
const pageSizes = [5, 10, 'all']

const displayMode = ref('full')
const showPageSizeSelector = ref(true)
const showInfo = ref(true)
const showNavButtons = ref(true)
const grid = ref(null)
const isCompactMode = computed(() => displayMode.value === 'compact')

const columnsFilter = ref({})

const isNotEmpty = (value) => value !== undefined && value !== null && value !== ''

const store = new CustomStore({
  key: 'ID',
  async load(loadOptions) {
    try {
      console.log({ ...GetGridSearch(grid.value.instance), sorts: loadOptions.sort })

      const result = await getATextList({ ...GetGridSearch(grid.value.instance), sorts: loadOptions.sort })
      console.log(result)

      if (result.code !== 0) {
        throw new Error('Data Loading Error')
      }
      return {
        data: result.data.data,
        totalCount: result.data.totalCount,
        // summary: result.summary,
        // groupCount: result.groupCount,
      }
    } catch (err) {
      throw new Error('Data Loading Error')
    }
  },
  async insert(aa) {
  }
})

const onRefreshClick = () => {
  window.location.reload()
}

const getcols = () => {
  console.log(GetColumnsFilter(grid))
}

const selectedItemKeys = ref([])
const onSelectionChanged = (e) => {
  selectedItemKeys.value = e.selectedRowKeys
}

const deleteRecords = () => {
  console.log(selectedItemKeys.value)
}

const query = () => {
  grid.value.instance.refresh()
}
const customizetext = (e) => {
  // console.log(e.component.getVisibleColumns())

  // console.log(GetColumnsFilter(e.component.getVisibleColumns()))
}
const onclick = () => {
  const state = grid.value.instance.state()
  console.log(state)
}
const onclick2 = () => {
  grid.value.instance.state(null)
}
</script>
<style scoped>
</style>
