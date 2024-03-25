<template>
  <div>
    <div class="gva-form-box">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="bool字段:" prop="bool">
          <el-switch v-model="formData.bool" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable />
        </el-form-item>
        <el-form-item label="type字段:" prop="type">
          <el-input v-model.number="formData.type" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="int32字段:" prop="int32">
          <el-input v-model.number="formData.int32" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="int64字段:" prop="int64">
          <el-input v-model.number="formData.int64" :clearable="true" placeholder="请输入" />
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
          <el-input-number v-model="formData.float" :precision="2" :clearable="true" />
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
  createAText,
  updateAText,
  findAText
} from '@/api/aText'

defineOptions({
  name: 'ATextForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
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

const elFormRef = ref()

// 初始化方法
const init = async() => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findAText({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.reaText
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
