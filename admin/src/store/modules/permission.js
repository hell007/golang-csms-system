import {
  asyncRouter,
  constantRouter,
  myimport
} from '@/router'

import {
  fetchGet,
  fetchPost,
} from '@/api'

import {
  getTree
} from '@/utils'

import Layout from '@/views/layout/Layout'
import View from '@/views/layout/View'


/**
 * 通过meta.role判断是否与当前用户权限匹配
 * @param roles
 * @param route
 */
function hasPermission(roles, route) {
  if (route.meta && route.meta.role) {
    return roles.some(role => route.meta.role.indexOf(role) >= 0)
  } else {
    return true
  }
}

/**
 * 递归过滤异步路由表，返回符合用户角色权限的路由表
 * @param asyncRouter
 * @param roles
 */
function filterAsyncRouter(asyncRouter, roles) {
  const accessedRouters = asyncRouter.filter(route => {
    if (hasPermission(roles, route)) {
      if (route.children && route.children.length) {
        route.children = filterAsyncRouter(route.children, roles)
      }
      return true
    }
    return false
  })
  return accessedRouters
}


const permission = {
  state: {
    routers: constantRouter,
    addRouters: []
  },
  mutations: {
    SET_ROUTERS: (state, routers) => {
      state.addRouters = routers
      state.routers = constantRouter.concat(routers)
    }
  },
  actions: {
    generateRoutes({
      commit
    }, data) {
      let roles = data.roles
      let access = data.access
      let asyncRouter = []
      let accessList = []
      let accessedRouters

      return new Promise(resolve => {
        // if (roles.indexOf('superadmin') >= 0) {
        //   accessedRouters = asyncRouter
        // } else {

          // 三级 level = 123
          // 1 2 级需要菜单显示 hidden=false
          for (const a of access) {
            switch (a.level) {
              case 1:
                a.noDropdown = true
                a.component = Layout
                a.hidden = false
                break;
              case 2:
                a.component = View
                a.hidden = false
                a.redirect = a.redirect ? `/${a.redirect}` : ''
                a.path = `/${a.path}`
                break;
              case 3:
                a.component = myimport(a.redirect)
                a.hidden = true
                delete a.redirect
                break;
            }

            a.icon = a.icon ? `xa-icon ${a.icon}` : ''
            a.meta = {
              role: roles
            }
          }

          accessList = getTree(access)

          for (const p of accessList) {
            asyncRouter.push(p)
          }

          accessedRouters = filterAsyncRouter(asyncRouter, roles)

        //}
        commit('SET_ROUTERS', accessedRouters)
        resolve(accessedRouters)
      })
    },
    // 列表
    getPermissionList({
      commit
    }) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchGet('/sys/permission/list');
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    // 获取
    getPermissions({
      commit
    }, id) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchGet('/sys/permission/list?', {level:2});
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    // 获取
    getPermission({
      commit
    }, id) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchGet('/sys/permission/item?', {
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
    savePermission({
      commit,
      state
    }, form) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            let response = await fetchPost('/sys/permission/save', form);
            resolve(response)
          } catch (ex) {
            reject(ex)
          }
        })();
      })
    },
    // 删除
    deletePermission({
      commit,
      state
    }, rows) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            // rows [1,2,3,...]
            let ids = []
            rows.map(function(row) {
              ids.push(row)
            })

            let response = await fetchGet('/sys/permission/delete?', {
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
    closePermission({
      commit,
      state
    }, rows) {
      return new Promise((resolve, reject) => {
        (async() => {
          try {
            //rows [1,2,3,...]
            let ids = []
            rows.map(function(row) {
              ids.push(row)
            })

            let response = await fetchGet('/sys/permission/close?', {
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

export default permission