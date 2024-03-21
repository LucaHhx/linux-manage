<template>
  <DxFileManager
    ref="fileManager"
    :file-system-provider="customFileProvider"
    :height="450"
    :current-path="currentItem.key"
    current-path-keys=""
    :root-folder-name="currentItem.name"
    :on-selected-file-opened="openFilePopup"
    :upload="{chunkSize: 1024 * 1024 * 10}"
    @context-menu-item-click="onItemClick"
    @context-menu-showing="onContextMenuShowing"
  >
    <DxPermissions
      :create="true"
      :copy="true"
      :move="true"
      :delete="true"
      :rename="true"
      :upload="true"
      :download="true"
    />
    <DxToolbar>

      <DxItem name="showNavPane" :visible="true" />
      <DxItem name="separator" />
      <DxItem name="create" />
      <DxItem name="upload" />
      <DxItem widget="dxMenu" location="before" :options="newFileMenuOptions" />
      <DxItem widget="dxDropDownButton" location="before" :options="attributesOptions" />
      <DxItem name="refresh" />
      <DxItem name="separator" location="after" />
      <DxItem name="switchView" />
      <DxItem widget="dxDropDownButton" location="after" :options="gridCacheOperate" />

      <DxFileSelectionItem name="download" />
      <DxFileSelectionItem name="separator" />
      <DxFileSelectionItem name="move" />
      <DxFileSelectionItem name="copy" />
      <DxFileSelectionItem name="rename" />
      <DxFileSelectionItem name="separator" />
      <DxFileSelectionItem name="delete" />
      <DxFileSelectionItem name="clearSelection" />
      <DxFileSelectionItem name="separator" location="after" />
      <DxFileSelectionItem name="refresh" />
    </DxToolbar>

    <DxItemView
      :show-parent-folder="false"
    >
      <DxDetails>
        <DxColumn data-field="thumbnail" />
        <DxColumn data-field="name" />
        <DxColumn
          data-field="tag"
          caption="Category"
          :width="95"
        />
        <DxColumn data-field="dateModified" />
        <DxColumn data-field="size" />
      </DxDetails>
    </DxItemView>
    <template #list-item="{ data }">
      <DxCheckBox
        :value="data.value"
        :text="data.text"
        readonly="true"
      />
    </template>
  </DxFileManager>
</template>

<script setup>
import {
  DxFileManager,
  DxPermissions,
  DxToolbar,
  DxItem,
  DxFileSelectionItem,
  DxDetails,
  DxColumn,
  DxItemView,
} from 'devextreme-vue/file-manager'
import DxDropDownButton from 'devextreme-vue/drop-down-button'
import CustomFileSystemProvider from 'devextreme/file_management/custom_provider'
import { DxCheckBox } from 'devextreme-vue/check-box'
import {
  getFileItems,
  uploadFileChunk,
  abortFileUpload,
  deleteItem,
  renameItem,
  createDirectory,
  copyItem,
  addTag,
  getMainList,
  getRoot,
  upAttribute
} from '@/api/aFileManage'
import { fileErr, fileErrOther, getFileMD5 } from '@/utils/fileManage'
import { ref } from 'vue'
import SparkMD5 from 'spark-md5'
import hide from '@/static/icons/hide.png'

const currentItem = ref({
  key: '/',
})
getRoot().then((res) => {
  currentItem.value = res.data
})
console.log(currentItem.value)
const fileManager = ref(null)

const getMainDir = (key) => {
  if (key === '') {
    return currentItem.value.key
  }
  return key
}
const rootFolderName = ref('/')
const customFileProvider = new CustomFileSystemProvider({
  // 获取文件夹内容
  async getItems(parentDirectory) {
    return await getFileItems({ parentDirectory: getMainDir(parentDirectory.key) }).then((res) => {
      if (res.code !== 0) {
        return fileErrOther(res.msg)
      }
      return res.data
    })
  },
  // 分片上传文件
  async uploadFileChunk(file, uploadInfo, destinationDirectory) {
    let fileMd5 = ''
    await getFileMD5(file).then((md5) => {
      fileMd5 = md5
    })
    let chunkMd5 = ''
    await getFileMD5(uploadInfo.chunkBlob).then((md5) => {
      chunkMd5 = md5
    })

    return await uploadFileChunk({
      file: uploadInfo.chunkBlob,
      fileMd5: fileMd5,
      fileName: file.name,
      chunkTotal: uploadInfo.chunkCount,
      chunkNumber: uploadInfo.chunkIndex,
      savePath: destinationDirectory.key,
      chunkMd5: chunkMd5
    }).then((res) => {
      if (res.code !== 0) {
        return fileErrOther(res.msg)
      }
    })
  },
  // 取消文件上传
  async  abortFileUpload(file, uploadInfo, destinationDirectory) {
    let fileMd5 = ''
    await getFileMD5(file).then((md5) => {
      fileMd5 = md5
    })
    return await abortFileUpload({ fileMd5: fileMd5, filePath: destinationDirectory.key }).then((res) => {
      if (res.code !== 0) {
        return fileErrOther(res.msg)
      }
    })
  },
  // 删除文件
  async deleteItem(item) {
    return await deleteItem({ path: item.key }).then((res) => {
      if (res.code !== 0) {
        return fileErrOther(res.msg)
      }
    })
  },
  // 重命名文件
  async renameItem(item, newName) {
    // console.log(item, newName)
    return await renameItem({ file: item.dataItem, newPath: item.dataItem.path, newName: newName }).then((res) => {
      if (res.code !== 0) {
        return fileErrOther(res.msg)
      }
    })
  },
  // 移动文件
  async moveItem(item, destinationDirectory) {
    return await renameItem({ file: item.dataItem, newPath: destinationDirectory.dataItem.key, newName: item.dataItem.name }).then((res) => {
      if (res.code !== 0) {
        return fileErrOther(res.msg)
      }
    })
  },
  // 创建文件夹
  async createDirectory(parentDirectory, name) {
    console.log(parentDirectory, name)
    return await createDirectory({ path: parentDirectory.dataItem.key, name: name }).then((res) => {
      console.log(res)
      if (res.code !== 0) {
        return fileErrOther(res.msg)
      }
    })
  },
  // 复制文件
  async copyItem(item, destinationDirectory) {
    return await copyItem({ file: item.dataItem.key, newPath: destinationDirectory.dataItem.key, newName: item.dataItem.name }).then((res) => {
      if (res.code !== 0) {
        return fileErrOther(res.msg)
      }
    })
  }

})
const onItemClick = ({ itemData, viewArea, fileSystemItem }) => {
  // console.log(itemData, viewArea, fileSystemItem)
  if (itemData.onItemClick) {
    itemData.onItemClick({ itemData, viewArea, fileSystemItem })
  }
  // console.log(itemData, viewArea, fileSystemItem)
}
const newFileMenuOptionsA = ref(null)

const getNewFileMenuOptions = async() => {
  return await getMainList().then((result) => {
    if (result.code !== 0) {
      throw new Error(result.msg)
    }
    const list = result.data.map((item) => {
      return {
        text: item.name,
        onItemClick: async(e) => {
          currentItem.value = e.itemData.data
          fileManager.value.instance.refresh()
        },
        data: item
      }
    })
    list.push({
      data: { key: '/' },
      text: 'Main',
      onItemClick: async(e) => {
        currentItem.value = { key: '/' }
        fileManager.value.instance.refresh()
      }
    })
    return {
      items: [
        {
          text: '目录',
          icon: 'bulletlist',
          items: list
        },
      ],
      onItemClick
    }
  })
}
const newFileMenuOptions = ref(null)
getNewFileMenuOptions().then((res) => {
  newFileMenuOptions.value = res
})

const onContextMenuShowing = (e) => {
  const fileContextMenuItems = ['create', 'upload', 'rename', 'move', 'copy', 'delete', 'refresh', 'download']
  const dirContextMenuItems = ['create', 'rename', 'move', 'copy', 'delete', 'refresh',
    {
      text: 'Category',
      items: [{
        text: 'Root',
        onItemClick: async(e) => {
          const file = e.fileSystemItem.dataItem
          file.tag = e.itemData.text
          const result = await addTag(file)
          if (result.code !== 0) {
            throw new Error(result.msg)
          }
          fileManager.value.instance.refresh()
        }
      }, {
        text: 'Main',
        onItemClick: async(e) => {
          const file = e.fileSystemItem.dataItem
          file.tag = e.itemData.text
          const result = await addTag(file)
          if (result.code !== 0) {
            throw new Error(result.msg)
          }
          fileManager.value.instance.refresh()
        }
      }, {
        text: 'clear',
        icon: 'clear',
        onItemClick: async(e) => {
          const file = e.fileSystemItem.dataItem
          file.tag = ''
          const result = await addTag(file)
          if (result.code !== 0) {
            throw new Error(result.msg)
          }
          fileManager.value.instance.refresh()
        }
      }],
      icon: 'tags',
    }]
  if (e?.fileSystemItem?.dataItem?.isDirectory) {
    // your code
    e.component.option('contextMenu.items', dirContextMenuItems)
  } else {
    // your code
    e.component.option('contextMenu.items', fileContextMenuItems)
  }
}
const attributesOptionOnItemClick = (e) => {
  console.log(e)
}

// {
//   text: '属性',
//   dropDownOptions: { width: 230 },
//   items: [
//     {
//       widget: 'dxCheckBox',
//       text: '隐藏文件',
//       key: 'attribute',
//       value: 'attribute',
//       template: 'list-item',
//       closeMenuOnClick: false
//     }
//   ],
// }
// 属性管理

const attributesOptions = ref({
  items: [
    {
      text: '目录',
      icon: 'bulletlist',
      items: [
        {
          widget: 'dxCheckBox',
          text: '隐藏文件',
          key: 'hiddenFiles',
          value: true,
          template: 'list-item',
          onItemClick: async(e) => {
            e.itemData.value = !e.itemData.value
            await upAttribute({ key: e.itemData.key, val: e.itemData.value })
            fileManager.value.instance.refresh()
          }
        }
      ]
    },
  ],
  onItemClick
})
const gridCacheOperate = {
  // text: '目录',
  icon: 'bulletlist',
  dropDownOptions: { width: 150 },
  items: [
    {
      text: '保存布局',
      icon: 'save',
    }, {
      text: '获取布局',
      icon: 'download',
    }, {
      text: '清除布局',
      icon: 'clear',
    }]
}
</script>
