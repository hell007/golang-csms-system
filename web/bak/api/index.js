/**
 * name: axios 的封装
 * user: jie
 */

import axios from '@/utils/axios'

import { obj2Param } from '@/utils'

/*get*/
export function fetchGet(url, query) {
  return axios({
    url: url + obj2Param(query),
    method: 'get'
  })
}

/*post*/
export function fetchPost(url, form) {
  return axios({
    url: url, //mock在这里拦截,返回数据
    method: 'post',
    data: form,
  })
}