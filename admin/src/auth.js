import router from './router'
import store from './store'
import NProgress from 'nprogress' // Progress 进度条
import 'nprogress/nprogress.css' // Progress 进度条样式
import {
  hasToken
} from '@/utils/storage'


// permissiom judge
// array.some() 是对数组中每一项运行指定函数，如果该函数对任一项返回true，则返回true
function hasPermission(roles, permissionRoles) {
  if (roles.indexOf('admin') >= 0) return true // admin权限 直接通过
  if (!permissionRoles) return true
  return roles.some(role => permissionRoles.indexOf(role) >= 0)
}

// register global progress.
const whiteList = ['/login', '/authredirect'] // 不重定向白名单

router.beforeEach((to, from, next) => {
  // 开启Progress
  NProgress.start()

  if (to.meta.title) {
    document.title = to.meta.title
  }

  // 判断是否有token
  if (hasToken()) {

    if (to.path === '/login') {
      next({
        path: '/'
      })
    } else {

      // 判断当前用户是否已拉取完admin信息
      if (!store.getters.roles || store.getters.roles.length === 0) {
        // 拉取admin
        store.dispatch('getAdmin').then(result => {
          store.dispatch('generateRoutes', result).then(() => {
            // 动态添加可访问路由表
            router.addRoutes(store.getters.addRouters)
              // hack方法 确保addRoutes已完成
            next({
              ...to,
              replace: true
            })
          })
        }).catch(() => {
          store.dispatch('LoginOut').then(res => {
            next({
              path: '/login'
            })
          })
        })

      } else {
        // 没有动态改变权限的需求可直接next() 删除下方权限判断 ↓
        if (hasPermission(store.getters.roles, to.meta.role)) {
          next()
        } else {
          next({
            path: '/401',
            query: {
              noGoBack: true
            }
          })
        }
        // 可删 ↑
      }
    }

  } else {

    // 在免登录白名单，直接进入
    if (whiteList.indexOf(to.path) !== -1) {
      next()
    } else {
      next('/login')
      NProgress.done()
    }

  }
})

router.afterEach(() => {
  NProgress.done() // 结束Progress
})