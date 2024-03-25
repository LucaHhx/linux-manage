<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
        <el-form-item label="创建日期" prop="createdAt">
          <template #label>
            <span>
              创建日期
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false" />
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false" />
        </el-form-item>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
      >
        <el-table-column type="selection" width="55" />

        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>

        <el-table-column sortable align="left" label="bool字段" prop="bool" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.bool) }}</template>
        </el-table-column>
        <el-table-column sortable align="left" label="type字段" prop="type" width="120" />
        <el-table-column sortable align="left" label="int32字段" prop="int32" width="120" />
        <el-table-column sortable align="left" label="int64字段" prop="int64" width="120" />
        <el-table-column sortable align="left" label="varchar字段" prop="varchar" width="120" />
        <el-table-column label="text字段" width="200">
          <template #default="scope">
            [富文本内容]
          </template>
        </el-table-column>
        <el-table-column label="tinyText字段" width="200">
          <template #default="scope">
            [富文本内容]
          </template>
        </el-table-column>
        <el-table-column label="mediumText字段" width="200">
          <template #default="scope">
            [富文本内容]
          </template>
        </el-table-column>
        <el-table-column label="longText字段" width="200">
          <template #default="scope">
            [富文本内容]
          </template>
        </el-table-column>
        <el-table-column sortable align="left" label="blob字段" prop="blob" width="120" />
        <el-table-column sortable align="left" label="tinyblob字段" prop="tinyblob" width="120" />
        <el-table-column sortable align="left" label="mediumBlob字段" prop="mediumBlob" width="120" />
        <el-table-column sortable align="left" label="longBlob字段" prop="longBlob" width="120" />
        <el-table-column sortable align="left" label="float字段" prop="float" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
              <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
              查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateATextFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-drawer v-model="dialogFormVisible" size="800" :show-close="false" :before-close="closeDialog">
      <template #title>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type==='create'?'添加':'修改' }}</span>
          <div>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>

      <el-form ref="elFormRef" :model="formData" label-position="top" :rules="rule" label-width="80px">
        <el-form-item label="bool字段:" prop="bool">
          <el-switch v-model="formData.bool" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable />
        </el-form-item>
        <el-form-item label="type字段:" prop="type">
          <el-input v-model.number="formData.type" :clearable="true" placeholder="请输入type字段" />
        </el-form-item>
        <el-form-item label="int32字段:" prop="int32">
          <el-input v-model.number="formData.int32" :clearable="true" placeholder="请输入int32字段" />
        </el-form-item>
        <el-form-item label="int64字段:" prop="int64">
          <el-input v-model.number="formData.int64" :clearable="true" placeholder="请输入int64字段" />
        </el-form-item>
        <el-form-item label="varchar字段:" prop="varchar">
          <el-input v-model="formData.varchar" :clearable="true" placeholder="请输入varchar字段" />
        </el-form-item>
        <el-form-item label="text字段:" prop="text">
          <RichEdit v-model="formData.text" />
        </el-form-item>
        <el-form-item label="tinyText字段:" prop="tinyText">
          <RichEdit v-model="formData.tinyText" />
        </el-form-item>
        <el-form-item label="mediumText字段:" prop="mediumText">
          <RichEdit v-model="formData.mediumText" />
        </el-form-item>
        <el-form-item label="longText字段:" prop="longText">
          <RichEdit v-model="formData.longText" />
        </el-form-item>
        <el-form-item label="blob字段:" prop="blob" />
        <el-form-item label="tinyblob字段:" prop="tinyblob" />
        <el-form-item label="mediumBlob字段:" prop="mediumBlob" />
        <el-form-item label="longBlob字段:" prop="longBlob" />
        <el-form-item label="float字段:" prop="float">
          <el-input-number v-model="formData.float" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer v-model="detailShow" size="800" :before-close="closeDetailShow" title="查看详情" destroy-on-close>
      <template #title>
        <div class="flex justify-between items-center">
          <span class="text-lg">查看详情</span>
        </div>
      </template>
      <el-descriptions :column="1" border>
        <el-descriptions-item label="bool字段">
          {{ formatBoolean(formData.bool) }}
        </el-descriptions-item>
        <el-descriptions-item label="type字段">
          {{ formData.type }}
        </el-descriptions-item>
        <el-descriptions-item label="int32字段">
          {{ formData.int32 }}
        </el-descriptions-item>
        <el-descriptions-item label="int64字段">
          {{ formData.int64 }}
        </el-descriptions-item>
        <el-descriptions-item label="varchar字段">
          {{ formData.varchar }}
        </el-descriptions-item>
        <el-descriptions-item label="text字段">
          [富文本内容]
        </el-descriptions-item>
        <el-descriptions-item label="tinyText字段">
          [富文本内容]
        </el-descriptions-item>
        <el-descriptions-item label="mediumText字段">
          [富文本内容]
        </el-descriptions-item>
        <el-descriptions-item label="longText字段">
          [富文本内容]
        </el-descriptions-item>
        <el-descriptions-item label="blob字段">
          {{ formData.blob }}
        </el-descriptions-item>
        <el-descriptions-item label="tinyblob字段">
          {{ formData.tinyblob }}
        </el-descriptions-item>
        <el-descriptions-item label="mediumBlob字段">
          {{ formData.mediumBlob }}
        </el-descriptions-item>
        <el-descriptions-item label="longBlob字段">
          {{ formData.longBlob }}
        </el-descriptions-item>
        <el-descriptions-item label="float字段">
          {{ formData.float }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createAText,
  deleteAText,
  deleteATextByIds,
  updateAText,
  findAText,
  getATextList
} from '@/api/aText'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'AText'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  bool: false,
  type: 0,
  int32: 0,
  int64: 0,
  varchar: '',
  text: '',
  tinyText: '',
  mediumText: '',
  longText: '',
  float: 0,
})

// 验证规则
const rule = reactive({
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
    bool: 'bool',
    type: 'type',
    int32: 'int32',
    int64: 'int64',
    varchar: 'varchar',
    text: 'text',
    tinyText: 'tiny_text',
    mediumText: 'medium_text',
    longText: 'long_text',
    blob: 'blob',
    tinyblob: 'tinyblob',
    mediumBlob: 'medium_blob',
    longBlob: 'long_blob',
    float: 'float',
  }

  let sort = sortMap[prop]
  if (!sort) {
    sort = prop.replace(/[A-Z]/g, match => `_${match.toLowerCase()}`)
  }

  searchInfo.value.sort = sort
  searchInfo.value.order = order
  getTableData()
}

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    if (searchInfo.value.bool === '') {
      searchInfo.value.bool = null
    }
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getATextList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteATextFunc(row)
  })
}

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const IDs = []
    if (multipleSelection.value.length === 0) {
      ElMessage({
        type: 'warning',
        message: '请选择要删除的数据'
      })
      return
    }
    multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
    const res = await deleteATextByIds({ IDs })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === IDs.length && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateATextFunc = async(row) => {
  const res = await findAText({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reaText
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteATextFunc = async(row) => {
  const res = await deleteAText({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 查看详情控制标记
const detailShow = ref(false)

// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}

// 打开详情
const getDetails = async(row) => {
  // 打开弹窗
  const res = await findAText({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.reaText
    openDetailShow()
  }
}

// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
    bool: false,
    type: 0,
    int32: 0,
    int64: 0,
    varchar: '',
    float: 0,
  }
}

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    bool: false,
    type: 0,
    int32: 0,
    int64: 0,
    varchar: '',
    float: 0,
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createAText(formData.value)
        break
      case 'update':
        res = await updateAText(formData.value)
        break
      default:
        res = await createAText(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      closeDialog()
      getTableData()
    }
  })
}

</script>

<style>

</style>
