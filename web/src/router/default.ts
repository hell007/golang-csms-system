/*
 * @Descripttion: 
 * @Author: zenghua.wang
 * @Date: 2021-02-23 21:12:37
 * @LastEditors: zenghua.wang
 * @LastEditTime: 2022-02-11 15:32:00
 */
import { RouteRecordRaw } from 'vue-router';

const unitList: Array<RouteRecordRaw> = [
  {
    path: '/unit',
    name: '单元',
    component: () => import(/* webpackChunkName: "unit-index" */ '@/views/unit/index.vue')
  },
  {
    path: '/unitList',
    name: '单元-上拉加载',
    component: () => import(/* webpackChunkName: "unit-list" */ '@/views/unit/list.vue')
  },
  {
    path: '/unitRefresh',
    name: '单元-下拉刷新',
    component: () => import(/* webpackChunkName: "unit-refresh" */ '@/views/unit/refresh.vue')
  },
  {
    path: '/unitUtils',
    name: '单元-工具类',
    component: () => import(/* webpackChunkName: "unit-utils" */ '@/views/unit/utils.vue')
  },
  {
    path: '/unitDialog',
    name: '单元-弹窗使用',
    component: () => import(/* webpackChunkName: "unit-dialog" */ '@/views/unit/modal.vue')
  },
  {
    path: '/',
    name: '单元-测试',
    component: () => import(/* webpackChunkName: "unit-dialog" */ '@/views/unit/test.vue')
  }
];

export { unitList };
