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
    login({
      commit,
      state
    }, form) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchPost('/api/user/login', form);
            const res = response.data.data
            commit("SET_TOKEN", res)
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    register({
      commit,
      state
    }, form) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchPost('/api/user/register', form);
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    logout({
      commit
    }, query) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchGet('/api/user/logout?', query);
            const res = response.data.data
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    getUser({
      commit
    }, query) {
      return new Promise((resolve, reject) => {
        (async() => {
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
        (async() => {
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
        (async() => {
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
    getAddress({
      commit
    }, query) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchGet('/api/user/address?', query);
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    saveAddress({
      commit
    }, form) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchPost('/api/user/saveaddr', form);
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    delAddress({
      commit
    }, query) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchGet('/api/user/deladdr?', query);
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