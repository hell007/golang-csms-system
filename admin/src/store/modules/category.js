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
    categorylist: [],
    category: {}
  },
  mutations: {
    SET_CATEGORYLIST: (state, categorylist) => {
      state.categorylist = categorylist
    },
    SET_CATEGORY: (state, category) => {
      state.category = category
    }
  },
  actions: {
    getCategoryList({
      commit
    }, listQuery) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchGet('/goods/category/list?', listQuery);
            const res = response.data.data
            commit('SET_CATEGORYLIST', res)
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
            const item = response.data.data
            commit('SET_CATEGORY', item)
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
            const item = response.data.data
            commit('SET_CATEGORY', item)
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

            let list = state.categorylist
            for (let i = 0, len = rows.length; i < len; i++) {
              const index = list.indexOf(rows[i])
              list.splice(index, 1)
            }
            commit('SET_CATEGORYLIST', list)
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