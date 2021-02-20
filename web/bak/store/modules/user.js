import {
  fetchGet,
  fetchPost
} from '@/api'

const user = {
  state: {
    token: {},
  },
  mutations: {
    SET_TOKEN: (state, token) => {
      state.token = token
    }
  },
  actions: {
    getUser({
      commit
    }, query) {
      return new Promise((resolve, reject) => {
        (async () => {
          try {
            let response = await fetchGet('/api/user/v1?', query);
            const res = response.data.data
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    saveUser({
      commit,
      state
    }, form) {
      return new Promise((resolve, reject) => {
        (async () => {
          try {
            let response = await fetchPost('/api/user/find', form);
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    getCity({
      commit
    }, query) {
      return new Promise((resolve, reject) => {
        (async () => {
          try {
            let response = await fetchGet('/api/user/city?', query);
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

export default user