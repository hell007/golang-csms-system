/**
 * name: goods store
 * author: jie
 */

import {
  fetchPost,
  fetchGet
} from '@/api'


const goods = {

  state: {
    goodslist: [],
    goods: {}
  },
  mutations: {
    SET_GOODSLIST: (state, goodslist) => {
      state.goodslist = goodslist
    },
    SET_GOODS: (state, goods) => {
      state.goods = goods
    }
  },
  actions: {
    getGoodsList({
      commit
    }, listQuery) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchGet('/goods/product/list?', listQuery);
            const res = response.data.data
            commit('SET_GOODSLIST', res.rows)
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    getGoods({
      commit
    }, id) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchGet('/goods/product/item?', {
              id: id
            });
            const item = response.data.data
            commit('SET_GOODS', item.goods)
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    // 保存
    saveGoods({
      commit,
      state
    }, form) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchPost('/goods/product/save', form);
            const item = response.data.data
            commit('SET_GOODS', item)
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    saveGallery({
      commit,
      state
    }, form) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchPost('/goods/gallery/upload', form);
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    deleteGallery({
      commit
    }, form) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchPost('/goods/gallery/delete', form);
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    saveSku({
      commit,
      state
    }, form) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchPost('/goods/sku/val', form);
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    // 删除
    deleteGoods({
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

            let response = await fetchGet('/goods/product/delete?', {
              id: ids
            });

            let list = state.goodslist
            for (let i = 0, len = rows.length; i < len; i++) {
              const index = list.indexOf(rows[i])
              list.splice(index, 1)
            }
            commit('SET_GOODSLIST', list)
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    // 下架
    closeGoods({
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

            let response = await fetchGet('/goods/product/close?', {id: ids});
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    }
  }
}

export default goods