
import {
  fetchPost,
  fetchGet
} from '@/api'


const services = {

  state: {
    // orderlist: [],
    // order: {}
  },
  mutations: {
    // SET_ORDERLIST: (state, orderlist) => {
    //   state.orderlist = orderlist
    // },
    // SET_ORDER: (state, order) => {
    //   state.order = order
    // }
  },
  actions: {
    getHome({
      commit
    }, listQuery) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchGet('/api/home/v1?', listQuery);
            const res = response.data.data
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    getSearch({
      commit
    }, listQuery) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchGet('/api/search/v1?', listQuery);
            const res = response.data.data
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    getCategory({
      commit
    }) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchGet('/api/category/v1?', {});
            const res = response.data.data
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    getCategoryGoods({
      commit
    }, listQuery) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchGet('/api/category/goods?', listQuery);
            const res = response.data.data
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    getGoods({
      commit
    }, listQuery) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchGet('/api/product/v1?', listQuery);
            const res = response.data.data
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    postGoods({
      commit
    }, form) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchPost('/api/product/do', form);
            const res = response.data.data
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
  }
}



export default services