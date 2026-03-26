const viewModules = import.meta.glob('../view/**/*.vue')
const pluginModules = import.meta.glob('../plugin/**/*.vue')

export const asyncRouterHandle = (asyncRouter) => {
  asyncRouter.forEach((item) => {
    if (item.component && typeof item.component === 'string') {
      // 特殊处理Layout组件
      if (item.component === 'Layout') {
        item.meta.path = '/src/view/layout/index.vue'
        item.component = viewModules['../view/layout/index.vue']
      } else {
        item.meta.path = '/src/' + item.component
        if (item.component.split('/')[0] === 'view') {
          item.component = dynamicImport(viewModules, item.component)
        } else if (item.component.split('/')[0] === 'plugin') {
          item.component = dynamicImport(pluginModules, item.component)
        }
      }
    }
    if (item.children) {
      asyncRouterHandle(item.children)
    }
  })
}

function dynamicImport(dynamicViewsModules, component) {
  const keys = Object.keys(dynamicViewsModules)
  const matchKeys = keys.filter((key) => {
    const k = key.replace('../', '')
    return k === component
  })
  const matchKey = matchKeys[0]

  if (matchKey) {
    return dynamicViewsModules[matchKey]
  } else {
    console.error(`Component not found: ${component}`)
    // 返回一个默认的错误组件
    return () => import('../view/error/index.vue')
  }
}
