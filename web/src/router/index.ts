import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import View from '@/views/layout/view.vue';
import { unitList } from './default';

let routes: Array<RouteRecordRaw> = [
  {
    path: '/user',
    component: View,
    redirect: '/user/profile',
    name: '个人中心',
    children: [
      {
        path: 'profile',
        //name: '个人中心',
        component: () => import('@/views/user/profile.vue')
      },
      {
        path: 'address',
        component: () => import('../views/user/address.vue')
      },
      {
        path: 'editAddress',
        component: () => import('../views/user/editAddress.vue')
      },
      {
        path: 'register',
        component: () => import('../views/user/register.vue')
      },
      {
        path: 'login',
        component: () => import('../views/user/login.vue')
      }
    ]
  }
  // {
  //   path: '/order',
  //   redirect: 'order/list',
  //   name: '我的订单',
  //   children: [
  //     {
  //       path: '/list',
  //       component: () => import('../views/order/list.vue')
  //     },
  //     {
  //       path: '/detail',
  //       component: () => import('../views/order/lidetailst.vue')
  //     },
  //     {
  //       path: '/confirm',
  //       component: () => import('../views/order/index.vue')
  //     }
  //   ]
  // },
  // {
  //   path: '/login',
  //   component: () => import('../views/user/login.vue')
  // },
  // {
  //   path: '/404',
  //   component: () => import('../views/error/404.vue')
  // },
  // {
  //   path: '/500',
  //   component: () => import('../views/error/500.vue')
  // },
  // {
  //   path: '/',
  //   component: () => import('../views/home/index.vue')
  // },
  // {
  //   path: '/category',
  //   component: () => import('../views/category/index.vue')
  // },
  // {
  //   path: '/search',
  //   component: () => import('../views/search/index.vue')
  // },
  // {
  //   path: '/product',
  //   component: () => import('../views/product/index.vue')
  // }
];

for (let item of unitList) {
  routes.push(item);
}

//routes.concat(unitList);

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
});

export default router;
