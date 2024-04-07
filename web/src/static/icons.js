// 动态导入目录下所有的 PNG 文件
const imageModules = import.meta.globEager('./icons/*.png')

// 将导入的模块转换为一个具有相同属性的对象
export const icons = Object.keys(imageModules).reduce((icons, path) => {
  // 提取文件名作为对象的键
  const key = path.split('/').pop().replace('.png', '')
  icons[key] = imageModules[path].default
  return icons
}, {})

