/**
 * name: order store
 * author: jie
 */

import {
  fetchPost,
  fetchGet
} from '@/api'


const order = {

  state: {
    orderlist: [],
    order: {}
  },
  mutations: {
    SET_ORDERLIST: (state, orderlist) => {
      state.orderlist = orderlist
    },
    SET_ORDER: (state, order) => {
      state.order = order
    }
  },
  actions: {
    getOrderList({
      commit
    }, listQuery) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchGet('/order/list?', listQuery);
            const res = response.data.data
            commit('SET_ORDERLIST', res.rows)
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    getOrder({
      commit
    }, id) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchGet('/order/item?', {
              ordersn: id
            });
            const item = response.data.data
            commit('SET_ORDER', item)
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    // 保存
    saveOrder({
      commit,
      state
    }, form) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchPost('/order/save', form);
            const item = response.data.data
            commit('SET_ORDER', item)
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    // 删除
    deleteOrder({
      commit,
      state
    }, rows) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            // 后台需要数组
            let ids = []
            rows.map(function(row) {
              ids.push(row.id)
            })

            let response = await fetchGet('/order/delete?', {
              id: ids
            });

            let list = state.orderlist
            for (let i = 0, len = rows.length; i < len; i++) {
              const index = list.indexOf(rows[i])
              list.splice(index, 1)
            }
            commit('SET_ORDERLIST', list)
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    // 下架
    closeOrder({
      commit,
      state
    }, rows) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let ids = []
            rows.map(function(row) {
              ids.push(row.id)
            })

            let response = await fetchGet('/order/close?', {id: ids});
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    }
  }
}

export default order