<template>
  <div id="app">
    <el-config-provider :locale="zhCn">
      <router-view />
    </el-config-provider>
  </div>
</template>

<script>
import deMessages from 'devextreme/localization/messages/zh.json'
import { locale, loadMessages } from 'devextreme/localization'

export default {
  created() {
    loadMessages(deMessages)
    locale(navigator.language)
  }
}

</script>

<script setup>
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'
// element 2.3.8之前使用下面的语句
// import zhCn from 'element-plus/lib/locale/lang/zh-cn'

defineOptions({
  name: 'App'
})

</script>
<style lang="scss">
@import "../dx-styles.scss";
@tailwind base;
@tailwind components;
@tailwind utilities;
// 引入初始化样式
#app {
  background: #eee;
  height: 100vh;
  overflow: hidden;
  font-weight: 400 !important;
}
.el-button{
  font-weight: 400 !important;
}
</style>
