import router from './router'

import {
  hasToken
} from '@/utils/storage'

const  routes = router.options.routes
let blackList = []
//let whiteList = []

for (const item of routes) {
  if (item.hidden) {
    blackList.push(item.path)
  }
}

router.beforeEach((to, from, next) => {
  console.log("是否已经登录=", hasToken())

  if (hasToken()) {
    next()
  } else {
    if (blackList.indexOf(to.path) !== -1) {
      next({
        path: '/login'
      })
    } else {
      next()
    }
  }
})

router.afterEach(() => {
})
