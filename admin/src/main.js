import Vue from 'vue'
import Element from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import App from './App'
import router from './router'
import store from './store'
import * as filters from './filters' // 全局filter
import './auth' // 权限
// import './mock' // 该项目所有请求使用mockjs模拟

//加载模块
Vue.use(Element)

//Vue filter介绍及其使用
Object.keys(filters).forEach(key => {
	Vue.filter(key, filters[key])
})

//全局
Vue.config.productionTip = false

new Vue({
	router,
	store,
	template: '<App/>',
	render: h => h(App)
}).$mount('#app')