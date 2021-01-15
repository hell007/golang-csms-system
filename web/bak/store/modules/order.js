import {
  fetchPost,
  fetchGet
} from '@/api'


const order = {
  actions: {
    getOrderList({
      commit
    }, listQuery) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchGet('/api/order/v1?', listQuery);
            const res = response.data.data
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
            let response = await fetchGet('/api/order/item?', {
              id: id
            });
            const item = response.data.data
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
            let response = await fetchPost('/api/order/save', form);
            const item = response.data.data
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