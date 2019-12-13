import Vue from 'vue'
import Vuex from 'vuex'
import services from './modules/services'
import user from './modules/user'
import order from './modules/order'
import getters from './getters'

//使用vuex
Vue.use(Vuex)

const store = new Vuex.Store({
	modules: {
		services,
		user,
		order
	},
	getters
})

export default store