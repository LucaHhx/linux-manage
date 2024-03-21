<template>
  <div>
    <div class="gva-form-box">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="80px">
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
          <el-date-picker v-model="formData.dateModified" type="date" placeholder="选择日期" :clearable="true" />
        </el-form-item>
        <el-form-item label="尺寸:" prop="size">
          <el-input v-model.number="formData.size" :clearable="true" placeholder="请输入" />
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
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createAFileInfo,
  updateAFileInfo,
  findAFileInfo
} from '@/api/aFileInfo'

defineOptions({
  name: 'AFileInfoForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
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
  }],
  name: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
})

const elFormRef = ref()

// 初始化方法
const init = async() => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findAFileInfo({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.reaFileInfo
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
}

init()
// 保存按钮
const save = async() => {
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
    }
  })
}

// 返回按钮
const back = () => {
  router.go(-1)
}

</script>

<style>
</style>
