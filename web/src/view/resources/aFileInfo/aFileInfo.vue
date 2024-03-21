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

        <el-table-column sortable align="left" label="全路径" prop="key" width="120" />
        <el-table-column sortable align="left" label="路径" prop="path" width="120" />
        <el-table-column sortable align="left" label="名称" prop="name" width="120" />
        <el-table-column sortable align="left" label="修改日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.dateModified) }}</template>
        </el-table-column>
        <el-table-column sortable align="left" label="尺寸" prop="size" width="120" />
        <el-table-column sortable align="left" label="目录" prop="isDirectory" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.isDirectory) }}</template>
        </el-table-column>
        <el-table-column sortable align="left" label="有子目录" prop="hasSubDirectories" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.hasSubDirectories) }}</template>
        </el-table-column>
        <el-table-column sortable align="left" label="图标路径" prop="thumbnail" width="120" />
        <el-table-column sortable align="left" label="标签" prop="tag" width="120" />
        <el-table-column sortable align="left" label="主目录" prop="isMain" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.isMain) }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
              <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
              查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateAFileInfoFunc(scope.row)">变更</el-button>
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
        <el-form-item label="全路径:" prop="key">
          <el-input v-model="formData.key" :clearable="true" placeholder="请输入全路径" />
        </el-form-item>
        <el-form-item label="路径:" prop="path">
          <el-input v-model="formData.path" :clearable="true" placeholder="请输入路径" />
        </el-form-item>
        <el-form-item label="名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入名称" />
        </el-form-item>
        <el-form-item label="修改日期:" prop="dateModified">
          <el-date-picker v-model="formData.dateModified" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
        </el-form-item>
        <el-form-item label="尺寸:" prop="size">
          <el-input v-model.number="formData.size" :clearable="true" placeholder="请输入尺寸" />
        </el-form-item>
        <el-form-item label="目录:" prop="isDirectory">
          <el-switch v-model="formData.isDirectory" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable />
        </el-form-item>
        <el-form-item label="有子目录:" prop="hasSubDirectories">
          <el-switch v-model="formData.hasSubDirectories" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable />
        </el-form-item>
        <el-form-item label="图标路径:" prop="thumbnail">
          <el-input v-model="formData.thumbnail" :clearable="true" placeholder="请输入图标路径" />
        </el-form-item>
        <el-form-item label="标签:" prop="tag">
          <el-input v-model="formData.tag" :clearable="true" placeholder="请输入标签" />
        </el-form-item>
        <el-form-item label="主目录:" prop="isMain">
          <el-switch v-model="formData.isMain" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable />
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
        <el-descriptions-item label="全路径">
          {{ formData.key }}
        </el-descriptions-item>
        <el-descriptions-item label="路径">
          {{ formData.path }}
        </el-descriptions-item>
        <el-descriptions-item label="名称">
          {{ formData.name }}
        </el-descriptions-item>
        <el-descriptions-item label="修改日期">
          {{ formatDate(formData.dateModified) }}
        </el-descriptions-item>
        <el-descriptions-item label="尺寸">
          {{ formData.size }}
        </el-descriptions-item>
        <el-descriptions-item label="目录">
          {{ formatBoolean(formData.isDirectory) }}
        </el-descriptions-item>
        <el-descriptions-item label="有子目录">
          {{ formatBoolean(formData.hasSubDirectories) }}
        </el-descriptions-item>
        <el-descriptions-item label="图标路径">
          {{ formData.thumbnail }}
        </el-descriptions-item>
        <el-descriptions-item label="标签">
          {{ formData.tag }}
        </el-descriptions-item>
        <el-descriptions-item label="主目录">
          {{ formatBoolean(formData.isMain) }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createAFileInfo,
  deleteAFileInfo,
  deleteAFileInfoByIds,
  updateAFileInfo,
  findAFileInfo,
  getAFileInfoList
} from '@/api/aFileInfo'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'AFileInfo'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  key: '',
  path: '',
  name: '',
  dateModified: new Date(),
  size: 0,
  isDirectory: false,
  hasSubDirectories: false,
  thumbnail: '',
  tag: '',
  isMain: false,
})

// 验证规则
const rule = reactive({
  key: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  name: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
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
    key: 'key',
    path: 'path',
    name: 'name',
    dateModified: 'date_modified',
    size: 'size',
    isDirectory: 'is_directory',
    hasSubDirectories: 'has_sub_directories',
    thumbnail: 'thumbnail',
    tag: 'tag',
    isMain: 'is_main',
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
    if (searchInfo.value.isDirectory === '') {
      searchInfo.value.isDirectory = null
    }
    if (searchInfo.value.hasSubDirectories === '') {
      searchInfo.value.hasSubDirectories = null
    }
    if (searchInfo.value.isMain === '') {
      searchInfo.value.isMain = null
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
  const table = await getAFileInfoList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteAFileInfoFunc(row)
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
    const res = await deleteAFileInfoByIds({ IDs })
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
const updateAFileInfoFunc = async(row) => {
  const res = await findAFileInfo({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reaFileInfo
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteAFileInfoFunc = async(row) => {
  const res = await deleteAFileInfo({ ID: row.ID })
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
  const res = await findAFileInfo({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.reaFileInfo
    openDetailShow()
  }
}

// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
    key: '',
    path: '',
    name: '',
    dateModified: new Date(),
    size: 0,
    isDirectory: false,
    hasSubDirectories: false,
    thumbnail: '',
    tag: '',
    isMain: false,
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
    key: '',
    path: '',
    name: '',
    dateModified: new Date(),
    size: 0,
    isDirectory: false,
    hasSubDirectories: false,
    thumbnail: '',
    tag: '',
    isMain: false,
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createAFileInfo(formData.value)
        break
      case 'update':
        res = await updateAFileInfo(formData.value)
        break
      default:
        res = await createAFileInfo(formData.value)
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
