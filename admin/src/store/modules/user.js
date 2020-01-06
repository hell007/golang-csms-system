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

const user = {
  //State负责存储整个应用的状态数据，一般需要在使用的时候在跟节点注入store对象，后期就可以使用this.$store.state.xxx直接获取状态。
  //即：...mapState(['token']) = this.$store.state.token

  state: {
    admin: {},
    access: [],
    roles: []
  },

  //Mutations的中文意思是“变化”，利用它可以更改状态，本质就是用来处理数据的函数，其接收唯一参数值state
  //store.commit(mutationName)是用来触发一个mutation的方法。
  //即：...mapMutations(['SET_USER'])调用

  //需要记住的是，定义的mutation必须是同步函数，否则devtool中的数据将可能出现问题，使状态改变变得难以跟踪。
  mutations: {
    SET_ADMIN: (state, admin) => {
      state.admin = admin
    },
    SET_ACCESS: (state, access) => {
      state.access = access
    },
    SET_ROLES: (state, roles) => {
      state.roles = roles
    }
  },

  //Actions也可以用于改变状态，不过是通过触发mutation实现的，重要的是可以包含异步操作。

  //其辅助函数是mapActions,与mapMutations类似，也是绑定在组件的methods上的。

  //如果选择直接触发的话，使用this.$store.dispatch(actionName)方法
  //即: ...mapActions(['actionName']) = this.$store.dispatch(actionName)
  actions: {
    // 用户名登录
    LoginByName({
      commit
    }, form) {
      return new Promise((resolve, reject) => {
        fetchPost('/sys/user/login', form).then(response => {
          const user = response.data.data
          if(user){
            const roles = []
            roles.push(user.rolename)

            //local storage
            storage.set(CSMSKEY.Token, user.token)
            storage.set(CSMSKEY.User, crypto.encrypt(user))

            //注意：这里暂时不设置roles,需要获取admin的权限菜单  
            commit('SET_ADMIN', user)
          }
          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 退出
    LoginOut({
      commit,
      state
    }) {
      return new Promise((resolve, reject) => {
        fetchPost('/sys/user/loginout', {
          token: getToken()
        }).then(response => {
          commit('SET_ADMIN', {})
          commit('SET_ROLES', [])
          commit('SET_ACCESS', [])

          storage.clear()

          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 登录管理员
    getAdmin({
      commit,
      state
    }) {
      // 页面刷新 state都会丢失，需要重新获取登录信息
      return new Promise((resolve, reject) => {
        const user = JSON.parse(crypto.decrypt(storage.get(CSMSKEY.User)))

        let result = {
          roles: [],
          access: []
        }

        if (hasAccess()) {
          result.access = JSON.parse(crypto.decrypt(storage.get(CSMSKEY.Access)))
          result.roles.push(user.rolename)

          commit('SET_ADMIN', user)
          commit('SET_ROLES', result.roles)
          commit('SET_ACCESS', result.access)

          resolve(result)
        } else {

          fetchGet('/sys/rule/menus?', {
            rid: user.roleId
          }).then(response => {
            const res = response.data.data
            result.access = res.rows
            result.roles.push(user.rolename)

            // local storage
            storage.set(CSMSKEY.Access, crypto.encrypt(res.rows))

            commit('SET_ADMIN', user)
            commit('SET_ROLES', result.roles)
            commit('SET_ACCESS', result.access)

            resolve(result)
          }).catch(error => {
            reject(error)
          })
        }
      })
    },
    // 列表
    getUserList({
      commit
    }, listQuery) {
      //调用 api/index.js 里面 fetchGet
      //es6 Promise
      return new Promise((resolve, reject) => {
        // es7 async await 
        (async() => {
          try {
            let response = await fetchGet('/sys/user/list?', listQuery);
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    // 获取
    getUser({
      commit
    }, id) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchGet('/sys/user/item?', {
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
    saveUser({
      commit,
      state
    }, form) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchPost('/sys/user/save', form);
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    // 删除
    deleteUser({
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

            let response = await fetchGet('/sys/user/delete?', {
              id: ids
            });
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    // 停用
    closeUser({
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
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    }
  },

  //Getters 对State 里面的数据二次处理（对数据进行过滤类似filter的作用）
  //比如State返回的为一个对象，我们想取对象中一个键的值用这个方法
  //通过this.$store.getters.user对派生出来的状态进行访问。
  //或者直接使用辅助函数mapGetters将其映射到本地计算属性中去
  //即: ...mapGetters(['user']) = this.$store.getters.user

  getters: {
    test: state => state.access
  }

}

export default user