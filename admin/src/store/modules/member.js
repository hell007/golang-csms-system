import {
  fetchGet,
  fetchPost
} from '@/api'
import * as storage from '@/utils/storage'
import * as crypto from '@/utils/crypto'
import {
  CSMSKEY
} from '@/api/config'


function hasAccess() {
  return storage.get(CSMSKEY.Access) ? true : false
}

function getToken() {
  return storage.get(CSMSKEY.Token)
}

const member = {
  state: {
    member: {},
    memberlist: []
  },
  mutations: {
    SET_MEMBER: (state, member) => {
      state.member = member
    },
    SET_MEMBERLIST: (state, memberlist) => {
      state.memberlist = memberlist
    }
  },
  actions: {
    getMemberList({
      commit
    }, listQuery) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchGet('/member/list?', listQuery);
            const res = response.data.data
            commit('SET_MEMBERLIST', res.rows)

            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    getMember({
      commit
    }, id) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchGet('/member/item?', {
              id: id
            });
            const item = response.data.data
            commit('SET_MEMBER', item)

            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    saveMember({
      commit,
      state
    }, form) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchPost('/member/save', form);
            const item = response.data.data
            commit('SET_MEMBER', item)

            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    deleteMember({
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

            let response = await fetchGet('/member/delete?', {
              id: ids
            });

            let list = state.memberlist
            for (let i = 0, len = rows.length; i < len; i++) {
              const index = list.indexOf(rows[i])
              list.splice(index, 1)
            }
            commit('SET_MEMBERLIST', list)

            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    closeMember({
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

            let response = await fetchGet('/member/close?', {
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

export default member