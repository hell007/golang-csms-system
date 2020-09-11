import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store'
import './auth' // 权限
import './mock' // 该项目所有请求使用mockjs模拟

// 组件
import Element from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
Vue.use(Element)

// 全局filter及其使用
import * as filters from './filters' 
Object.keys(filters).forEach(key => {
	Vue.filter(key, filters[key])
})

//全局
Vue.config.productionTip = false

window.drag = new Vue();

new Vue({
	router,
	store,
	template: '<App/>',
	render: h => h(App)
}).$mount('#app')