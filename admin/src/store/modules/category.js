/**
 * name: category store
 * author: jie
 */

import {
  fetchPost,
  fetchGet
} from '@/api'


const category = {

  state: {
  },
  mutations: {
  },
  actions: {
    getCategoryList({
      commit
    }, listQuery) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchGet('/goods/category/list?', listQuery);
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    getCategory({
      commit
    }, id) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchGet('/goods/category/item?', {
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
    saveCategory({
      commit,
      state
    }, form) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchPost('/goods/category/save', form);
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    // 删除
    deleteCategory({
      commit,
      state
    }, rows) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            // 后台需要数组
            //console.log("rows===",rows)
            let ids = []
            rows.map(function(row) {
              ids.push(row)
            })

            let response = await fetchGet('/goods/category/delete?', {
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
    closeCategory({
      commit,
      state
    }, rows) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let ids = []
            rows.map(function(row) {
              ids.push(row)
            })

            let response = await fetchGet('/goods/category/close?', {id: ids});
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    }
  }
}

export default category