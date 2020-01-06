import axios from '@/utils/axios'
import { obj2Param } from '@/utils'

/**
 * get方法
 * @param  {[type]} url 
 * @param  {[type]} query
 * @return {[type]}       
 */
export function fetchGet(url, query) {
  return axios({
    url: url + obj2Param(query), //mock在这里拦截,返回数据
    method: 'get'
  })
}

/**
 * post请求
 * @param  {[type]} url
 * @param  {[type]} form 
 * @return {[type]}      
 */
export function fetchPost(url, form) {
  return axios({
    url: url, //mock在这里拦截,返回数据
    method: 'post',
    data: form,
  })
}