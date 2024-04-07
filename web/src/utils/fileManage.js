
import FileSystemError from 'devextreme/file_management/error'
import SparkMD5 from 'spark-md5'
export function fileErr(code, message) {
  return new Promise((resolve, reject) => {
    reject(new FileSystemError(code, null, message))
  })
}

export function fileErrOther(message) {
  return new Promise((resolve, reject) => {
    reject(new FileSystemError(32767, null, message))
  })
}

export const getFileMD5 = (file) => {
  return new Promise((resolve, reject) => {
    const reader = new FileReader()

    reader.onload = function(e) {
      // 当文件完全读取完成后，计算其MD5
      const md5Hash = SparkMD5.ArrayBuffer.hash(e.target.result)
      resolve(md5Hash)
    }

    reader.onerror = function(error) {
      // 如果读取文件过程中发生错误，拒绝这个Promise
      reject(error)
    }

    // 读取文件内容，以便计算MD5
    reader.readAsArrayBuffer(file)
  })
}

export function openDownload(blob, fileName) {
  const url = window.URL.createObjectURL(blob)
  const a = document.createElement('a')
  a.href = url
  a.download = fileName.replace(/['"]/g, '') // 去掉可能的引号
  document.body.appendChild(a) // 这一行可选
  a.click()
  window.URL.revokeObjectURL(url)
  document.body.removeChild(a) // 这一行可选
}