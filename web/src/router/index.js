import Vue from 'vue'
import VueRouter from 'vue-router'

export const myimport = require('./_import_' + process.env.NODE_ENV)

Vue.use(VueRouter)

export const constantRouterMap = [{
  path: '/register',
  component: myimport('user/register'),
  hidden: false
},{
  path: '/login',
  component: myimport('user/login'),
  hidden: false
}, {
  path: '/auth',
  component: myimport('user/authredirect'),
  hidden: false
}, {
  path: '/info',
  component: myimport('user/info'),
  hidden: false
}, {
  path: '/address',
  component: myimport('user/address'),
  hidden: true
}, {
  path: '/agree',
  component: myimport('user/agree'),
  hidden: false
}, {
  path: '/confirm',
  component: myimport('order/index'),
  hidden: true
},{
  path: '/order',
  component: myimport('order/list'),
  hidden: true
}, {
  path: '/detail',
  component: myimport('order/detail'),
  hidden: true
}, {
  path: '/404',
  component: myimport('error/404'),
  hidden: false
}, {
  path: '/500',
  component: myimport('error/500'),
  hidden: false
}, {
  path: '/',
  component: myimport('home/index'),
  hidden: false,
  name: '首页',
}, {
  path: '/category',
  component: myimport('category/index'),
  hidden: false,
  name: '分类',
}, {
  path: '/search',
  component: myimport('search/index'),
  hidden: false,
  name: '搜索',
}, {
  path: '/product',
  component: myimport('product/index'),
  hidden: false,
  name: '商品'
}, {
  path: '/profile',
  component: myimport('user/profile'),
  hidden: true,
  name: '个人中心'
}]


//路由详细配置
export default new VueRouter({
  mode: 'history',
  routes: constantRouterMap,
  strict: process.env.NODE_ENV !== 'production',
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      if (from.meta.keepAlive) {
        from.meta.savedPosition = document.body.scrollTop;
      }
      return {
        x: 0,
        y: to.meta.savedPosition || 0
      }
    }
  }
})