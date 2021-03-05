/**
 * name: role store
 * author: jie
 */

import {
  fetchPost,
  fetchGet
} from '@/api'


const role = {

  state: {
    rolelist: [],
    role: {}
  },
  mutations: {
    SET_ROLELIST: (state, rolelist) => {
      state.rolelist = rolelist
    },
    SET_ROLE: (state, role) => {
      state.role = role
    }
  },
  actions: {
    getRoleList({
      commit
    }, listQuery) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchGet('/sys/role/list?', listQuery);
            const res = response.data.data
            commit('SET_ROLELIST', res.rows)
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    getRole({
      commit
    }, id) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchGet('/sys/role/item?', {
              id: id
            });
            const role = response.data.data
            commit('SET_ROLE', role)
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    // 保存
    saveRole({
      commit,
      state
    }, form) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchPost('/sys/role/save', form);
            const role = response.data.data
            commit('SET_ROLE', role)
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    // 删除
    deleteRole({
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

            let response = await fetchGet('/sys/role/delete?', {
              id: ids
            });

            let list = state.rolelist
            for (let i = 0, len = rows.length; i < len; i++) {
              const index = list.indexOf(rows[i])
              list.splice(index, 1)
            }
            commit('SET_ROLELIST', list)
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    // 停用
    closeRole({
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

            let response = await fetchGet('/sys/role/close?', {id: ids});
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    }
  }
}

export default role