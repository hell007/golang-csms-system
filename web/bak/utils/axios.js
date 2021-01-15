import axios from 'axios'
import router from '../router'

import {
  hasToken,
  getToken
} from '@/utils/storage'

// 创建axios实例
const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_API, // api的base_url
  timeout: 10000 // 请求超时时间
})

// request拦截器
service.interceptors.request.use(config => {
  // Do something before request is sent
  if (hasToken()) {
    config.headers['Authorization'] = 'bearer ' + getToken() // 让每个请求携带token--['X-Token']为自定义key 请根据实际情况自行修改
  }
  return config
}, error => {
  Promise.reject(error)
})

// respone拦截器
service.interceptors.response.use(
  response => response,
  error => {
    console.log('error.response=', error.response)
    if(error && error.response) {
      let status = error.response.status
      switch(status) {
        case 404:
          router.push('/404')
          break;
        case 500:
          router.push('/500')
          break;  
      }
    }
    return Promise.reject(error)
  }
)

export default service