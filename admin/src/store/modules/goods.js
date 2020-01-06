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
  },
  mutations: {
  },
  actions: {
    getGoodsList({
      commit
    }, listQuery) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchGet('/goods/product/list?', listQuery);
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