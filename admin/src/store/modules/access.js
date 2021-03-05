/**
 * name: access store
 * author: jie
 */

import {
  fetchGet,
  fetchPost
} from '@/api'


const access = {

  state: {
    accesslist: []
  },
  mutations: {
    SET_ACCESSLIST: (state, accesslist) => {
      state.accesslist = accesslist
    }
  },
  actions: {
    // 列表
    getAccessList({
      commit
    }, rid) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchGet('/sys/access/list?', {rid:rid});
            const res = response.data.data
            commit('SET_ACCESSLIST', res.rows)
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    // 保存
    saveAccess({
      commit,
      state
    }, form) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchPost('/sys/access/save', form);
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    }
  }
}

export default access