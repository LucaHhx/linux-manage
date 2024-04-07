<template>
  <div>

    <DxFileManager
      ref="fileManager"
      :file-system-provider="customFileProvider"
      :height="450"
      :current-path="options.mainFolder.key"
      current-path-keys=""
      :root-folder-name="options.mainFolder.name"
      :upload="{chunkSize: 1024 * 1024 * 10}"
      :on-selected-file-opened="displayImagePopup"
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
        <DxItem widget="dxDropDownButton" location="before" :options="favoriteFolders" />

        <DxItem name="refresh" />
        <DxItem name="separator" location="after" />
        <DxItem name="switchView" />
        <DxItem widget="dxDropDownButton" location="after" :options="attributesOperate" />
      <!-- <DxItem name="separator" location="after" /> -->
      <!-- <DxItem widget="dxDropDownButton" location="after" :options="gridCacheOperate" /> -->

      <!-- <DxFileSelectionItem name="download" /> -->
      <!-- <DxFileSelectionItem name="separator" />
      <DxFileSelectionItem name="move" />
      <DxFileSelectionItem name="copy" />
      <DxFileSelectionItem name="rename" />
      <DxFileSelectionItem name="separator" />
      <DxFileSelectionItem name="delete" />
      <DxFileSelectionItem name="clearSelection" />
      <DxFileSelectionItem name="separator" location="after" />
      <DxFileSelectionItem name="refresh" /> -->
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
          :value="options.hide"
          :text="data.text"
          readonly="true"
        />
      </template>

    </DxFileManager>
    <DxPopup
      v-model:visible="openFileContent.visible"
      v-model:title="openFileContent.title"
      :hide-on-outside-click="true"
      :show-close-button="true"
      max-height="600"
      class="photo-popup-content"
    >
      <div>
        <!-- <div v-if="openFileContent.type === '.png'">弹窗是打开的，文件类型是PNG</div> -->
        <el-image v-if="openFileContent.type==='.png'" :src="openFileContent.content" :fit="fit" />
        <DxTextArea
          v-if="openFileContent.type!=='.png'"
          v-model:value="openFileContent.content"
          :auto-resize-enabled="true"
          :max-length="maxLength"
          :input-attr="{ 'aria-label': 'Notes' }"
        />
      </div>

      <!-- <img
        v-if="openFileContent.type==='.png'"
        :src="openFileContent.content"
        class="photo-popup-image"
      > -->
    </DxPopup>
  </div>
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
import CustomFileSystemProvider from 'devextreme/file_management/custom_provider'
import DxTextArea from 'devextreme-vue/text-area'
import { DxCheckBox } from 'devextreme-vue/check-box'
import { DxPopup } from 'devextreme-vue/popup'
// import {
//   renameItem,
//   createDirectory,
//   copyItem,
//   addTag,
//   getMainList,
//   getRoot,
//   upAttribute
// } from '@/api/aFileManage'

import {
  getFileItems,
  upOption,
  uploadFileChunk,
  abortFileUpload,
  deleteItem,
  renameItem,
  createDirectory,
  copyItem,
  downloadItem
} from '@/api/fileManage'

import {
  saveAFileInfo,
  getTagAFileInfoList
} from '@/api/aFileInfo'
import { fileErrOther, getFileMD5, openDownload } from '@/utils/fileManage'
import { reactive, ref } from 'vue'
import { onServerPrefetch, onMounted, onBeforeMount } from 'vue'
const key = 'aFileManage'
const options = ref({
  mainFolder: {
    key: '/',
    name: 'Root'
  },
  hide: true
})
upOption({ key: key, type: 'load', option: options.value }).then((res) => {
  if (res.data) {
    options.value = res.data
  }
})

// 收藏文件夹列表
const favoriteFolders = ref()

const favoriteFolderClick = (e) => {
  options.value.mainFolder = e.itemData.data
}
getTagAFileInfoList({ tag: 'Collect' }).then((res) => {
  if (res.code !== 0) {
    throw new Error(res.msg)
  }
  const foldList = {
  // text: '目录',
    icon: 'src/static/icons/folder.png',
    dropDownOptions: { width: 150 },
    items: [{ text: 'Root', data: { key: '/', name: 'Root' }, onClick: favoriteFolderClick }]
  }
  res.data.forEach((item) => {
    foldList.items.push({
      text: item.name,
      icon: item.thumbnail,
      data: item,
      onClick: favoriteFolderClick
    })
  })
  console.log(foldList)
  favoriteFolders.value = foldList
})

const customFileProvider = new CustomFileSystemProvider({
  // 获取文件夹内容
  async getItems(parentDirectory) {
    return await getFileItems({ path: parentDirectory.key === '' ? options.value.mainFolder.key : parentDirectory.key,
      options: options.value

    }).then((res) => {
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
    console.log(parentDirectory?.dataItem)
    return await createDirectory({ path: parentDirectory?.dataItem === undefined ? options.value.mainFolder.key : parentDirectory.dataItem.key, name: name }).then((res) => {
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
  },
  async downloadItems(items) {
    items.forEach(async(item) => {
      console.log(item.dataItem)
      await downloadItem(item.dataItem).then((res) => {
        console.log(Object.prototype.toString.call(res))
        console.log(res)
        if (res.status !== 200) {
          throw new Error(res.msg)
        }
        const filename = res.headers.get('Content-Disposition').split('filename=')[1]
        console.log(filename)
        openDownload(res.data, filename)
      })
    })

    // const list = items.map((item) => {
    //   return item.dataItem
    // })
    // console.log(list)
  },
  async getItemsContent(items) {
    console.log('getItemsContent', items)
  }

})

const onItemClick = ({ itemData, viewArea, fileSystemItem }) => {
  if (itemData.onSend) {
    itemData.onSend({ itemData, viewArea, fileSystemItem })
  }
}

const onContextMenuShowing = (e) => {
  const items = []
  e.component.option().contextMenu.items.forEach(element => {
    if (!element?.text) {
      items.push(element)
    }
  })
  if (!e.fileSystemItem.isDirectory) {
    // items = ['create', 'upload', 'refresh', 'download', 'move', 'separator', 'copy', 'rename', 'delete']
  } else {
    items.push({
      text: '收藏',
      icon: 'src/static/icons/collection.png',
      beginGroup: true,
      onSend: async({ itemData, viewArea, fileSystemItem }) => {
        fileSystemItem.dataItem.tag = 'Collect'
        await saveAFileInfo(fileSystemItem.dataItem).then((res) => {
          if (res.code !== 0) {
            throw new Error(res.msg)
          }
          console.log(res)
        })
      }
    })
  }

  e.component.option('contextMenu.items', items)
}

const attributesOperate = {
  // text: '目录',
  icon: 'src/static/icons/slider.vertical.3@2x.png',
  dropDownOptions: { width: 150 },
  closeMenuOnClick: false,
  items: [
    {
      text: '隐藏文件',
      template: 'list-item',
      onClick() {
        options.value.hide = !options.value.hide
      }
    }, {
      text: '保存设置',
      icon: 'save',
      async onClick() {
        await upOption({ key: key, type: 'save', option: options.value }).then((res) => {
          if (res.code !== 0) {
            throw new Error(res.msg)
          }
        })
      }
    }, {
      text: '获取设置',
      icon: 'download',
      async onClick() {
        await upOption({ key: key, type: 'load', option: options.value }).then((res) => {
          if (res.data) {
            options.value = res.data
          }
        })
      }
    }, {
      text: '清除设置',
      icon: 'clear',
      async onClick() {
        await upOption({ key: key, type: 'rm', option: options.value }).then((res) => {
          if (res.code !== 0) {
            throw new Error(res.msg)
          }
          options.value = {
            mainFolder: {
              key: '/',
              name: 'Root'
            },
            hide: true
          }
        })
      }
    }]
}

const openFileContent = ref({
  visible: false,
  title: '',
  content: '',
  type: ''
})

const displayImagePopup = async(e) => {
  await downloadItem(e.file).then(async(res) => {
    if (res.status !== 200) {
      throw new Error(res.msg)
      // return
    }
    const filename = res.headers.get('Content-Disposition').split('filename=')[1]
    openFileContent.value.visible = true
    openFileContent.value.title = filename
    openFileContent.value.type = e.file.getFileExtension()
    if (isImageFile(openFileContent.value.type)) {
      openFileContent.value.content = URL.createObjectURL(res.data)
    } else {
      openFileContent.value.content = await res.data.text()
    }
  })
}

function isImageFile(extension) {
  // 定义图片文件的扩展名
  const imageExtensions = ['.jpg', '.jpeg', '.png', '.gif', '.bmp']

  // 判断文件扩展名是否在图片扩展名列表中
  return imageExtensions.includes(extension)
}
</script>
