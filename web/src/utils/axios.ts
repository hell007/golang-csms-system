/*
 * @Descripttion: 
 * @Author: zenghua.wang
 * @Date: 2021-02-23 21:12:37
 * @LastEditors: zenghua.wang
 * @LastEditTime: 2022-02-11 17:53:40
 */
import axios from 'axios';
import { hasToken, getToken } from './storage';

// 创建axios实例
const publicAxios = axios.create({
  baseURL: process.env.VUE_APP_BASE_API, // api的base_url
  timeout: 10000
});

// request拦截器
publicAxios.interceptors.request.use(
  function(config) {
    if (hasToken()) {
      // 让每个请求携带token--['X-Token']为自定义key 请根据实际情况自行修改
      config.headers['Authorization'] = getToken();
      config.headers['cache-control'] = 'no-cache';
      config.headers['Pragma'] = 'no-cache';
    }
    return config;
  },
  function(error) {
    console.log(error);
    Promise.reject(error);
  }
);

// respone拦截器
// publicAxios.interceptors.response.use(
//   (response) => response,
//   (error) => {
//     console.log('error.response=', error.response);
//     if (error && error.response) {
//       let status = error.response.status;
//       switch (status) {
//         case 404:
//           router.push('/404');
//           break;
//         case 500:
//           router.push('/500');
//           break;
//       }
//     }
//     return Promise.reject(error);
//   }
// );

export default publicAxios;
