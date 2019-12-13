import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store'
import * as filters from './filters' // 全局filter
import './auth' // 权限

//自定义组件
import toast from '@/components/toast/index'
Vue.use(toast)

import message from '@/components/message/index'
Vue.use(message)


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