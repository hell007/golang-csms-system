/*
 * @Descripttion: 
 * @Author: zenghua.wang
 * @Date: 2021-02-23 21:12:37
 * @LastEditors: zenghua.wang
 * @LastEditTime: 2022-02-11 17:54:12
 */
import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import { vantPlugins } from './plugins/vant';
import Bmob, { globalPlugins } from './plugins/global';
import publicAxios from './utils/axios'
import VueAxios from 'vue-axios'
import { Toast, Dialog } from 'vant';

//import { Modal } from './types';

const app = createApp(App);
app.config.globalProperties.$bmob = Bmob;

// app.directive('loading', {
//   mounted(el) {
//     el.Loading();
//   }
// });

app
  .use(store)
  .use(router)
  .use(globalPlugins)
  .use(vantPlugins)
  .use(VueAxios, publicAxios)
  .use(Toast)
  .use(Dialog)
  //.use(Modal)
  .mount('#app');
