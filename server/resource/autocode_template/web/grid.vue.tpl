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
        <DxColumn data-field="ID" data-type="number" width="100" />
        {{- range .Fields}}
          {{- if eq .FieldType "bool" }}
        <DxColumn data-field="{{.FieldJson}}"  data-type="boolean" caption="{{.FieldDesc}}"/>
          {{- end }}
          {{- if eq .FieldType "string" "richtext" }}
        <DxColumn data-field="{{.FieldJson}}"  data-type="string" caption="{{.FieldDesc}}"/>
          {{- end }}
          {{- if eq .FieldType "json" }}
        <DxColumn data-field="{{.FieldJson}}"  data-type="object" caption="{{.FieldDesc}}"/>
          {{- end }}
          {{- if eq .FieldType "int" "float64"  }}
        <DxColumn data-field="{{.FieldJson}}"  data-type="number" caption="{{.FieldDesc}}"/>
          {{- end }}
          {{- if eq .FieldType "time.Time" }}
        <DxColumn data-field="{{.FieldJson}}"  data-type="datetime" caption="{{.FieldDesc}}"/>
          {{- end }}
          {{- if eq .FieldType "float64" }}
        <DxColumn data-field="{{.FieldJson}}"  data-type="number" caption="{{.FieldDesc}}"/>
          {{- end }}
          {{- if eq .FieldType "enum" }}
        <DxColumn data-field="{{.FieldJson}}"  data-type="number" caption="{{.FieldDesc}}">
          <DxLookup :data-source="{{.FieldJson}}Opts" display-expr="value" value-expr="key" />
        </DxColumn>
          {{- end }}
          {{- else }}
        <DxColumn data-field="{{.FieldJson}}"  data-type="string" caption="{{.FieldDesc}}"/>
          {{- end }}


        <DxColumn data-field="CreatedAt" data-type="datetime" width="150" caption="创建时间"/>
        <DxColumn data-field="UpdatedAt" data-type="datetime" width="150" caption="最近修改"/>


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
            title="Add {{.StructName}}"
          />
          <DxForm :show-validation-summary="true">
            <FDxItem data-field="ID" data-type="number" :editor-options="{disabled:true}" />
{{- range .Fields}}
          {{- if eq .FieldType "bool" }}
            <FDxItem data-field="bool" data-type="boolean" />
          {{- end }}
          {{- if eq .FieldType "string" "richtext" }}
            <FDxItem data-field="{{.FieldJson}}"  data-type="string" />
          {{- end }}
          {{- if eq .FieldType "json" }}
        <FDxItem data-field="{{.FieldJson}}"  data-type="object" />
          {{- end }}
          {{- if eq .FieldType "int" "float64" "enum" }}
        <FDxItem data-field="{{.FieldJson}}"  data-type="number" />
          {{- end }}
          {{- if eq .FieldType "time.Time" }}
        <FDxItem data-field="{{.FieldJson}}"  data-type="datetime" />
          {{- end }}
          {{- if eq .FieldType "float64" }}
        <FDxItem data-field="{{.FieldJson}}"  data-type="number" />
          {{- end }}
          {{- else }}
        <FDxItem data-field="{{.FieldJson}}"  data-type="string" />
{{- end }}

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
          <el-tag :type="tagStyle[data.value].type" :effect="tagStyle[data.value].effect">  {{"{{"}} data.displayValue {{"}}"}}</el-tag>
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
  create{{.StructName}},
  delete{{.StructName}},
  delete{{.StructName}}ByIds,
  update{{.StructName}},
  find{{.StructName}},
  get{{.StructName}}List
} from '@/api/{{.PackageName}}'

// 选项
{{- range .Fields}}
  {{- if eq .FieldType "enum" }}
const {{.FieldJson}}Opts = [
  { key: 1, value: 'Type1' },
  { key: 2, value: 'Type2' },
  { key: 3, value: 'Type3' },
  { key: 4, value: 'Type4' }
]
  {{- end }}
{{- end }}

// 缓存
const gridCacheKey = '{{.StructName}}Grid'
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
      const result = await get{{.StructName}}List({
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
 {{- range .Fields }}
   {{- if eq .Require true }}
    if (isNullOrUndefined(value. {{.FieldJson }})) {
      throw new Error(' {{.FieldJson }} is required')
    }
    {{- end }}
{{- end }}

    const result = await create{{.StructName}}(value)
    if (result.code !== 0) {
      throw new Error(result.msg)
    }
  },
  // 按key删除
  async remove(key) {
    const result = await delete{{.StructName}}({ ID: key })
    if (result.code !== 0) {
      throw new Error(result.msg)
    }
  },
  // 更新数据
  async update(key, value) {
    const result = await update{{.StructName}}({ id: key, fields: value })
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
  const result = await find{{.StructName}}({ ID: e.key })
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
      const result = await delete{{.StructName}}ByIds({ IDs: ids })
      if (result.code === 0) {
        grid.value.instance.refresh()
      }
    }
  })
}

</script>
<style scoped></style>
