import {
  fetchGet,
  fetchPost
} from '@/api'



const unit = {
  state: {
    unit: {}
  },
  mutations: {
    SET_UNIT: (state, unit) => {
      state.unit = unit
    }
  },
  actions: {
    // 列表
    getUnitList({
      commit
    }, listQuery) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchGet('/unit/list?', listQuery);
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    // 获取
    getUnit({
      commit
    }, id) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchGet('/unit/item?', {
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
    saveUnit({
      commit,
      state
    }, form) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchPost('/unit/save', form);
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    // 删除
    deleteUnit({
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

            let response = await fetchGet('/unit/delete?', {
              id: ids
            });
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    }
  }
}

export default unit