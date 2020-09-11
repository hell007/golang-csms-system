import Vue from 'vue'
import VueRouter from 'vue-router'
import Layout from '@/views/layout/Layout'
import View from '@/views/layout/View'

Vue.use(VueRouter)


const unit = {
  path: '/unit',
  component: Layout,
  hidden: false,
  name: '单元管理',
  icon: 'xa-icon xa-icon-copy',
  meta: {
    role: ['admin', 'editor']
  },
  dropdown: true,
  children: [{
    path: '/unit1',
    component: View,
    redirect: '/unit1/list',
    name: '页面测试',
    meta: {
      role: ['superadmin', 'admin']
    },
    children: [{
      path: 'list',
      component: () =>
        import ('@/views/unit/list'),
      name: '列表页',
      meta: {
        role: ['superadmin', 'admin']
      }
    }, {
      path: 'form',
      component: () =>
        import ('@/views/unit/form'),
      name: '表单页',
      meta: {
        role: ['superadmin', 'admin']
      }
    }, {
      path: 'form',
      component: () =>
        import ('@/views/unit/detail'),
      name: '详情页',
      meta: {
        role: ['superadmin', 'admin']
      }
    }]
  }, {
    path: '/unit2',
    component: View,
    redirect: '/unit2/list',
    name: 'UI组件',
    meta: {
      role: ['superadmin', 'admin']
    },
    children: [{
      path: 'list',
      component: () =>
        import ('@/views/unit/components'),
      name: '组件使用',
      meta: {
        role: ['superadmin', 'admin']
      }
    }]
  }]
}

const charts = {
  path: '/charts',
  component: Layout,
  redirect: '/d3',
  hidden: false,
  name: '数据可视化',
  icon: 'xa-icon xa-icon-data',
  noDropdown: true,
  children: [{
    path: '/d3',
    component: () =>
      import ('@/views/charts/d3'),
    name: 'd3'
  }]
}

export const constantRouter = [{
    path: '/login',
    component: () =>
      import ('@/views/login/index'),
    hidden: true
  }, {
    path: '/authredirect',
    component: () =>
      import ('@/views/login/authredirect'),
    hidden: true
  }, {
    path: '/404',
    component: () =>
      import ('@/views/error/404'),
    hidden: true
  }, {
    path: '/401',
    component: () =>
      import ('@/views/error/401'),
    hidden: true
  }, {
    path: '/',
    component: Layout,
    hidden: false,
    redirect: '/home',
    name: '首页',
    icon: 'xa-icon xa-icon-station',
    children: [{
      path: 'home',
      component: () =>
        import ('@/views/home/index')
    }]
  },
  unit,
  charts
]


//路由详细配置
export default new VueRouter({
  routes: constantRouter,
  //base: __dirname,//默认值: /，应用的基路径，一般就是项目的根目录，webpack中有配置好
  //linkActiveClass:'link-active',//默认值: router-link-active，就是当前组件被激活，相应路由会自动添加类”router-link-active”，这里是为了全局设置激活类名，如果不设置，直接用默认的也是可以的
  //注意:打开node后刷新二级页面空白，也影响保存更新
  //mode: 'history', //路由模式 默认hash值， history: 依赖 HTML5 History API 和服务器配置支持
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
  //通过这个这个属性（是个函数），可以让应用像浏览器的原生表现那样，在按下 后退/前进 按钮时，简单地让页面滚动到顶部或原来的位置
  //注意: 这个功能只在 HTML5 history 模式下可用
})

export const asyncRouter = [{
  path: '*',
  redirect: '/404',
  hidden: true
}]